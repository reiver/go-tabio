package tabio


type internalClosedComplainer struct{}


func (internalClosedComplainer) Error() string {
	return "Closed"
}


func (internalClosedComplainer) RuntimeError() {
	// Nothing here.
}


func (internalClosedComplainer) ClosedComplainer() {
	// Nothing here.
}
