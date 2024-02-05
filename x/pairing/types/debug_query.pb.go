// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/pairing/debug_query.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/x/staking/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/lavanet/lava/x/dualstaking/types"
	_ "github.com/lavanet/lava/x/epochstorage/types"
	_ "github.com/lavanet/lava/x/projects/types"
	_ "github.com/lavanet/lava/x/subscription/types"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QueryDebugQueryResponse struct {
	BlockReward              uint64 `protobuf:"varint,1,opt,name=block_reward,json=blockReward,proto3" json:"block_reward,omitempty"`
	ValDistPoolBalance       uint64 `protobuf:"varint,2,opt,name=val_dist_pool_balance,json=valDistPoolBalance,proto3" json:"val_dist_pool_balance,omitempty"`
	ValAllocPoolBalance      uint64 `protobuf:"varint,3,opt,name=val_alloc_pool_balance,json=valAllocPoolBalance,proto3" json:"val_alloc_pool_balance,omitempty"`
	ProviderDistPoolBalance  uint64 `protobuf:"varint,4,opt,name=provider_dist_pool_balance,json=providerDistPoolBalance,proto3" json:"provider_dist_pool_balance,omitempty"`
	ProviderAllocPoolBalance uint64 `protobuf:"varint,5,opt,name=provider_alloc_pool_balance,json=providerAllocPoolBalance,proto3" json:"provider_alloc_pool_balance,omitempty"`
	ProviderFullReward       uint64 `protobuf:"varint,6,opt,name=provider_full_reward,json=providerFullReward,proto3" json:"provider_full_reward,omitempty"`
	ProviderRewardNoBonus    uint64 `protobuf:"varint,7,opt,name=provider_reward_no_bonus,json=providerRewardNoBonus,proto3" json:"provider_reward_no_bonus,omitempty"`
	Block                    uint64 `protobuf:"varint,8,opt,name=block,proto3" json:"block,omitempty"`
	MonthsLeft               uint64 `protobuf:"varint,9,opt,name=months_left,json=monthsLeft,proto3" json:"months_left,omitempty"`
	ValidatorReward          uint64 `protobuf:"varint,10,opt,name=validator_reward,json=validatorReward,proto3" json:"validator_reward,omitempty"`
}

func (m *QueryDebugQueryResponse) Reset()         { *m = QueryDebugQueryResponse{} }
func (m *QueryDebugQueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDebugQueryResponse) ProtoMessage()    {}
func (*QueryDebugQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a315cb51b0edeaef, []int{0}
}
func (m *QueryDebugQueryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDebugQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDebugQueryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDebugQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDebugQueryResponse.Merge(m, src)
}
func (m *QueryDebugQueryResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDebugQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDebugQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDebugQueryResponse proto.InternalMessageInfo

func (m *QueryDebugQueryResponse) GetBlockReward() uint64 {
	if m != nil {
		return m.BlockReward
	}
	return 0
}

func (m *QueryDebugQueryResponse) GetValDistPoolBalance() uint64 {
	if m != nil {
		return m.ValDistPoolBalance
	}
	return 0
}

func (m *QueryDebugQueryResponse) GetValAllocPoolBalance() uint64 {
	if m != nil {
		return m.ValAllocPoolBalance
	}
	return 0
}

func (m *QueryDebugQueryResponse) GetProviderDistPoolBalance() uint64 {
	if m != nil {
		return m.ProviderDistPoolBalance
	}
	return 0
}

func (m *QueryDebugQueryResponse) GetProviderAllocPoolBalance() uint64 {
	if m != nil {
		return m.ProviderAllocPoolBalance
	}
	return 0
}

func (m *QueryDebugQueryResponse) GetProviderFullReward() uint64 {
	if m != nil {
		return m.ProviderFullReward
	}
	return 0
}

func (m *QueryDebugQueryResponse) GetProviderRewardNoBonus() uint64 {
	if m != nil {
		return m.ProviderRewardNoBonus
	}
	return 0
}

func (m *QueryDebugQueryResponse) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *QueryDebugQueryResponse) GetMonthsLeft() uint64 {
	if m != nil {
		return m.MonthsLeft
	}
	return 0
}

