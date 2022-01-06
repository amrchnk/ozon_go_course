package models

import (
	"errors"
	"strings"
)

type CallbackBody struct {
	CallbackName string
	CallbackData string
}

var ErrUnknownCallback = errors.New("unknown callback.go")

func ParseCallback(callbackData string) (CallbackBody, error) {
	callbackParts := strings.SplitN(callbackData, "_",2)
	if len(callbackParts) != 2 {
		return CallbackBody{}, ErrUnknownCallback
	}

	return CallbackBody{
		CallbackName: callbackParts[0],
		CallbackData: callbackParts[1],
	}, nil
}
