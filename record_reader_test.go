package tabio_test


import (
	"github.com/reiver/go-tabio"

	"io/ioutil"
	"strings"

	"testing"
)


func TestNewRecordReader(t *testing.T) {
	reader := ioutil.NopCloser( strings.NewReader("apple banana cherry") )

	rr := tabio.NewRecordReader(reader)
	if nil == rr {
		t.Errorf("Did not expect nil but actually got: %v", rr)
		return
	}
}


func TestRecordReaderColumns(t *testing.T) {

	const RS string = "\x1e"
	const US string = "\x1f"

	tests := []struct{
		String     string
		Expected []string
	}{
		{
			String: "given_name" +US+ "family_name" +US+ "city"      +RS+
			        "Joe"        +US+ "Blow"        +US+ "Vancouver" +RS+
			        "Jane"       +US+ "Doe"         +US+ "Surrey"    +RS,

			Expected: []string{"given_name", "family_name", "city"},
		},
		{
			String: "given_name" +US+ "family_name" +US+ "city"      +RS+
			        "Joe"        +US+ "Blow"        +US+ "Vancouver" +RS+
			        "Jane"       +US+ "Doe"         +US+ "Surrey",

			Expected: []string{"given_name", "family_name", "city"},
		},



		{
			String: "given_name" +US+ "family_name" +US+ "city"      +RS,

			Expected: []string{"given_name", "family_name", "city"},
		},
		{
			String: "given_name" +US+ "family_name" +US+ "city",

			Expected: []string{"given_name", "family_name", "city"},
		},



		{
			String: "fruit"  +RS+
			        "apple"  +RS+
			        "banana" +RS+
			        "cherry" +RS,

			Expected: []string{"fruit"},
		},
		{
			String: "fruit"  +RS+
			        "apple"  +RS+
			        "banana" +RS+
			        "cherry",

			Expected: []string{"fruit"},
		},



		{
			String: "fruit"  +RS,

			Expected: []string{"fruit"},
		},
		{
			String: "fruit",

			Expected: []string{"fruit"},
		},
	}


	TestLoop: for testNumber, test := range tests {

		rr := tabio.NewRecordReader( ioutil.NopCloser(strings.NewReader(test.String)) )

		columns := rr.MustColumns()

		if expected, actual := len(test.Expected), len(columns); expected != actual {
			t.Errorf("For test #%d, expected %d columns but actually got %d columns; expected columns = %#v; actual columns = %#v.", testNumber, expected, actual, test.Expected, columns)
			continue
		}

		for columnNumber, expected := range test.Expected {
			actual := columns[columnNumber]

			if expected != actual {
				t.Errorf("For test #%d and column #%d, did not expected %q but actually got %q.", testNumber, columnNumber, expected, actual)
				continue TestLoop
			}
		}
	}
}


