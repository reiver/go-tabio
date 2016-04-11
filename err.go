package tabio


import (
	"errors"
)


var (
	errInternalError  = errors.New("Internal Error")
	errNilReceiver    = internalNilReceiverComplainer{}
)
