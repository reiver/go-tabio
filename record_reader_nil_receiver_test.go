package tabio


import (
	"testing"
)


func TestRecordReaderNilReceiverClose(t *testing.T) {

	rr := (*RecordReader)(nil)

	err := rr.Close()

	if nil == err {
		t.Errorf("Expected an error but did not get one: %v", err)
		return
	}

	if _, ok := err.( interface{ RuntimeError() } ); !ok {
		t.Errorf("Expected a runetime error but did not get one.")
		return
	}

	if _, ok := err.( interface{ BadRequestComplainer() } ); !ok {
		t.Errorf("Expected a bad request complainer error but did not get one.")
		return
	}

	if _, ok := err.( interface{ NilReceiverComplainer() } ); !ok {
		t.Errorf("Expected a nil receiver complainer error but did not get one.")
		return
	}
}


func TestRecordReaderNilReceiverColumns(t *testing.T) {

	rr := (*RecordReader)(nil)

	_, err := rr.Columns()

	if nil == err {
		t.Errorf("Expected an error but did not get one: %v", err)
		return
	}

	if _, ok := err.( interface{ RuntimeError() } ); !ok {
		t.Errorf("Expected a runetime error but did not get one.")
		return
	}

	if _, ok := err.( interface{ BadRequestComplainer() } ); !ok {
		t.Errorf("Expected a bad request complainer error but did not get one.")
		return
	}

	if _, ok := err.( interface{ NilReceiverComplainer() } ); !ok {
		t.Errorf("Expected a nil receiver complainer error but did not get one.")
		return
	}
}


func TestRecordReaderNilReceiverErr(t *testing.T) {

	rr := (*RecordReader)(nil)

	err := rr.Err()

	if nil == err {
		t.Errorf("Expected an error but did not get one: %v", err)
		return
	}

	if _, ok := err.( interface{ RuntimeError() } ); !ok {
		t.Errorf("Expected a runetime error but did not get one.")
		return
	}

	if _, ok := err.( interface{ BadRequestComplainer() } ); !ok {
		t.Errorf("Expected a bad request complainer error but did not get one.")
		return
	}

	if _, ok := err.( interface{ NilReceiverComplainer() } ); !ok {
		t.Errorf("Expected a nil receiver complainer error but did not get one.")
		return
	}
}


func TestRecordReaderNilReceiverFields(t *testing.T) {

	rr := (*RecordReader)(nil)

	_, err := rr.Fields()

	if nil == err {
		t.Errorf("Expected an error but did not get one: %v", err)
		return
	}

	if _, ok := err.( interface{ RuntimeError() } ); !ok {
		t.Errorf("Expected a runetime error but did not get one.")
		return
	}

	if _, ok := err.( interface{ BadRequestComplainer() } ); !ok {
		t.Errorf("Expected a bad request complainer error but did not get one.")
		return
	}

	if _, ok := err.( interface{ NilReceiverComplainer() } ); !ok {
		t.Errorf("Expected a nil receiver complainer error but did not get one.")
		return
	}
}


func TestRecordReaderNilReceiverNext(t *testing.T) {

	defer func() {
		if r := recover(); nil != r {

			err, ok := r.(error)
			if !ok {
				t.Errorf("Expected an error but did not get one.")
				return
			}

			if _, ok := err.( interface{ RuntimeError() } ); !ok {
				t.Errorf("Expected a runetime error but did not get one.")
				return
			}

			if _, ok := err.( interface{ BadRequestComplainer() } ); !ok {
				t.Errorf("Expected a bad request complainer error but did not get one.")
				return
			}

			if _, ok := err.( interface{ NilReceiverComplainer() } ); !ok {
				t.Errorf("Expected a nil receiver complainer error but did not get one.")
				return
			}	
		}
	}()

	rr := (*RecordReader)(nil)

	_ = rr.Next()

	t.Errorf("This should never get here!")
}


func TestRecordReaderNilReceiverScan(t *testing.T) {

	rr := (*RecordReader)(nil)

	err := rr.Scan()

	if nil == err {
		t.Errorf("Expected an error but did not get one: %v", err)
		return
	}

	if _, ok := err.( interface{ RuntimeError() } ); !ok {
		t.Errorf("Expected a runetime error but did not get one.")
		return
	}

	if _, ok := err.( interface{ BadRequestComplainer() } ); !ok {
		t.Errorf("Expected a bad request complainer error but did not get one.")
		return
	}

	if _, ok := err.( interface{ NilReceiverComplainer() } ); !ok {
		t.Errorf("Expected a nil receiver complainer error but did not get one.")
		return
	}
}


func TestRecordReaderNilReceiverUnmarshal(t *testing.T) {

	rr := (*RecordReader)(nil)

	err := rr.Unmarshal([]interface{}{})

	if nil == err {
		t.Errorf("Expected an error but did not get one: %v", err)
		return
	}

	if _, ok := err.( interface{ RuntimeError() } ); !ok {
		t.Errorf("Expected a runetime error but did not get one.")
		return
	}

	if _, ok := err.( interface{ BadRequestComplainer() } ); !ok {
		t.Errorf("Expected a bad request complainer error but did not get one.")
		return
	}

	if _, ok := err.( interface{ NilReceiverComplainer() } ); !ok {
		t.Errorf("Expected a nil receiver complainer error but did not get one.")
		return
	}
}
