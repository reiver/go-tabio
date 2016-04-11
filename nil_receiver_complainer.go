package tabio


type internalNilReceiverComplainer struct{}


func (internalNilReceiverComplainer) Error() string {
	return "Nil Receiver"
}


func (internalNilReceiverComplainer) RuntimeError() {
	// Nothing here.
}


func (internalNilReceiverComplainer) BadRequestComplainer() {
	// Nothing here.
}


func (internalNilReceiverComplainer) NilReceiverComplainer() {
	// Nothing here.
}
