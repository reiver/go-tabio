package tabio


import (
	"io/ioutil"
	"strings"

	"testing"
)



func TestRecordReaderClose(t *testing.T) {

	tests := []struct{
		String string
	}{
		{
			String: "",
		},
		{
			String: "\x1f",
		},



		{
			String: "\x1e",
		},
		{
			String: "\x1e\x1f",
		},



		{
			String: "\x1e\x1e",
		},
		{
			String: "\x1e\x1e\x1f",
		},



		{
			String: "\x1e\x1e\x1f\x1e\x1e\x1f\x1e\x1e",
		},
		{
			String: "\x1e\x1e\x1f\x1e\x1e\x1f\x1e\x1e\x1f",
		},



		{
			String: "apple banana cherry",
		},
		{
			String: "apple banana cherry\x1f",
		},



		{
			String: "fruit\x1fdruit",
		},
		{
			String: "fruit\x1fdruit\x1f",
		},



		{
			String: "fruit\x1fapple\x1fbanana\x1fcherry",
		},
		{
			String: "fruit\x1fapple\x1fbanana\x1fcherry\x1f",
		},



		{
			String: "apple\x1ebanana\x1echerry\x1fapple\x1ebanana\x1echerry",
		},
		{
			String: "apple\x1ebanana\x1echerry\x1fapple\x1ebanana\x1echerry\x1f",
		},



		{
			String: "given_name\x1efamily_name\x1ecity\x1fJoe\x1eBlow\x1eVancouver\x1fJane\x1eDoe\x1eSurrey",
		},
		{
			String: "given_name\x1efamily_name\x1ecity\x1fJoe\x1eBlow\x1eVancouver\x1fJane\x1eDoe\x1eSurrey\x1f",
		},
	}


	for testNumber, test := range tests {
		reader := ioutil.NopCloser( strings.NewReader(test.String) )

	        rr := NewRecordReader(reader)

		if err := rr.Close(); nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one: %v", testNumber, err)
			return
		}
	}
}



func TestRecordReaderCloseThenNext(t *testing.T) {

	tests := []struct{
		String string
	}{
		{
			String: "",
		},
		{
			String: "\x1f",
		},



		{
			String: "\x1e",
		},
		{
			String: "\x1e\x1f",
		},



		{
			String: "\x1e\x1e",
		},
		{
			String: "\x1e\x1e\x1f",
		},



		{
			String: "\x1e\x1e\x1f\x1e\x1e\x1f\x1e\x1e",
		},
		{
			String: "\x1e\x1e\x1f\x1e\x1e\x1f\x1e\x1e\x1f",
		},



		{
			String: "apple banana cherry",
		},
		{
			String: "apple banana cherry\x1f",
		},



		{
			String: "fruit\x1fdruit",
		},
		{
			String: "fruit\x1fdruit\x1f",
		},



		{
			String: "fruit\x1fapple\x1fbanana\x1fcherry",
		},
		{
			String: "fruit\x1fapple\x1fbanana\x1fcherry\x1f",
		},



		{
			String: "apple\x1ebanana\x1echerry\x1fapple\x1ebanana\x1echerry",
		},
		{
			String: "apple\x1ebanana\x1echerry\x1fapple\x1ebanana\x1echerry\x1f",
		},



		{
			String: "given_name\x1efamily_name\x1ecity\x1fJoe\x1eBlow\x1eVancouver\x1fJane\x1eDoe\x1eSurrey",
		},
		{
			String: "given_name\x1efamily_name\x1ecity\x1fJoe\x1eBlow\x1eVancouver\x1fJane\x1eDoe\x1eSurrey\x1f",
		},
	}


	for testNumber, test := range tests {

		reader := ioutil.NopCloser( strings.NewReader(test.String) )

	        rr := NewRecordReader(reader)

		rr.MustClose()

		next := rr.Next()
		if expected, actual := false, next; expected != actual {
			t.Errorf("For test #%d, expected %t but actually got %t.", testNumber, expected, actual)
			return
		}
	}
}
