package estype_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	estype "github.com/ngicks/elastic-type/es_type"
	"github.com/stretchr/testify/require"
)

type SampleFields struct {
	A      estype.Field[string]
	B      estype.Field[int]                `esjson:"pseudo,single"`   // contains a fake tag
	C      estype.Field[[]byte]             `json:"c"`                 // has a json tag, no esfield tag
	D      estype.Field[estype.BooleanStr]  `json:"d" esjson:"single"` // has both
	Nested estype.Field[NestedSampleFields] `esjson:"single"`          // nested
}

type NestedSampleFields struct {
	A estype.Field[string]
	B estype.Field[int] `esjson:"single"`
}

func (n NestedSampleFields) MarshalJSON() ([]byte, error) {
	return estype.MarshalFieldsJSON(n)
}

func TestMarshalFieldsJSON_happy_path(t *testing.T) {
	require := require.New(t)

	cases := []struct {
		input  SampleFields
		expect []byte
	}{
		// all undefined
		{
			SampleFields{},
			[]byte(`{}`),
		},
		{
			SampleFields{
				A: estype.NewFieldSingleValue("foo", true),
				B: estype.NewFieldSingleValue(123, true),
				C: estype.NewFieldSingleValue([]byte("baz"), true),
				D: estype.NewFieldSingleValue[estype.BooleanStr](true, true),
			},
			[]byte(`{"A":["foo"],"B":123,"c":["YmF6"],"d":"true"}`),
		},
		// contains nested
		{
			SampleFields{
				A: estype.NewFieldSingleValue("bar", true),
				Nested: estype.NewFieldSingleValue(NestedSampleFields{
					A: estype.NewFieldSingleValue("baz", false),
					B: estype.NewFieldSingleValue(95, false),
				}, false),
			},
			[]byte(`{"A":["bar"],"Nested":{"A":["baz"],"B":95}}`),
		},
		// contains null
		{
			SampleFields{
				A: estype.NewFieldSingleValue("foo", true),
				B: estype.NewFieldNull[int](false),
				D: estype.NewFieldNull[estype.BooleanStr](true),
				Nested: estype.NewFieldSingleValue(NestedSampleFields{
					A: estype.NewFieldNull[string](false),
					B: estype.NewFieldSingleValue(95, false),
				}, false),
			},
			[]byte(`{"A":["foo"],"B":null,"d":null,"Nested":{"A":null,"B":95}}`),
		},
	}

	for _, tc := range cases {
		jsonEncoded, err := estype.MarshalFieldsJSON(tc.input)
		require.NoError(err)
		require.Conditionf(
			func() bool { return cmp.Equal(tc.expect, jsonEncoded) },
			"diff: %s", cmp.Diff(string(tc.expect), string(jsonEncoded)),
		)
	}
}

func TestMarshalFieldsJSON_err_if_input_contains_non_Field(t *testing.T) {
	t.Skip("not implemented")
}

func TestMarshalFieldsJSON_err_of_marshaling_inner_value(t *testing.T) {
	t.Skip("not implemented")
}
