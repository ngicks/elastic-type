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
	type TestBool struct {
		A estype.Boolean
	}
	type TestBoolStr struct {
		A estype.BooleanStr
	}
	testBool := TestBool{}
	testBoolStr := TestBoolStr{}

	testBooleanUnmarshal(t, testBool)
	testBooleanUnmarshal(t, testBoolStr)

	var err error
	var syntaxError *json.SyntaxError
	err = testBool.A.UnmarshalJSON([]byte(`dawju9813`))
	require.ErrorAs(t, err, &syntaxError)
	err = testBoolStr.A.UnmarshalJSON([]byte(`dawju9813`))
	require.ErrorAs(t, err, &syntaxError)
}

func testBooleanUnmarshal[T any](t *testing.T, testBool T) {
	var invalidTypeError *estype.InvalidTypeError
	for _, testCase := range []string{
		`{"A": "foo"}`,
		`{"A": 123}`,
		`{"A": 123.5}`,
	} {
		err := json.Unmarshal([]byte(testCase), &testBool)
		require.ErrorAs(t, err, &invalidTypeError)
	}

	err := json.Unmarshal([]byte(`{"A": kc;a123}`), &testBool)
	var syntaxError *json.SyntaxError
	require.ErrorAs(t, err, &syntaxError)
}

func TestBooleanString(t *testing.T) {
	esBoolean := estype.Boolean(false)
	esBooleanStr := estype.BooleanStr(false)

	require.Equal(t, "false", esBoolean.String())
	require.Equal(t, "false", esBooleanStr.String())

	esBoolean = estype.Boolean(true)
	esBooleanStr = estype.BooleanStr(true)

	require.Equal(t, "true", esBoolean.String())
	require.Equal(t, "true", esBooleanStr.String())
}
