package mapping

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/nested.html#nested-params
type NestedProperty struct {
	types.NestedProperty
	Properties *Properties `json:"properties,omitempty"`
}