func TestRecordReaderNext(t *testing.T) {

	const RS string = "\x1e"
	const US string = "\x1f"

	tests := []struct{
		String   string
		Expected int
	}{
		{
			String: "",
			Expected: 0,
		},
		{
			String: "" +RS,
			Expected: 0,
		},



		{
			String: "" +US+ "",
			Expected: 0,
		},
		{
			String: "" +US+ "" +RS,
			Expected: 0,
		},



		{
			String: "" +US+ "" +US+ "",
			Expected: 0,
		},
		{
			String: "" +US+ "" +US+ "" +RS,
			Expected: 0,
		},



		{
			String: "" +RS+
			        "",
			Expected: 0,
		},
		{
			String: "" +RS+
			        "" +RS,
			Expected: 1,
		},



		{
			String: "" +US+ "" +RS+
			        "" +US+ "",
			Expected: 1,
		},
		{
			String: "" +US+ "" +RS+
			        "" +US+ "" +RS,
			Expected: 1,
		},



		{
			String: "" +US+ "" +US+ "" +RS+
			        "" +US+ "" +US+ "",
			Expected: 1,
		},
		{
			String: "" +US+ "" +US+ "" +RS+
			        "" +US+ "" +US+ "" +RS,
			Expected: 1,
		},



		{
			String: "" +RS+
			        "" +RS+
			        "",
			Expected: 1,
		},
		{
			String: "" +RS+
			        "" +RS+
			        "" +RS,
			Expected: 2,
		},



		{
			String: "" +US+ "" +RS+
			        "" +US+ "" +RS+
			        "" +US+ "",
			Expected: 2,
		},
		{
			String: "" +US+ "" +RS+
			        "" +US+ "" +RS+
			        "" +US+ "" +RS,
			Expected: 2,
		},



		{
			String: "" +US+ "" +US+ "" +RS+
			        "" +US+ "" +US+ "" +RS+
			        "" +US+ "" +US+ "",
			Expected: 2,
		},
		{
			String: "" +US+ "" +US+ "" +RS+
			        "" +US+ "" +US+ "" +RS+
			        "" +US+ "" +US+ "" +RS,
			Expected: 2,
		},



		{
			String: "given_name" +US+ "family_name" +US+ "city"      +RS+
			        "Joe"        +US+ "Blow"        +US+ "Vancouver" +RS+
			        "Jane"       +US+ "Doe"         +US+ "Surrey"    +RS,

			Expected: 2,
		},
		{
			String: "given_name" +US+ "family_name" +US+ "city"      +RS+
			        "Joe"        +US+ "Blow"        +US+ "Vancouver" +RS+
			        "Jane"       +US+ "Doe"         +US+ "Surrey",

			Expected: 2,
		},



		{
			String: "given_name" +US+ "family_name" +US+ "city"      +RS,

			Expected: 0,
		},
		{
			String: "given_name" +US+ "family_name" +US+ "city",

			Expected: 0,
		},



		{
			String: "fruit"  +RS+
			        "apple"  +RS+
			        "banana" +RS+
			        "cherry",

			Expected: 3,
		},
		{
			String: "fruit"  +RS+
			        "apple"  +RS+
			        "banana" +RS+
			        "cherry" +RS,

			Expected: 3,
		},



		{
			String: "fruit",

			Expected: 0,
		},
		{
			String: "fruit"  +RS,

			Expected: 0,
		},



		{
			String: "fruit"  +RS+
			        "apple",

			Expected: 1,
		},
		{
			String: "fruit"  +RS+
			        "apple"  +RS,

			Expected: 1,
		},



		{
			String: "fruit"  +RS+
			        "apple"  +RS+
			        "banana",

			Expected: 2,
		},
		{
			String: "fruit"  +RS+
			        "apple"  +RS+
			        "banana" +RS,

			Expected: 2,
		},



		{
			String: "fruit",

			Expected: 0,
		},
		{
			String: "fruit"  +RS,

			Expected: 0,
		},



		{
			String: "fruit" +RS+
			        "fruit",

			Expected: 1,
		},
		{
			String: "fruit" +RS+
			        "fruit" +RS,

			Expected: 1,
		},



		{
			String: "fruit" +RS+
			        "fruit" +RS+
			        "fruit",

			Expected: 2,
		},
		{
			String: "fruit" +RS+
			        "fruit" +RS+
			        "fruit" +RS,

			Expected: 2,
		},
	}


	TestLoop: for testNumber, test := range tests {

		rr := tabio.NewRecordReader( ioutil.NopCloser( strings.NewReader(test.String) ) )

		count := 0
		for rr.Next() {
			count++

			if expected, actual := test.Expected, count; expected < actual {
				t.Errorf("For test #%d, expected count to be less than or equal to %d, but became %d; for string = %q.", testNumber, expected, actual, test.String)
				continue TestLoop
			}
		}

		if expected, actual := test.Expected, count; expected != actual {
			t.Errorf("For test #%d, expected %d but actually got %d; for string = %q.", testNumber, expected, actual, test.String)
			continue
		}
	}
}


