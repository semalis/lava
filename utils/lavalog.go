package utils

import (
	"errors"
	"fmt"
	"os"

	sdk "github.com/cosmos/cosmos-sdk/types"
	zerolog "github.com/rs/zerolog"
	zerologlog "github.com/rs/zerolog/log"
	"github.com/tendermint/tendermint/libs/log"
)

const (
	EventPrefix = "lava_"
)

func LogLavaEvent(ctx sdk.Context, logger log.Logger, name string, attributes map[string]string, description string) {
	attributes_str := ""
	eventAttrs := []sdk.Attribute{}
	for key, val := range attributes {
		attributes_str += fmt.Sprintf("%s: %s,", key, val)
		eventAttrs = append(eventAttrs, sdk.NewAttribute(key, val))
	}
	logger.Info(fmt.Sprintf("%s%s: %s %s", EventPrefix, name, description, attributes_str))
	ctx.EventManager().EmitEvent(sdk.NewEvent(EventPrefix+name, eventAttrs...))
}

func LavaError(ctx sdk.Context, logger log.Logger, name string, attributes map[string]string, description string) error {
	attributes_str := ""
	eventAttrs := []sdk.Attribute{}
	for key, val := range attributes {
		attributes_str += fmt.Sprintf("%s: %s,", key, val)
		eventAttrs = append(eventAttrs, sdk.NewAttribute(key, val))
	}
	err_msg := fmt.Sprintf("ERR_%s: %s %s", name, description, attributes_str)
	logger.Error(err_msg)
	// ctx.EventManager().EmitEvent(sdk.NewEvent("ERR_"+name, eventAttrs...))
	//TODO: add error types, create them here and return
	return errors.New(err_msg)
}

func LavaFormatLog(description string, err error, extraAttributes *map[string]string, severity uint) error {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	var logEvent *zerolog.Event
	switch severity {
	case 4:
		// prefix = "Fatal:"
		logEvent = zerologlog.Fatal()

	case 3:
		// prefix = "Error:"
		logEvent = zerologlog.Error()
	case 2:
		// prefix = "Warning:"
		logEvent = zerologlog.Warn()
	case 1:
		logEvent = zerologlog.Info()
		// prefix = "Info:"
	case 0:
		logEvent = zerologlog.Debug()
		// prefix = "Debug:"
	}
	output := description
	if err != nil {
		logEvent = logEvent.Err(err)
		output = fmt.Sprintf("%s ErrMsg: %s", output, err.Error())
	}
	if extraAttributes != nil {
		for key, val := range *extraAttributes {
			logEvent = logEvent.Str(key, val)
		}
		output = fmt.Sprintf("%s -- %v", output, extraAttributes)
	}
	zerologlog.Logger = zerologlog.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	logEvent.Msg(description)
	// golog.Println(output)
	return fmt.Errorf(output)
}

func LavaFormatFatal(description string, err error, extraAttributes *map[string]string) {
	LavaFormatLog(description, err, extraAttributes, 4)
	os.Exit(1)
}

func LavaFormatError(description string, err error, extraAttributes *map[string]string) error {
	return LavaFormatLog(description, err, extraAttributes, 3)
}

func LavaFormatWarning(description string, err error, extraAttributes *map[string]string) error {
	return LavaFormatLog(description, err, extraAttributes, 2)
}

func LavaFormatInfo(description string, extraAttributes *map[string]string) error {
	return LavaFormatLog(description, nil, extraAttributes, 1)
}

func LavaFormatDebug(description string, extraAttributes *map[string]string) error {
	return LavaFormatLog(description, nil, extraAttributes, 0)
}