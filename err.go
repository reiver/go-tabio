package tabio


import (
	"errors"
)


var (
	errClosed         = internalClosedComplainer{}
	errInternalError  = errors.New("Internal Error")
	errNilReceiver    = internalNilReceiverComplainer{}
)