func TestRecordReaderFields(t *testing.T) {

	const RS string = "\x1e"
	const US string = "\x1f"

	tests := []struct{
		String     string
		Expected [][]string
	}{
		{
			String: "given_name" +US+ "family_name" +US+ "city"      +RS+
			        "Joe"        +US+ "Blow"        +US+ "Vancouver" +RS+
			        "Jane"       +US+ "Doe"         +US+ "Surrey"    +RS,

			Expected: [][]string{
				[]string{"Joe",  "Blow", "Vancouver"},
				[]string{"Jane", "Doe",  "Surrey"},
			},
		},
		{
			String: "given_name" +US+ "family_name" +US+ "city"      +RS+
			        "Joe"        +US+ "Blow"        +US+ "Vancouver" +RS+
			        "Jane"       +US+ "Doe"         +US+ "Surrey",

			Expected: [][]string{
				[]string{"Joe",  "Blow", "Vancouver"},
				[]string{"Jane", "Doe",  "Surrey"},
			},
		},



		{
			String: "given_name" +US+ "family_name" +US+ "city"      +RS,

			Expected: [][]string{},
		},
		{
			String: "given_name" +US+ "family_name" +US+ "city",

			Expected: [][]string{},
		},



		{
			String: "fruit"  +RS+
			        "apple"  +RS+
			        "banana" +RS+
			        "cherry" +RS,

			Expected: [][]string{
				[]string{"apple"},
				[]string{"banana"},
				[]string{"cherry"},
			},
		},
		{
			String: "fruit"  +RS+
			        "apple"  +RS+
			        "banana" +RS+
			        "cherry",

			Expected: [][]string{
				[]string{"apple"},
				[]string{"banana"},
				[]string{"cherry"},
			},
		},



		{
			String: "fruit"  +RS,

			Expected: [][]string{},
		},
		{
			String: "fruit",

			Expected: [][]string{},
		},
	}


	TestLoop: for testNumber, test := range tests {

		rr := tabio.NewRecordReader( ioutil.NopCloser( strings.NewReader(test.String) ) )

		i := 0
		for rr.Next() {
			fields := rr.MustFields()
			if expected, actual := len(test.Expected[i]), len(fields); expected != actual {
				t.Errorf("For test #%d, expected %d fields but actually got %d.", testNumber, expected, actual)
				continue TestLoop
			}

			for fieldNumber, expected := range test.Expected[i] {
				if actual := fields[fieldNumber]; expected != actual {
					t.Errorf("For test #%d and field #%d, expected %q fields but actually got %q.", testNumber, fieldNumber, expected, actual)
					continue TestLoop
				}
			}

			i++
		}
	}
}


