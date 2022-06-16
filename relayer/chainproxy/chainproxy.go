package chainproxy

import (
	"context"
	"fmt"
	"time"

	"github.com/btcsuite/btcd/btcec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/relayer/sentry"
	"github.com/lavanet/lava/relayer/sigs"
	pairingtypes "github.com/lavanet/lava/x/pairing/types"
	spectypes "github.com/lavanet/lava/x/spec/types"
)

type NodeMessage interface {
	GetServiceApi() *spectypes.ServiceApi
	Send(ctx context.Context) (*pairingtypes.RelayReply, error)
	RequestedBlock() int64
}

type ChainProxy interface {
	Start(context.Context) error
	GetSentry() *sentry.Sentry
	ParseMsg(string, []byte) (NodeMessage, error)
	PortalStart(context.Context, *btcec.PrivateKey, string)
}

func GetChainProxy(nodeUrl string, nConns uint, sentry *sentry.Sentry) (ChainProxy, error) {
	switch sentry.ApiInterface {
	case "jsonrpc":
		return NewJrpcChainProxy(nodeUrl, nConns, sentry), nil
	case "tendermintrpc":
		return NewtendermintRpcChainProxy(nodeUrl, nConns, sentry), nil
	case "rest":
		return NewRestChainProxy(nodeUrl, sentry), nil
	}
	return nil, fmt.Errorf("chain proxy for apiInterface (%s) not found", sentry.ApiInterface)
}

func VerifyRelayReply(reply *pairingtypes.RelayReply, relayRequest *pairingtypes.RelayRequest, addr string, comparesHashes bool) error {

	serverKey, err := sigs.RecoverPubKeyFromRelayReply(reply, relayRequest)
	if err != nil {
		return err
	}
	serverAddr, err := sdk.AccAddressFromHex(serverKey.Address().String())
	if err != nil {
		return err
	}
	if serverAddr.String() != addr {
		return fmt.Errorf("server address mismatch in reply (%s) (%s)", serverAddr.String(), addr)
	}

	if comparesHashes {
		serverKey, err = sigs.RecoverPubKeyFromResponseFinalizationData(reply, relayRequest)
		if err != nil {
			return err
		}

		serverAddr, err = sdk.AccAddressFromHex(serverKey.Address().String())
		if err != nil {
			return err
		}

		if serverAddr.String() != addr {
			return fmt.Errorf("server address mismatch in reply sigblocks (%s) (%s)", serverAddr.String(), addr)
		}
	}

	return nil
}

// Client requests and queries
func SendRelay(
	ctx context.Context,
	cp ChainProxy,
	privKey *btcec.PrivateKey,
	url string,
	req string,
) (*pairingtypes.RelayReply, error) {

	//
	// Unmarshal request
	nodeMsg, err := cp.ParseMsg(url, []byte(req))
	if err != nil {
		return nil, err
	}
	blockHeight := int64(-1) //to sync reliability blockHeight in case it changes
	requestedBlock := int64(0)
	callback_send_relay := func(clientSession *sentry.ClientSession) (*pairingtypes.RelayReply, *pairingtypes.RelayRequest, error) {
		//client session is locked here
		err := CheckComputeUnits(clientSession, nodeMsg.GetServiceApi().ComputeUnits)
		if err != nil {
			return nil, nil, err
		}

		blockHeight = cp.GetSentry().GetBlockHeight()
		relayRequest := &pairingtypes.RelayRequest{
			Provider:        clientSession.Client.Acc,
			ApiUrl:          url,
			Data:            []byte(req),
			SessionId:       uint64(clientSession.SessionId),
			ChainID:         cp.GetSentry().ChainID,
			CuSum:           clientSession.CuSum,
			BlockHeight:     blockHeight,
			RelayNum:        clientSession.RelayNum,
			RequestBlock:    nodeMsg.RequestedBlock(),
			QoSReport:       clientSession.QoSInfo.LastQoSReport,
			DataReliability: nil,
		}

		sig, err := sigs.SignRelay(privKey, *relayRequest)
		if err != nil {
			return nil, nil, err
		}
		relayRequest.Sig = sig
		c := *clientSession.Client.Client
		relaySentTime := time.Now()
		clientSession.QoSInfo.TotalRelays++
		reply, err := c.Relay(ctx, relayRequest)
		if err != nil {
			if err.Error() == context.DeadlineExceeded.Error() { //check if this is right
				clientSession.QoSInfo.ConsecutiveTimeOut++
			}
			return nil, nil, err
		}
		currentLatency := time.Since(relaySentTime)
		clientSession.QoSInfo.ConsecutiveTimeOut = 0
		clientSession.QoSInfo.AnsweredRelays++

		//update relay request requestedBlock to the provided one in case it was arbitrary
		sentry.UpdateRequestedBlock(relayRequest, reply)
		requestedBlock = relayRequest.RequestBlock

		err = VerifyRelayReply(reply, relayRequest, clientSession.Client.Acc, cp.GetSentry().GetSpecComparesHashes())
		if err != nil {
			return nil, nil, err
		}

		clientSession.CalculateQoS(nodeMsg.GetServiceApi().ComputeUnits, currentLatency)

		return reply, relayRequest, nil
	}
	callback_send_reliability := func(clientSession *sentry.ClientSession, dataReliability *pairingtypes.VRFData) (*pairingtypes.RelayReply, error) {
		//client session is locked here

		if blockHeight < 0 {
			return nil, fmt.Errorf("expected callback_send_relay to be called first and set blockHeight")
		}

		relayRequest := &pairingtypes.RelayRequest{
			Provider:        clientSession.Client.Acc,
			ApiUrl:          url,
			Data:            []byte(req),
			SessionId:       uint64(0), //sessionID for reliability is 0
			ChainID:         cp.GetSentry().ChainID,
			CuSum:           clientSession.CuSum,
			BlockHeight:     blockHeight,
			RelayNum:        clientSession.RelayNum,
			RequestBlock:    requestedBlock,
			QoSReport:       nil,
			DataReliability: dataReliability,
		}

		sig, err := sigs.SignRelay(privKey, *relayRequest)
		if err != nil {
			return nil, err
		}
		relayRequest.Sig = sig

		sig, err = sigs.SignVRFData(privKey, relayRequest.DataReliability)
		if err != nil {
			return nil, err
		}
		relayRequest.DataReliability.Sig = sig
		c := *clientSession.Client.Client
		reply, err := c.Relay(ctx, relayRequest)
		if err != nil {
			return nil, err
		}

		err = VerifyRelayReply(reply, relayRequest, clientSession.Client.Acc, cp.GetSentry().GetSpecComparesHashes())
		if err != nil {
			return nil, err
		}

		return reply, nil
	}
	//
	//
	reply, err := cp.GetSentry().SendRelay(ctx, callback_send_relay, callback_send_reliability, nodeMsg.GetServiceApi().Category)

	return reply, err
}

func CheckComputeUnits(clientSession *sentry.ClientSession, apiCu uint64) error {
	clientSession.Client.SessionsLock.Lock()
	defer clientSession.Client.SessionsLock.Unlock()

	if clientSession.Client.UsedComputeUnits+apiCu > clientSession.Client.MaxComputeUnits {
		return fmt.Errorf("used all the available compute units")
	}

	clientSession.CuSum += apiCu
	clientSession.Client.UsedComputeUnits += apiCu
	clientSession.RelayNum += 1
	return nil
}
