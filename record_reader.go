package tabio


import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)


// RecordReader implement reading of tabular data, in plain text format, where
// records are separated by a "record separator" (RS) ASCII and UNICODE control
// character and fields are separated by a "unit separator" (US) ASCII and UNICODE
// control character.
//
// (RS = 30. US = 31)
//
// *RecordReader has intentionally been made to include the methods of *sql.Rows.
//
// Because of this, if desired, a *RecordReader can be conceptually "converted"
// into a *sql.Rows by using the "github.com/reiver/go-shunt" package's shunt.Rows
// func.
type RecordReader struct {
	reader io.ReadCloser
	runeReader io.RuneReader

	buffer bytes.Buffer
	err error

	cachedColumns []string

	closed bool
}


// NewRecordReader returns a new *RecordReader, reading from the io.ReadCloser r.
func NewRecordReader(r io.ReadCloser) *RecordReader {
	runeReader, ok := r.(io.RuneReader)
	if !ok {
		runeReader = bufio.NewReader(r)
	}

	rr := RecordReader{
		reader:     r,
		runeReader: runeReader,
	}


	if !rr.Next() {
		return &rr
	}

	cachedColumns, err := rr.Fields()
	if nil != err {
		rr.err = err
		return &rr
	}

	rr.cachedColumns = cachedColumns


	return &rr
}


// Close closes the *RecordReader, preventing further iteration.
//
// If Next returns false then the *RecordReader is closed automatically.
//
// Close is idempotent and does not affect the result of Err.
//
// The first time Close is called it will in turn call Close on the wrapped io.ReadCloser that way passed to NewRecordReader.
func (rr *RecordReader) Close() error {
	if rr.closed {
		return nil
	}

	if err := rr.reader.Close(); nil != err {
		return err
	}

	rr.closed = true

	return nil
}


// MustClose is like Close, except it panic()s if there is an error.
func (rr *RecordReader) MustClose() {
	if err := rr.Close(); nil != err {
		panic(err)
	}
}


// Columns returns the column names.
//
// Columns returns an error if the *RecordReader is closed.
func (rr *RecordReader) Columns() ([]string, error) {
	if nil == rr {
		return nil, errNilReceiver
	}

	if rr.closed {
		return nil, errClosed
	}

	cachedColumns := rr.cachedColumns
	if nil == cachedColumns {
		return []string{}, nil
	}

//@TODO: Should we return a copy of this instead?
	return cachedColumns, nil
}


// MustColumns is like Columns, except it panic()s if there is an error.
func (rr *RecordReader) MustColumns() []string {
	cols, err := rr.Columns()
	if nil != err {
		panic(err)
	}

	return cols
}


// Error return an err, if an error occurred when calling Next.
//
// Usually Err is called after getting back a false result from a call to Next.
func (rr *RecordReader) Err() error {
	return rr.err
}


// Fields returns the field values.
//
// Fields returns an error if the *RecordReader is closed.
func (rr *RecordReader) Fields() ([]string, error) {
	const US = "\x1f" // Unit Separator

	if nil == rr {
		return nil, errNilReceiver
	}

	if rr.closed {
		return nil, errClosed
	}

	return strings.Split(rr.buffer.String(), US), nil
}


// MustFields is like Fields, except it panic()s if there is an error.
func (rr *RecordReader) MustFields() []string {
	fields, err := rr.Fields()
	if nil != err {
		panic(err)
	}

	return fields
}


// Next prepares the next record for reading with the Fields, Scan or Unmarshal methods.
//
// It returns true if it was successful.
//
// It returns false if there is no next record or if an error occurred while trying to prepare it.
//
// Err should be consulted to be able to tell the difference between the two cases.
//
// Every call to Fields, Scan, and Unmarshal, even the first one, must be preceded by a call to Next.
func (rr *RecordReader) Next() bool {

	const RS rune = 30 // Record Separator

	if nil == rr {
		panic(errNilReceiver)
	}

	if rr.closed {
		return false
	}

	if nil != rr.err {
		return false
	}


	runeReader := rr.runeReader
	if nil == runeReader {
		panic(errInternalError)
	}


	rr.buffer.Reset()


	var r rune
	var err error
	for {
		r, _, err = runeReader.ReadRune()
		if io.EOF == err {
			next := 0 < rr.buffer.Len()
			if !next {
				rr.Close()
			}

			return next
		}
		if nil != err {
			rr.err = err
			rr.Close()
			return false
		}

		if RS == r {
			return true
		}

		rr.buffer.WriteRune(r)
	}
}


// Scan copies the fields in the current record into the values pointed at by dest.
//
// The number of values in dest must be the same as the number of columns in *RecordReader.
func (rr *RecordReader) Scan(dest ...interface{}) error {
	if nil == rr {
		return errNilReceiver
	}

	lenColumns := 0
	if cachedColumns := rr.cachedColumns; nil != cachedColumns {
		lenColumns = len(cachedColumns)
	}

	if lenDest := len(dest); lenColumns != lenDest {
		return internalWrongNumberOfArgumentsComplainer{expected: uint(lenColumns), actual: uint(lenDest) }
	}


//@TODO: Do this better.
	fields, err := rr.Fields()
	if nil != err {
		return err
	}
	for i, _ := range dest {
//@TODO: need to handle converions... probably.
		switch dest[i].(type) {
		case *string:
			dest[i] = &fields[i]
		default:
			return internalUnsupportedScanComplainer{
				srcType:  fmt.Sprintf("%T", &fields[i]),
				destType: fmt.Sprintf("%T", dest[i]),
			}
		}
	}


	return nil
}

// MustScan is like Scan, except it panic()s if there is an error.
func (rr *RecordReader) MustScan(dest ...interface{}) {
	if err := rr.Scan(dest...); nil != err {
		panic(err)
	}
}

func (rr *RecordReader) Unmarshal(target interface{}) error {
	if nil == rr {
		return errNilReceiver
	}


	switch x:= target.(type) {
	case RecordUnmarshaler:
		bs := rr.buffer.Bytes()

		if err := x.UnmarshalRecord(bs); nil != err {
			return err
		}
	default:
panic("//@TODO TODO TODO TODO TODO")

	}


	return nil
}

// MustUnmarshal is like Unmarshal, except it panic()s if there is an error.
func (rr *RecordReader) MustUnmarshal(target interface{}) {
	if err := rr.Unmarshal(target); nil != err {
		panic(err)
	}
}