func TestRecordReaderScan(t *testing.T) {

	const RS string = "\x1e"
	const US string = "\x1f"

	tests := []struct{
		String     string
		Expected [][]string
		Dest     []interface{}
	}{
		{ // 0
			String: "given_name" +US+ "family_name" +US+ "city"      +RS+
			        "Joe"        +US+ "Blow"        +US+ "Vancouver" +RS+
			        "Jane"       +US+ "Doe"         +US+ "Surrey"    +RS,

			Expected: [][]string{
				[]string{"Joe",  "Blow", "Vancouver"},
				[]string{"Jane", "Doe",  "Surrey"},
			},
			Dest: []interface{}{ (*string)(nil), (*string)(nil), (*string)(nil) },
		},
		{ // 1
			String: "given_name" +US+ "family_name" +US+ "city"      +RS+
			        "Joe"        +US+ "Blow"        +US+ "Vancouver" +RS+
			        "Jane"       +US+ "Doe"         +US+ "Surrey",

			Expected: [][]string{
				[]string{"Joe",  "Blow", "Vancouver"},
				[]string{"Jane", "Doe",  "Surrey"},
			},
			Dest: []interface{}{ (*string)(nil), (*string)(nil), (*string)(nil) },
		},



		{ // 2
			String: "given_name" +US+ "family_name" +US+ "city"      +RS,

			Expected: [][]string{},
			Dest: []interface{}{ (*string)(nil), (*string)(nil), (*string)(nil) },
		},
		{ // 3
			String: "given_name" +US+ "family_name" +US+ "city",

			Expected: [][]string{},
			Dest: []interface{}{ (*string)(nil), (*string)(nil), (*string)(nil) },
		},



		{ // 4
			String: "fruit"  +RS+
			        "apple"  +RS+
			        "banana" +RS+
			        "cherry" +RS,

			Expected: [][]string{
				[]string{"apple"},
				[]string{"banana"},
				[]string{"cherry"},
			},
			Dest: []interface{}{ (*string)(nil) },
		},
		{ // 5
			String: "fruit"  +RS+
			        "apple"  +RS+
			        "banana" +RS+
			        "cherry",

			Expected: [][]string{
				[]string{"apple"},
				[]string{"banana"},
				[]string{"cherry"},
			},
			Dest: []interface{}{ (*string)(nil) },
		},



		{ // 6
			String: "fruit"  +RS,

			Expected: [][]string{},
			Dest: []interface{}{ (*string)(nil) },
		},
		{ // 7
			String: "fruit",

			Expected: [][]string{},
			Dest: []interface{}{ (*string)(nil) },
		},



		{ // 8
			String:  "x" +US+ "y" +RS+
			        "-2" +US+ "4" +RS+
			        "-1" +US+ "2" +RS+
			         "0" +US+ "0" +RS+
			         "1" +US+ "2" +RS+
			         "2" +US+ "4" +RS,

			Expected: [][]string{
				[]string{"-2", "4"},
				[]string{"-1", "2"},
				[]string{ "0", "0"},
				[]string{ "1", "2"},
				[]string{ "2", "4"},
			},
			Dest: []interface{}{ (*string)(nil), (*string)(nil) },
		},
		{ // 9
			String:  "x" +US+ "y" +RS+
			        "-2" +US+ "4" +RS+
			        "-1" +US+ "2" +RS+
			         "0" +US+ "0" +RS+
			         "1" +US+ "2" +RS+
			         "2" +US+ "4",

			Expected: [][]string{
				[]string{"-2", "4"},
				[]string{"-1", "2"},
				[]string{ "0", "0"},
				[]string{ "1", "2"},
				[]string{ "2", "4"},
			},
			Dest: []interface{}{ (*string)(nil), (*string)(nil) },
		},









		{
			String:  "x" +US+ "y" +RS+
			        "-2" +US+ "4" +RS+
			        "-1" +US+ "2" +RS+
			         "0" +US+ "0" +RS+
			         "1" +US+ "2" +RS+
			         "2" +US+ "4" +RS,

			Expected: [][]string{
				[]string{"-2", "4"},
				[]string{"-1", "2"},
				[]string{ "0", "0"},
				[]string{ "1", "2"},
				[]string{ "2", "4"},
			},
			Dest: []interface{}{ (*int64)(nil), (*int64)(nil) },
		},
		{
			String:  "x" +US+ "y" +RS+
			        "-2" +US+ "4" +RS+
			        "-1" +US+ "2" +RS+
			         "0" +US+ "0" +RS+
			         "1" +US+ "2" +RS+
			         "2" +US+ "4",

			Expected: [][]string{
				[]string{"-2", "4"},
				[]string{"-1", "2"},
				[]string{ "0", "0"},
				[]string{ "1", "2"},
				[]string{ "2", "4"},
			},
			Dest: []interface{}{ (*int64)(nil), (*int64)(nil) },
		},



		{
			String:  "x" +US+ "y" +RS+
			        "-2" +US+ "4" +RS+
			        "-1" +US+ "2" +RS+
			         "0" +US+ "0" +RS+
			         "1" +US+ "2" +RS+
			         "2" +US+ "4" +RS,

			Expected: [][]string{
				[]string{"-2", "4"},
				[]string{"-1", "2"},
				[]string{ "0", "0"},
				[]string{ "1", "2"},
				[]string{ "2", "4"},
			},
			Dest: []interface{}{ (*int32)(nil), (*int32)(nil) },
		},
		{
			String:  "x" +US+ "y" +RS+
			        "-2" +US+ "4" +RS+
			        "-1" +US+ "2" +RS+
			         "0" +US+ "0" +RS+
			         "1" +US+ "2" +RS+
			         "2" +US+ "4",

			Expected: [][]string{
				[]string{"-2", "4"},
				[]string{"-1", "2"},
				[]string{ "0", "0"},
				[]string{ "1", "2"},
				[]string{ "2", "4"},
			},
			Dest: []interface{}{ (*int32)(nil), (*int32)(nil) },
		},



		{
			String:  "x" +US+ "y" +RS+
			        "-2" +US+ "4" +RS+
			        "-1" +US+ "2" +RS+
			         "0" +US+ "0" +RS+
			         "1" +US+ "2" +RS+
			         "2" +US+ "4" +RS,

			Expected: [][]string{
				[]string{"-2", "4"},
				[]string{"-1", "2"},
				[]string{ "0", "0"},
				[]string{ "1", "2"},
				[]string{ "2", "4"},
			},
			Dest: []interface{}{ (*int16)(nil), (*int16)(nil) },
		},
		{
			String:  "x" +US+ "y" +RS+
			        "-2" +US+ "4" +RS+
			        "-1" +US+ "2" +RS+
			         "0" +US+ "0" +RS+
			         "1" +US+ "2" +RS+
			         "2" +US+ "4",

			Expected: [][]string{
				[]string{"-2", "4"},
				[]string{"-1", "2"},
				[]string{ "0", "0"},
				[]string{ "1", "2"},
				[]string{ "2", "4"},
			},
			Dest: []interface{}{ (*int16)(nil), (*int16)(nil) },
		},



		{
			String:  "x" +US+ "y" +RS+
			        "-2" +US+ "4" +RS+
			        "-1" +US+ "2" +RS+
			         "0" +US+ "0" +RS+
			         "1" +US+ "2" +RS+
			         "2" +US+ "4" +RS,

			Expected: [][]string{
				[]string{"-2", "4"},
				[]string{"-1", "2"},
				[]string{ "0", "0"},
				[]string{ "1", "2"},
				[]string{ "2", "4"},
			},
			Dest: []interface{}{ (*int8)(nil), (*int8)(nil) },
		},
		{
			String:  "x" +US+ "y" +RS+
			        "-2" +US+ "4" +RS+
			        "-1" +US+ "2" +RS+
			         "0" +US+ "0" +RS+
			         "1" +US+ "2" +RS+
			         "2" +US+ "4",

			Expected: [][]string{
				[]string{"-2", "4"},
				[]string{"-1", "2"},
				[]string{ "0", "0"},
				[]string{ "1", "2"},
				[]string{ "2", "4"},
			},
			Dest: []interface{}{ (*int8)(nil), (*int8)(nil) },
		},









		{
			String:  "x" +US+ "y" +RS+
			        "-2" +US+ "4" +RS+
			        "-1" +US+ "2" +RS+
			         "0" +US+ "0" +RS+
			         "1" +US+ "2" +RS+
			         "2" +US+ "4" +RS,

			Expected: [][]string{
				[]string{"-2", "4"},
				[]string{"-1", "2"},
				[]string{ "0", "0"},
				[]string{ "1", "2"},
				[]string{ "2", "4"},
			},
			Dest: []interface{}{ (*uint64)(nil), (*uint64)(nil) },
		},
		{
			String:  "x" +US+ "y" +RS+
			        "-2" +US+ "4" +RS+
			        "-1" +US+ "2" +RS+
			         "0" +US+ "0" +RS+
			         "1" +US+ "2" +RS+
			         "2" +US+ "4",

			Expected: [][]string{
				[]string{"-2", "4"},
				[]string{"-1", "2"},
				[]string{ "0", "0"},
				[]string{ "1", "2"},
				[]string{ "2", "4"},
			},
			Dest: []interface{}{ (*uint64)(nil), (*uint64)(nil) },
		},



		{
			String:  "x" +US+ "y" +RS+
			        "-2" +US+ "4" +RS+
			        "-1" +US+ "2" +RS+
			         "0" +US+ "0" +RS+
			         "1" +US+ "2" +RS+
			         "2" +US+ "4" +RS,

			Expected: [][]string{
				[]string{"-2", "4"},
				[]string{"-1", "2"},
				[]string{ "0", "0"},
				[]string{ "1", "2"},
				[]string{ "2", "4"},
			},
			Dest: []interface{}{ (*uint32)(nil), (*uint32)(nil) },
		},
		{
			String:  "x" +US+ "y" +RS+
			        "-2" +US+ "4" +RS+
			        "-1" +US+ "2" +RS+
			         "0" +US+ "0" +RS+
			         "1" +US+ "2" +RS+
			         "2" +US+ "4",

			Expected: [][]string{
				[]string{"-2", "4"},
				[]string{"-1", "2"},
				[]string{ "0", "0"},
				[]string{ "1", "2"},
				[]string{ "2", "4"},
			},
			Dest: []interface{}{ (*uint32)(nil), (*uint32)(nil) },
		},



		{
			String:  "x" +US+ "y" +RS+
			        "-2" +US+ "4" +RS+
			        "-1" +US+ "2" +RS+
			         "0" +US+ "0" +RS+
			         "1" +US+ "2" +RS+
			         "2" +US+ "4" +RS,

			Expected: [][]string{
				[]string{"-2", "4"},
				[]string{"-1", "2"},
				[]string{ "0", "0"},
				[]string{ "1", "2"},
				[]string{ "2", "4"},
			},
			Dest: []interface{}{ (*uint16)(nil), (*uint16)(nil) },
		},
		{
			String:  "x" +US+ "y" +RS+
			        "-2" +US+ "4" +RS+
			        "-1" +US+ "2" +RS+
			         "0" +US+ "0" +RS+
			         "1" +US+ "2" +RS+
			         "2" +US+ "4",

			Expected: [][]string{
				[]string{"-2", "4"},
				[]string{"-1", "2"},
				[]string{ "0", "0"},
				[]string{ "1", "2"},
				[]string{ "2", "4"},
			},
			Dest: []interface{}{ (*uint16)(nil), (*uint16)(nil) },
		},



		{
			String:  "x" +US+ "y" +RS+
			        "-2" +US+ "4" +RS+
			        "-1" +US+ "2" +RS+
			         "0" +US+ "0" +RS+
			         "1" +US+ "2" +RS+
			         "2" +US+ "4" +RS,

			Expected: [][]string{
				[]string{"-2", "4"},
				[]string{"-1", "2"},
				[]string{ "0", "0"},
				[]string{ "1", "2"},
				[]string{ "2", "4"},
			},
			Dest: []interface{}{ (*uint8)(nil), (*uint8)(nil) },
		},
		{
			String:  "x" +US+ "y" +RS+
			        "-2" +US+ "4" +RS+
			        "-1" +US+ "2" +RS+
			         "0" +US+ "0" +RS+
			         "1" +US+ "2" +RS+
			         "2" +US+ "4",

			Expected: [][]string{
				[]string{"-2", "4"},
				[]string{"-1", "2"},
				[]string{ "0", "0"},
				[]string{ "1", "2"},
				[]string{ "2", "4"},
			},
			Dest: []interface{}{ (*uint8)(nil), (*uint8)(nil) },
		},









		{
			String:  "x" +US+ "y" +RS+
			        "-2" +US+ "4" +RS+
			        "-1" +US+ "2" +RS+
			         "0" +US+ "0" +RS+
			         "1" +US+ "2" +RS+
			         "2" +US+ "4" +RS,

			Expected: [][]string{
				[]string{"-2", "4"},
				[]string{"-1", "2"},
				[]string{ "0", "0"},
				[]string{ "1", "2"},
				[]string{ "2", "4"},
			},
			Dest: []interface{}{ (*float32)(nil), (*float32)(nil) },
		},
		{
			String:  "x" +US+ "y" +RS+
			        "-2" +US+ "4" +RS+
			        "-1" +US+ "2" +RS+
			         "0" +US+ "0" +RS+
			         "1" +US+ "2" +RS+
			         "2" +US+ "4",

			Expected: [][]string{
				[]string{"-2", "4"},
				[]string{"-1", "2"},
				[]string{ "0", "0"},
				[]string{ "1", "2"},
				[]string{ "2", "4"},
			},
			Dest: []interface{}{ (*float32)(nil), (*float32)(nil) },
		},



		{
			String:  "x" +US+ "y" +RS+
			        "-2" +US+ "4" +RS+
			        "-1" +US+ "2" +RS+
			         "0" +US+ "0" +RS+
			         "1" +US+ "2" +RS+
			         "2" +US+ "4" +RS,

			Expected: [][]string{
				[]string{"-2", "4"},
				[]string{"-1", "2"},
				[]string{ "0", "0"},
				[]string{ "1", "2"},
				[]string{ "2", "4"},
			},
			Dest: []interface{}{ (*float64)(nil), (*float64)(nil) },
		},
		{
			String:  "x" +US+ "y" +RS+
			        "-2" +US+ "4" +RS+
			        "-1" +US+ "2" +RS+
			         "0" +US+ "0" +RS+
			         "1" +US+ "2" +RS+
			         "2" +US+ "4",

			Expected: [][]string{
				[]string{"-2", "4"},
				[]string{"-1", "2"},
				[]string{ "0", "0"},
				[]string{ "1", "2"},
				[]string{ "2", "4"},
			},
			Dest: []interface{}{ (*float64)(nil), (*float64)(nil) },
		},



		{
			String:  "x" +US+ "y" +RS,

			Expected: [][]string{},
			Dest: []interface{}{ (*string)(nil), (*string)(nil) },
		},
		{
			String:  "x" +US+ "y",

			Expected: [][]string{},
			Dest: []interface{}{ (*string)(nil), (*string)(nil) },
		},
	}


	TestLoop: for testNumber, test := range tests {

		rr := tabio.NewRecordReader( ioutil.NopCloser( strings.NewReader(test.String) ) )

		i := 0
		for rr.Next() {
			if err := rr.Scan(test.Dest...); nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one: %v", testNumber, err)
				t.Errorf("\tString: %q", test.String)
				t.Errorf("\tDest:...")
				for destNumber, dest := range test.Dest {
					t.Errorf("\t\tdest[%d] -> (%T)", destNumber, dest)
				}
				continue TestLoop
			}

			if expected, actual := len(test.Expected[i]), len(test.Dest); expected != actual {
				t.Errorf("For test #%d, expected %d fields but actually got %d.", testNumber, expected, actual)
				continue TestLoop
			}

			for fieldNumber, expected := range test.Expected[i] {
				if actual := test.Dest[fieldNumber]; expected == actual {
					t.Errorf("For test #%d and field #%d, expected (%T) %q but actually got (%T) %q.", testNumber, fieldNumber, expected, expected, actual, actual)
					continue TestLoop
				}
			}

			for fieldNumber, expected := range test.Expected[i] {
				actualValue := test.Dest[fieldNumber]

				if actual := stringAt(actualValue); expected != actual {
					t.Errorf("For test #%d and field #%d, expected (%T) %q, but actually got (%T) %q.", testNumber, fieldNumber, expected, expected, actualValue, actual)
				}
			}

			i++
		}
	}
}
