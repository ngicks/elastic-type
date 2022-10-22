package estype_test

import (
	"encoding/json"
	"testing"

	estype "github.com/ngicks/elastic-type/es_type"
	"github.com/stretchr/testify/require"
)

func TestBinary(t *testing.T) {
	var bin estype.Binary
	src := "SGVsbG8gV29ybGQ="
	srcJsonValue := `"` + src + `"`
	binData := []byte("Hello World")

	test := func() {
		require.Equal(t, src, bin.String())
		require.Equal(t, string(binData), string(bin.Bytes()))
	}

	err := json.Unmarshal([]byte(srcJsonValue), &bin)
	require.NoError(t, err)

	test()

	bin = estype.StringToBinaryUnchecked(src)
	test()

	bin, err = estype.StringToBinary(src, false)
	require.NoError(t, err)
	test()

	bin, err = estype.StringToBinary(src, true)
	require.NoError(t, err)
	test()

	bin = estype.BytesToBinary(binData)
	test()

	marshalled, err := json.Marshal(bin)
	require.NoError(t, err)
	require.Equal(t, []byte(srcJsonValue), marshalled)
}

func TestBinaryError(t *testing.T) {
	var bin estype.Binary

	invalidString := "Invalid"
	invalidStringJsonValue := `"` + invalidString + `"`

	var err error

	_, err = estype.StringToBinary(invalidString, false)
	require.Error(t, err)

	err = json.Unmarshal([]byte(invalidStringJsonValue), &bin)
	require.Error(t, err)

	err = json.Unmarshal([]byte("1235"), &bin)
	require.Error(t, err)
}
