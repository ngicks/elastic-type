package mapping

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/object.html#object-params
type ObjectProperty struct {
	types.ObjectProperty
	Properties *Properties `json:"properties,omitempty"`
}
