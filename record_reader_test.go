package tabio


import (
	"strings"

	"testing"
)


func TestNewRecordReader(t *testing.T) {
	reader := strings.NewReader("apple banana cherry")

	recordReader := NewRecordReader(reader)
	if nil == recordReader {
		t.Errorf("Did not expect nil but actually got: %v", recordReader)
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
			String: "fruit"  +RS,

			Expected: []string{"fruit"},
		},
		{
			String: "fruit",

			Expected: []string{"fruit"},
		},
	}


	TestLoop: for testNumber, test := range tests {

		rr := NewRecordReader( strings.NewReader(test.String) )

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
			String: "fruit"  +RS,

			Expected: 0,
		},
		{
			String: "fruit",

			Expected: 0,
		},
	}


	TestLoop: for testNumber, test := range tests {

		rr := NewRecordReader( strings.NewReader(test.String) )

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
			String: "fruit"  +RS,

			Expected: [][]string{},
		},
		{
			String: "fruit",

			Expected: [][]string{},
		},
	}


	TestLoop: for testNumber, test := range tests {

		rr := NewRecordReader( strings.NewReader(test.String) )

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
