package tabio


import (
	"errors"
)


var (
	errNilReceiver   = internalNilReceiverComplainer{}
	errInternalError = errors.New("Internal Error")
)