func (m *QueryDebugQueryResponse) GetValidatorReward() uint64 {
	if m != nil {
		return m.ValidatorReward
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryDebugQueryResponse)(nil), "lavanet.lava.pairing.QueryDebugQueryResponse")
}

func init() {
	proto.RegisterFile("lavanet/lava/pairing/debug_query.proto", fileDescriptor_a315cb51b0edeaef)
}

var fileDescriptor_a315cb51b0edeaef = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x13, 0xfa, 0x07, 0xd8, 0x22, 0x81, 0x4c, 0x4a, 0xad, 0x20, 0x99, 0xbf, 0xa2, 0x20,
	0x50, 0x4c, 0xd4, 0x03, 0x07, 0xc4, 0xa1, 0x51, 0xc5, 0x09, 0x21, 0xc8, 0x91, 0x8b, 0x35, 0xb6,
	0x27, 0x8e, 0xe9, 0xc4, 0xb3, 0xec, 0xae, 0x5d, 0xfa, 0x16, 0x3c, 0x16, 0xc7, 0x1e, 0x39, 0xa2,
	0xe4, 0xcc, 0x3b, 0x20, 0xaf, 0xbd, 0x16, 0xa6, 0x3d, 0xed, 0xcc, 0x7c, 0xbf, 0x6f, 0xbf, 0x64,
	0xb4, 0x16, 0xcf, 0x08, 0x2a, 0x28, 0xd0, 0x84, 0xf5, 0x19, 0x4a, 0xc8, 0x55, 0x5e, 0x64, 0x61,
	0x8a, 0x71, 0x99, 0x45, 0xdf, 0x4a, 0x54, 0xe7, 0x13, 0xa9, 0xd8, 0xb0, 0x37, 0x6a, 0xb9, 0x49,
	0x7d, 0x4e, 0x5a, 0x6e, 0x3c, 0xca, 0x38, 0x63, 0x0b, 0x84, 0x75, 0xd5, 0xb0, 0xe3, 0x20, 0x63,
	0xce, 0x08, 0x43, 0xdb, 0xc5, 0xe5, 0x22, 0x3c, 0x53, 0x20, 0x25, 0x2a, 0xdd, 0xea, 0x2f, 0x7b,
	0x99, 0x28, 0x39, 0x59, 0x6a, 0xc3, 0x0a, 0x32, 0x0c, 0xb5, 0x81, 0x53, 0x8c, 0xb0, 0x30, 0x2e,
	0x78, 0x7c, 0xd8, 0x83, 0xd3, 0x12, 0xa8, 0x66, 0x9a, 0x1f, 0x49, 0x98, 0x81, 0xc1, 0x16, 0x7c,
	0x9a, 0xb0, 0x5e, 0xb1, 0x0e, 0x9d, 0x5c, 0x4d, 0x63, 0x34, 0x30, 0x75, 0x7d, 0x4b, 0xbd, 0xea,
	0x5d, 0xa7, 0xcb, 0x58, 0x27, 0x2a, 0x97, 0x26, 0xe7, 0xa2, 0xd7, 0xb4, 0xf4, 0x93, 0xfe, 0x76,
	0x14, 0x7f, 0xc5, 0xc4, 0x68, 0x57, 0x34, 0xd0, 0xe3, 0x3f, 0x5b, 0xe2, 0xe0, 0x73, 0xbd, 0xaa,
	0x93, 0x7a, 0x6b, 0xb6, 0x9a, 0xa3, 0x96, 0x5c, 0x68, 0xf4, 0x1e, 0x89, 0x5b, 0x31, 0x71, 0x72,
	0x1a, 0x29, 0x3c, 0x03, 0x95, 0xfa, 0xc3, 0x87, 0xc3, 0xe7, 0xdb, 0xf3, 0x3d, 0x3b, 0x9b, 0xdb,
	0x91, 0x37, 0x15, 0xfb, 0x15, 0x50, 0x94, 0xe6, 0xda, 0x44, 0x92, 0x99, 0xa2, 0x18, 0x08, 0x8a,
	0x04, 0xfd, 0x6b, 0x96, 0xf5, 0x2a, 0xa0, 0x93, 0x5c, 0x9b, 0x4f, 0xcc, 0x34, 0x6b, 0x14, 0xef,
	0x48, 0xdc, 0xab, 0x2d, 0x40, 0xc4, 0x49, 0xdf, 0xb3, 0x65, 0x3d, 0x77, 0x2b, 0xa0, 0xe3, 0x5a,
	0xfc, 0xd7, 0xf4, 0x56, 0x8c, 0xa5, 0xe2, 0x2a, 0x4f, 0x51, 0x5d, 0x11, 0xb6, 0x6d, 0x8d, 0x07,
	0x8e, 0xf8, 0x3f, 0xf1, 0x9d, 0xb8, 0xdf, 0x99, 0xaf, 0x88, 0xdd, 0xb1, 0x6e, 0xdf, 0x21, 0x97,
	0xb2, 0x5f, 0x8b, 0x51, 0x67, 0x5f, 0x94, 0x44, 0x6e, 0x1d, 0xbb, 0xcd, 0x5f, 0x74, 0xda, 0xfb,
	0x92, 0xa8, 0xdd, 0xca, 0x1b, 0xd1, 0xdd, 0xd6, 0xc2, 0x51, 0xc1, 0x51, 0xcc, 0x45, 0xa9, 0xfd,
	0xeb, 0xd6, 0xb5, 0xef, 0xf4, 0xc6, 0xf1, 0x91, 0x67, 0xb5, 0xe8, 0x8d, 0xc4, 0x8e, 0xdd, 0xae,
	0x7f, 0xc3, 0x52, 0x4d, 0xe3, 0x3d, 0x10, 0x7b, 0x2b, 0x2e, 0xcc, 0x52, 0x47, 0x84, 0x0b, 0xe3,
	0xdf, 0xb4, 0x9a, 0x68, 0x46, 0x1f, 0x70, 0x61, 0xbc, 0x17, 0xe2, 0x4e, 0x05, 0x94, 0xa7, 0x60,
	0xd8, 0x05, 0xfa, 0xc2, 0x52, 0xb7, 0xbb, 0x79, 0x13, 0x34, 0x3b, 0xfe, 0xb9, 0x0e, 0x86, 0x17,
	0xeb, 0x60, 0xf8, 0x7b, 0x1d, 0x0c, 0x7f, 0x6c, 0x82, 0xc1, 0xc5, 0x26, 0x18, 0xfc, 0xda, 0x04,
	0x83, 0x2f, 0x87, 0x59, 0x6e, 0x96, 0x65, 0x3c, 0x49, 0x78, 0x15, 0xf6, 0x5e, 0xce, 0xf7, 0xee,
	0xcb, 0x32, 0xe7, 0x12, 0x75, 0xbc, 0x6b, 0x5f, 0xce, 0xd1, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x75, 0x48, 0x29, 0xd9, 0x7e, 0x03, 0x00, 0x00,
}

