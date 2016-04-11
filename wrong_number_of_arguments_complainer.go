package tabio


import (
	"fmt"
)


type internalWrongNumberOfArgumentsComplainer struct {
	expected uint
	actual   uint
}


func (err internalWrongNumberOfArgumentsComplainer) Error() string {
	return fmt.Sprintf("Expected %d arguments, not %d.", err.expected, err.actual)
}


func (internalWrongNumberOfArgumentsComplainer) RuntimeError() {
        // Nothing here.
}


func (internalWrongNumberOfArgumentsComplainer) BadRequestComplainer() {
	// Nothing here.
}
