package test_test

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ngicks/elastic-type/mapping"
	"github.com/stretchr/testify/require"
)

var (
	//go:embed example/all.json
	allJSONBin       []byte
	sampleMappingBin []byte
	testSettings     = map[string]any{
		"number_of_replicas": 0, // This prevents es from being yellow after creation of index. Only needed if es is single-node.
	}
)

func init() {
	// Comprehensive mapping setting with all fields populated with default or sensible value.
	var allJSON map[string]map[string]any

	var err error
	err = json.Unmarshal(allJSONBin, &allJSON)
	if err != nil {
		panic(err)
	}

	getOne := func(v map[string]map[string]any) map[string]any {
		for _, v := range v {
			return v
		}
		return nil
	}

	sampleMapping := map[string]any{
		"settings": testSettings,
		"mappings": getOne(allJSON)["mappings"],
	}

	sampleMappingBin, err = json.MarshalIndent(sampleMapping, "", "     ")
	if err != nil {
		panic(err)
	}
}

func TestMapping(t *testing.T) {
	require := require.New(t)

	// These parts are just doing validation of our input,
	// checking that input is legitimate for real Elasticsearch instance.
	//
	// Create index with our inputs, and then retrieve those mapping.
	// Then we can decode and encode the mapping through our type,
	// make sure our type does not loose any field of it.
	skipIfEsNotReachable(t, *ELASTICSEARCH_URL, false)
	indexName := must(createRandomIndex(*ELASTICSEARCH_URL, sampleMappingBin))
	defer deleteIndex(*ELASTICSEARCH_URL, indexName)

	bin := must(getMapping(*ELASTICSEARCH_URL, indexName))

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
	// But it fills doc_values field if type is "search_as_you_type"
}
