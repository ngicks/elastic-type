package estype_test

import (
	"encoding/json"
	"fmt"
	"testing"

	estype "github.com/ngicks/elastic-type/es_type"
	"github.com/stretchr/testify/require"
)

type TestBool struct {
	A    estype.Boolean
	B    estype.Boolean
	Astr estype.BooleanStr
	Bstr estype.BooleanStr
}

func TestBoolean(t *testing.T) {
	for _, testCase := range [][4]bool{
		{true, true, true, true},
		{true, false, true, false},
		{false, true, false, true},
		{false, false, false, false},
	} {
		testBool := TestBool{}
		err := json.Unmarshal(
			[]byte(fmt.Sprintf(
				`{"A": %t, "B": "%t", "Astr": %t, "Bstr": "%t"}`,
				testCase[0],
				testCase[1],
				testCase[2],
				testCase[3],
			)),
			&testBool,
		)
		require.NoError(t, err)
		require.Equal(
			t,
			TestBool{
				A:    estype.Boolean(testCase[0]),
				B:    estype.Boolean(testCase[1]),
				Astr: estype.BooleanStr(testCase[2]),
				Bstr: estype.BooleanStr(testCase[3]),
			},
			testBool,
		)
		bin, _ := json.Marshal(testBool)

		testBool2 := TestBool{}
		err = json.Unmarshal(bin, &testBool2)
		require.NoError(t, err)
		require.Equal(t, testBool, testBool2)
	}
}

func TestBooleanInvalidInput(t *testing.T) {
	testBool := TestBool{}

	var invalidTypeError *estype.InvalidTypeError
	for _, testCase := range []string{
		`{"A": "foo"}`,
		`{"A": 123}`,
		`{"A": 123.5}`,
	} {
		err := json.Unmarshal([]byte(testCase), &testBool)
		require.ErrorAs(t, err, &invalidTypeError)
	}
}
