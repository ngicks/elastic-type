package test_test

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ngicks/elastic-type/mapping"
	"github.com/stretchr/testify/require"
)

var (
	sampleMapping    = map[string]any{}
	sampleMappingBin []byte
	testSettings     = map[string]any{
		"number_of_replicas": 0, // This prevents es from being yellow after creation of index. Only needed if es is single-node.
	}
)

func init() {
	// Comprehensive mapping setting with all fields populated with default or sensible value.
	props := map[string]map[string]any{
		"agg": {
			"type":           "aggregate_metric_double",
			"metrics":        []string{"min", "max", "sum", "value_count"},
			"default_metric": "max",
		},
		"alias": {
			"type": "alias",
			"path": "blob",
		},
		"blob": {
			"type":       "binary",
			"doc_values": false,
			"store":      false,
		},
		"bool": {
			"type":            "boolean",
			"doc_values":      true,
			"index":           true,
			"null_value":      nil,
			"on_script_error": mapping.Continue,
			"script": map[string]any{
				"source": "emit(false)",
			},
			"store": false,
			"meta":  map[string]any{},
		},
		"comp": {
			"type":                         "completion",
			"analyzer":                     "simple",
			"search_analyzer":              "simple",
			"preserve_separators":          true,
			"preserve_position_increments": true,
			"max_input_length":             255,
		},
		"date": {
			"type":             "date",
			"doc_values":       true,
			"format":           "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis",
			"locale":           "ja-jp",
			"ignore_malformed": false,
			"index":            true,
			"null_value":       nil,
			"on_script_error":  mapping.Fail,
			"script": map[string]any{
				"source": "emit(new Date().getTime())",
				"lang":   "painless",
			},
			"store": false,
			"meta": map[string]any{
				"metric_type": "gauge",
			},
		},
		"dateNano": {
			"type":             "date",
			"doc_values":       true,
			"format":           "strict_date_optional_time_nanos||epoch_millis",
			"locale":           "ja-jp",
			"ignore_malformed": false,
			"index":            true,
			"null_value":       nil,
			"on_script_error":  mapping.Fail,
			"script": map[string]any{
				"source": "emit(new Date().getTime())",
				"lang":   "painless",
			},
			"store": false,
			"meta": map[string]any{
				"metric_type": "gauge",
			},
		},
		"dense_vector": {
			"type":       "dense_vector",
			"dims":       3,
			"index":      true,
			"similarity": mapping.L2Norm,
			"index_options": map[string]any{
				"type":            "hnsw",
				"m":               16,
				"ef_construction": 100,
			},
		},
		"flattened": {
			"type":                        "flattened",
			"depth_limit":                 20,
			"doc_values":                  true,
			"eager_global_ordinals":       false,
			"ignore_above":                255,
			"index":                       true,
			"index_options":               mapping.Freqs,
			"null_value":                  nil,
			"similarity":                  "BM25",
			"split_queries_on_whitespace": false,
		},
		// TODO: expand til the end...
	}

	sampleMapping = map[string]any{
		"settings": testSettings,
		"mappings": map[string]any{
			"properties": props,
		},
	}

	var err error
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
}
