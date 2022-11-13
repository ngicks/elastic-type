package test_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerate_all(t *testing.T) {
	require := require.New(t)

	skipIfEsNotReachable(t, *ELASTICSEARCH_URL, false)
	indexName := must(createRandomIndex(*ELASTICSEARCH_URL, allMappings))
	defer deleteIndex(*ELASTICSEARCH_URL, indexName)

	res, err := postDoc(*ELASTICSEARCH_URL, indexName, map[string]any{
		"date": "2022-11-14 23:22:12",
	})
	t.Logf("%+v", res)
	require.NoError(err)

	fetchedDoc, err := getDoc(*ELASTICSEARCH_URL, indexName, res.Id_)
	t.Logf("%+v", fetchedDoc.Source_)
	require.NoError(err)
}
