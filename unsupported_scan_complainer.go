package tabio


import (
	"fmt"
)


type internalUnsupportedScanComplainer struct {
	destType string
	srcType  string
}


func (err internalUnsupportedScanComplainer) Error() string {
	return fmt.Sprintf("Unsupported scan from %q to %q.", err.srcType, err.destType)
}


func (internalUnsupportedScanComplainer) RuntimeError() {
        // Nothing here.
}


func (internalUnsupportedScanComplainer) BadRequestComplainer() {
	// Nothing here.
}
