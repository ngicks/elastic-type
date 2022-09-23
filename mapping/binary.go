package mapping

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/binary.html#binary-params
type BinaryParams struct {
	// DocValues indicates whether it should save field on disk in a column-stride fashion,
	// so that it can later be used for sorting, aggregations, or scripting.
	// Default(nil) is true.
	DocValues *bool `json:"doc_values,omitempty"`
	// Store indicates whether the field value should be stored and retrievable separately from the _source field.
	// Default is false.
	Store *bool `json:"store,omitempty"`
}
