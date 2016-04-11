package tabio


import (
	"io/ioutil"
	"strings"

	"testing"
)



func TestRecordReaderClose(t *testing.T) {

	reader := ioutil.NopCloser( strings.NewReader("apple banana cherry") )

        rr := NewRecordReader(reader)

	if err := rr.Close(); nil != err {
		t.Errorf("Did not expect an error but actually got one: %v", err)
		return
	}
}



func TestRecordReaderCloseThenNext(t *testing.T) {

	reader := ioutil.NopCloser( strings.NewReader("apple\x1ebanana\x1echerry\x1fapple\x1ebanana\x1echerry\x1f") )

        rr := NewRecordReader(reader)

	rr.MustClose()

	next := rr.Next()
	if expected, actual := false, next; expected != actual {
		t.Errorf("Expected %t but actually got %t.", expected, actual)
		return
	}
}
