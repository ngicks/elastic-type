package estype_test

import (
	"errors"
	"fmt"
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
				A: estype.NewFieldSingleValue("foo"),
				B: estype.NewFieldSingleValue(123),
				C: estype.NewFieldSingleValue([]byte("baz")),
				D: estype.NewFieldSingleValue[estype.BooleanStr](true),
			},
			[]byte(`{"A":["foo"],"B":123,"c":["YmF6"],"d":"true"}`),
		},
		// contains nested
		{
			SampleFields{
				A: estype.NewFieldSingleValue("bar"),
				Nested: estype.NewFieldSingleValue(NestedSampleFields{
					A: estype.NewFieldSingleValue("baz"),
					B: estype.NewFieldSingleValue(95),
				}),
			},
			[]byte(`{"A":["bar"],"Nested":{"A":["baz"],"B":95}}`),
		},
		// contains null
		{
			SampleFields{
				A: estype.NewFieldSingleValue("foo"),
				B: estype.NewFieldNull[int](),
				D: estype.NewFieldNull[estype.BooleanStr](),
				Nested: estype.NewFieldSingleValue(NestedSampleFields{
					A: estype.NewFieldNull[string](),
					B: estype.NewFieldSingleValue(95),
				}),
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

type SampleWithNonField struct {
	A estype.Field[string]
	B struct {
		C string
		D int
	}
}

func TestMarshalFieldsJSON_contains_non_Field(t *testing.T) {
	require := require.New(t)

	input := SampleWithNonField{
		A: estype.NewFieldSingleValue("foo"),
		B: struct {
			C string
			D int
		}{
			C: "foo",
			D: 123,
		},
	}

	jsonEncoded, err := estype.MarshalFieldsJSON(input)
	require.NoError(err)
	require.Empty(cmp.Diff(
		[]byte(`{"A":["foo"],"B":{"C":"foo","D":123}}`),
		jsonEncoded,
	))
}

var errSample = errors.New("error")

type Erroneous string

func (e Erroneous) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("%w", errSample)
}

type SampleWithError struct {
	A estype.Field[string]
	B Erroneous
}

func TestMarshalFieldsJSON_err_of_marshaling_inner_value(t *testing.T) {
	require := require.New(t)

	input := SampleWithError{
		A: estype.NewFieldNull[string](),
		B: Erroneous("bar"),
	}

	_, err := estype.MarshalFieldsJSON(input)

	require.ErrorIs(err, errSample)
}
