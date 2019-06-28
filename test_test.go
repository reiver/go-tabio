package tabio_test

import (
	"fmt"
	"strconv"
)

func stringAt(value interface{}) string {

	switch casted := value.(type) {
	case *fmt.Stringer:
		return fmt.Sprintf("%s", *casted)
	case *bool:
		return fmt.Sprintf("%t", *casted)
	case *complex64:
		return fmt.Sprintf("%v", *casted)
	case *complex128:
		return fmt.Sprintf("%v", *casted)
	case *float32:
		return fmt.Sprintf("%f", *casted)
	case *float64:
		return strconv.FormatFloat(*casted, 'f', -1, 64)
	case *int8:
		return fmt.Sprintf("%d", *casted)
	case *int16:
		return fmt.Sprintf("%d", *casted)
	case *int32:
		return fmt.Sprintf("%d", *casted)
	case *int64:
		return fmt.Sprintf("%d", *casted)
	case *string:
		return *casted
	case *uint8:
		return fmt.Sprintf("%d", *casted)
	case *uint16:
		return fmt.Sprintf("%d", *casted)
	case *uint32:
		return fmt.Sprintf("%d", *casted)
	case *uint64:
		return fmt.Sprintf("%d", *casted)
	default:
		return fmt.Sprintf("ERROR ERROR ERROR (%T)", casted)
	}

}
