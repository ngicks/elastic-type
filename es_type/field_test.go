package estype_test

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	estype "github.com/ngicks/elastic-type/es_type"
	"github.com/stretchr/testify/require"
)

type testFieldStruct struct {
	A estype.Field[string]         `json:"a"`
	B estype.Field[estype.Boolean] `json:"b"`
}

func TestFieldMarshal(t *testing.T) {
	require := require.New(t)

	var bin []byte
	var err error

	var f testFieldStruct

	f.A.SetNull()
	f.B.SetNull()

	bin, err = json.Marshal(f)
	require.NoError(err)
	require.Equal([]byte(`{"a":null,"b":null}`), bin)

	f.A.SetUndefined()
	f.B.SetUndefined()

	bin, err = json.Marshal(f)
	require.NoError(err)
	// `undefined` is marshalled into null anyway. Upper structs must handle marshalling behavior.
	require.Equal([]byte(`{"a":null,"b":null}`), bin)

	f.A.SetValue([]string{"foo", "bar"})
	f.B.SetSingleValue(estype.Boolean(true))

	bin, err = json.Marshal(f)
	require.NoError(err)
	require.Equal([]byte(`{"a":["foo","bar"],"b":[true]}`), bin)

	bin, err = json.Marshal(f)
	require.NoError(err)
	require.Equal([]byte(`{"a":["foo","bar"],"b":[true]}`), bin)
}

func TestFieldUnmarshal(t *testing.T) {
	require := require.New(t)

	var f testFieldStruct

	err := json.Unmarshal([]byte(`{"a":null}`), &f)
	require.NoError(err)

	require.True(f.A.IsNull())
	require.True(f.B.IsUndefined())

	err = json.Unmarshal([]byte(`{"a":["foo","bar"],"b":true}`), &f)
	require.NoError(err)

	require.Condition(
		func() bool { return cmp.Equal([]string{"foo", "bar"}, f.A.Unwrap()) },
		cmp.Diff([]string{"foo", "bar"}, f.A.Unwrap()),
	)
	require.Condition(
		func() bool { return cmp.Equal([]estype.Boolean{estype.Boolean(true)}, f.B.Unwrap()) },
		cmp.Diff([]estype.Boolean{estype.Boolean(true)}, f.B.Unwrap()),
	)

	err = json.Unmarshal([]byte(`{"a":"foo","b":[true]}`), &f)
	require.NoError(err)

	require.Condition(
		func() bool { return cmp.Equal([]string{"foo"}, f.A.Unwrap()) },
		cmp.Diff([]string{"foo"}, f.A.Unwrap()),
	)
	require.Condition(
		func() bool { return cmp.Equal([]estype.Boolean{estype.Boolean(true)}, f.B.Unwrap()) },
		cmp.Diff([]estype.Boolean{estype.Boolean(true)}, f.B.Unwrap()),
	)

	err = json.Unmarshal([]byte(`{"a":123,"b":true}`), &f)
	require.Error(err)

	err = json.Unmarshal([]byte(`{"a":[123],"b":true}`), &f)
	require.Error(err)
}

func TestFieldUtilities(t *testing.T) {
	require := require.New(t)

	var f testFieldStruct
	testPanicOnUnwrap := func(shouldUnwrapPanic bool, shouldUnwrapSinglePanic bool) {
		func() {
			defer func() {
				rec := recover()
				if shouldUnwrapPanic {
					require.NotNil(rec)
				} else {
					require.Nil(rec)
				}
			}()
			f.A.Unwrap()
		}()

		func() {
			defer func() {
				rec := recover()
				if shouldUnwrapSinglePanic {
					require.NotNil(rec)
				} else {
					require.Nil(rec)
				}
			}()
			f.A.UnwrapSingle()
		}()
	}

	require.True(f.A.IsUndefined())
	require.False(f.A.IsNull())
	require.True(f.A.IsEmpty())
	require.Nil(f.A.Value())
	require.Nil(f.A.ValueSingle())
	testPanicOnUnwrap(true, true)

	f.A.SetNull()

	require.False(f.A.IsUndefined())
	require.True(f.A.IsNull())
	require.True(f.A.IsEmpty())
	testPanicOnUnwrap(true, true)

	f.A.SetValue([]string{})

	require.False(f.A.IsUndefined())
	require.False(f.A.IsNull())
	require.True(f.A.IsEmpty())
	require.Equal(&[]string{}, f.A.Value())
	require.Nil(f.A.ValueSingle())
	testPanicOnUnwrap(false, true)

	f.A.SetValue([]string{"foo", "bar"})

	require.Condition(
		func() bool { return cmp.Equal([]string{"foo", "bar"}, *f.A.Value()) },
		cmp.Diff([]string{"foo", "bar"}, f.A.Unwrap()),
	)
	require.Equal("foo", *f.A.ValueSingle())
	require.False(f.A.IsEmpty())
	testPanicOnUnwrap(false, false)

	f.A.SetSingleValue("baz")
	require.Condition(
		func() bool { return cmp.Equal([]string{"baz"}, f.A.Unwrap()) },
		cmp.Diff([]string{"baz"}, f.A.Unwrap()),
	)
	require.Equal("baz", f.A.UnwrapSingle())
	require.False(f.A.IsEmpty())
}
