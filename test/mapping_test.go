package test_test

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ngicks/elastic-type/mapping"
	"github.com/ngicks/elastic-type/test"
	"github.com/stretchr/testify/require"
)

func TestMapping(t *testing.T) {
	require := require.New(t)

	// These parts are just doing validation of our input,
	// checking that input is legitimate for real Elasticsearch instance.
	//
	// Create index with our inputs, and then retrieve those mapping.
	// Then we can decode and encode the mapping through our type,
	// make sure our type does not loose any field of it.
	skipIfEsNotReachable(t, *ELASTICSEARCH_URL, false)
	helper := must(createRandomIndex[any](client, test.AllMappings))
	t.Log(helper.IndexName)
	defer helper.Delete()

	bin := must(helper.GetMapping())

	var mappingOfOurDefinedType mapping.MappingSettings
	// error should not happen
	_ = json.Unmarshal(bin, &mappingOfOurDefinedType)
	anyMapEncodedThroughOurType := toAnyMap(mappingOfOurDefinedType)

	var storedMapping map[string]any
	_ = json.Unmarshal(bin, &storedMapping)

	require.Conditionf(
		func() bool {
			return cmp.Equal(anyMapEncodedThroughOurType, storedMapping)
		},
		"not equal: diff = %s",
		cmp.Diff(storedMapping, anyMapEncodedThroughOurType),
	)

	// Note that, I do not why. and I do not know it is documented.
	// But it fills doc_values field if type is "search_as_you_type".
}
