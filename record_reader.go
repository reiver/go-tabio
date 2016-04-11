package tabio


import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)


type RecordReader struct {
	reader io.ReadCloser
	runeReader io.RuneReader

	buffer bytes.Buffer
	err error

	cachedColumns []string

	closed bool
}


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

	return rr.reader.Close()
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


func (rr *RecordReader) MustColumns() []string {
	cols, err := rr.Columns()
	if nil != err {
		panic(err)
	}

	return cols
}


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


func (rr *RecordReader) MustFields() []string {
	fields, err := rr.Fields()
	if nil != err {
		panic(err)
	}

	return fields
}


func (rr *RecordReader) Next() bool {

	const RS rune = 30 // Record Separator

	if nil == rr {
		panic(errNilReceiver)
	}

	if rr.closed {
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

func (rr *RecordReader) MustUnmarshal(target interface{}) {
	if err := rr.Unmarshal(target); nil != err {
		panic(err)
	}
}
