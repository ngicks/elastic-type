package test_test

import (
	"testing"
	"time"

	estype "github.com/ngicks/elastic-type/es_type"
	"github.com/ngicks/elastic-type/test"
	"github.com/ngicks/elastic-type/test/example"
	"github.com/stretchr/testify/require"
)

func TestGenerate_all(t *testing.T) {
	require := require.New(t)

	skipIfEsNotReachable(t, *ELASTICSEARCH_URL, false)
	helper := must(createRandomIndex[example.AllRaw](client, test.AllMappings))
	defer helper.Delete()

	now := time.Now()

	id, err := helper.PostDoc(example.AllRaw{
		Date: estype.NewFieldSingleValue(example.AllDate(now), false),
	})
	require.NoError(err)

	fetchedDoc, err := helper.GetDoc(id)
	t.Logf("%+v", fetchedDoc.Source_)
	require.NoError(err)
}