func (m *QueryDebugQueryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDebugQueryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDebugQueryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ValidatorReward != 0 {
		i = encodeVarintDebugQuery(dAtA, i, uint64(m.ValidatorReward))
		i--
		dAtA[i] = 0x50
	}
	if m.MonthsLeft != 0 {
		i = encodeVarintDebugQuery(dAtA, i, uint64(m.MonthsLeft))
		i--
		dAtA[i] = 0x48
	}
	if m.Block != 0 {
		i = encodeVarintDebugQuery(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x40
	}
	if m.ProviderRewardNoBonus != 0 {
		i = encodeVarintDebugQuery(dAtA, i, uint64(m.ProviderRewardNoBonus))
		i--
		dAtA[i] = 0x38
	}
	if m.ProviderFullReward != 0 {
		i = encodeVarintDebugQuery(dAtA, i, uint64(m.ProviderFullReward))
		i--
		dAtA[i] = 0x30
	}
	if m.ProviderAllocPoolBalance != 0 {
		i = encodeVarintDebugQuery(dAtA, i, uint64(m.ProviderAllocPoolBalance))
		i--
		dAtA[i] = 0x28
	}
	if m.ProviderDistPoolBalance != 0 {
		i = encodeVarintDebugQuery(dAtA, i, uint64(m.ProviderDistPoolBalance))
		i--
		dAtA[i] = 0x20
	}
	if m.ValAllocPoolBalance != 0 {
		i = encodeVarintDebugQuery(dAtA, i, uint64(m.ValAllocPoolBalance))
		i--
		dAtA[i] = 0x18
	}
	if m.ValDistPoolBalance != 0 {
		i = encodeVarintDebugQuery(dAtA, i, uint64(m.ValDistPoolBalance))
		i--
		dAtA[i] = 0x10
	}
	if m.BlockReward != 0 {
		i = encodeVarintDebugQuery(dAtA, i, uint64(m.BlockReward))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDebugQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovDebugQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryDebugQueryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockReward != 0 {
		n += 1 + sovDebugQuery(uint64(m.BlockReward))
	}
	if m.ValDistPoolBalance != 0 {
		n += 1 + sovDebugQuery(uint64(m.ValDistPoolBalance))
	}
	if m.ValAllocPoolBalance != 0 {
		n += 1 + sovDebugQuery(uint64(m.ValAllocPoolBalance))
	}
	if m.ProviderDistPoolBalance != 0 {
		n += 1 + sovDebugQuery(uint64(m.ProviderDistPoolBalance))
	}
	if m.ProviderAllocPoolBalance != 0 {
		n += 1 + sovDebugQuery(uint64(m.ProviderAllocPoolBalance))
	}
	if m.ProviderFullReward != 0 {
		n += 1 + sovDebugQuery(uint64(m.ProviderFullReward))
	}
	if m.ProviderRewardNoBonus != 0 {
		n += 1 + sovDebugQuery(uint64(m.ProviderRewardNoBonus))
	}
	if m.Block != 0 {
		n += 1 + sovDebugQuery(uint64(m.Block))
	}
	if m.MonthsLeft != 0 {
		n += 1 + sovDebugQuery(uint64(m.MonthsLeft))
	}
	if m.ValidatorReward != 0 {
		n += 1 + sovDebugQuery(uint64(m.ValidatorReward))
	}
	return n
}

func sovDebugQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDebugQuery(x uint64) (n int) {
	return sovDebugQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryDebugQueryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDebugQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryDebugQueryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDebugQueryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockReward", wireType)
			}
			m.BlockReward = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebugQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockReward |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValDistPoolBalance", wireType)
			}
			m.ValDistPoolBalance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebugQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValDistPoolBalance |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValAllocPoolBalance", wireType)
			}
			m.ValAllocPoolBalance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebugQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValAllocPoolBalance |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderDistPoolBalance", wireType)
			}
			m.ProviderDistPoolBalance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebugQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProviderDistPoolBalance |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderAllocPoolBalance", wireType)
			}
			m.ProviderAllocPoolBalance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebugQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProviderAllocPoolBalance |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderFullReward", wireType)
			}
			m.ProviderFullReward = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebugQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProviderFullReward |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderRewardNoBonus", wireType)
			}
			m.ProviderRewardNoBonus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebugQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProviderRewardNoBonus |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebugQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MonthsLeft", wireType)
			}
			m.MonthsLeft = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebugQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MonthsLeft |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorReward", wireType)
			}
			m.ValidatorReward = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebugQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidatorReward |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDebugQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDebugQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDebugQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDebugQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDebugQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDebugQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthDebugQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDebugQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDebugQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDebugQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDebugQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDebugQuery = fmt.Errorf("proto: unexpected end of group")
)