package tabio_test

import (
	"github.com/reiver/go-tabio"

	"io/ioutil"
	"strings"

	"testing"
)

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
