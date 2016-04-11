package tabio


import (
	"io/ioutil"
	"strings"

	"testing"
)


func TestNewRecordReader(t *testing.T) {
	reader := ioutil.NopCloser( strings.NewReader("apple banana cherry") )

	rr := NewRecordReader(reader)
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

		rr := NewRecordReader( ioutil.NopCloser(strings.NewReader(test.String)) )

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
			        "cherry" +RS,

			Expected: 3,
		},
		{
			String: "fruit"  +RS+
			        "apple"  +RS+
			        "banana" +RS+
			        "cherry",

			Expected: 3,
		},



		{
			String: "fruit"  +RS,

			Expected: 0,
		},
		{
			String: "fruit",

			Expected: 0,
		},
	}


	TestLoop: for testNumber, test := range tests {

		rr := NewRecordReader( ioutil.NopCloser( strings.NewReader(test.String) ) )

		count := 0
		for rr.Next() {
			count++

			if expected, actual := test.Expected, count; expected < actual {
				t.Errorf("For test %#d, expected count to be less than or equal to %d, but became %d,", testNumber, expected, actual)
				continue TestLoop
			}
		}

		if expected, actual := test.Expected, count; expected != actual {
			t.Errorf("For test #%d, expected %d but actually got %d.", testNumber, expected, actual)
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

		rr := NewRecordReader( ioutil.NopCloser( strings.NewReader(test.String) ) )

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
		{
			String: "given_name" +US+ "family_name" +US+ "city"      +RS+
			        "Joe"        +US+ "Blow"        +US+ "Vancouver" +RS+
			        "Jane"       +US+ "Doe"         +US+ "Surrey"    +RS,

			Expected: [][]string{
				[]string{"Joe",  "Blow", "Vancouver"},
				[]string{"Jane", "Doe",  "Surrey"},
			},
			Dest: []interface{}{ (*string)(nil), (*string)(nil), (*string)(nil) },
		},
		{
			String: "given_name" +US+ "family_name" +US+ "city"      +RS+
			        "Joe"        +US+ "Blow"        +US+ "Vancouver" +RS+
			        "Jane"       +US+ "Doe"         +US+ "Surrey",

			Expected: [][]string{
				[]string{"Joe",  "Blow", "Vancouver"},
				[]string{"Jane", "Doe",  "Surrey"},
			},
			Dest: []interface{}{ (*string)(nil), (*string)(nil), (*string)(nil) },
		},



		{
			String: "given_name" +US+ "family_name" +US+ "city"      +RS,

			Expected: [][]string{},
			Dest: []interface{}{ (*string)(nil), (*string)(nil), (*string)(nil) },
		},
		{
			String: "given_name" +US+ "family_name" +US+ "city",

			Expected: [][]string{},
			Dest: []interface{}{ (*string)(nil), (*string)(nil), (*string)(nil) },
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
			Dest: []interface{}{ (*string)(nil) },
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
			Dest: []interface{}{ (*string)(nil) },
		},



		{
			String: "fruit"  +RS,

			Expected: [][]string{},
			Dest: []interface{}{ (*string)(nil) },
		},
		{
			String: "fruit",

			Expected: [][]string{},
			Dest: []interface{}{ (*string)(nil) },
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
			Dest: []interface{}{ (*string)(nil), (*string)(nil) },
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
			Dest: []interface{}{ (*int)(nil), (*int)(nil) },
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
			Dest: []interface{}{ (*int)(nil), (*int)(nil) },
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

		rr := NewRecordReader( ioutil.NopCloser( strings.NewReader(test.String) ) )

		i := 0
		for rr.Next() {
			if err := rr.Scan(test.Dest...); nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one: %v", testNumber, err)
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

			i++
		}
	}
}
