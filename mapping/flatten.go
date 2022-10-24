package mapping

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/flattened.html#flattened-params
type FlattenedParams struct {
	// Type is type of this property. Automatically filled if zero.
	Type esType `json:"type,omitempty"`
	// Defaults to 20.
	DepthLimit *int `json:"depth_limit,omitempty"`
	// DocValues indicates whether it should save field on disk in a column-stride fashion,
	// so that it can later be used for sorting, aggregations, or scripting.
	// Default(nil) is true.
	DocValues *bool `json:"doc_values,omitempty"`
	// Defaults to false.
	EagerGlobalOrdinals *bool `json:"eager_global_ordinals,omitempty"`
	IgnoreAbove         *int  `json:"ignore_above,omitempty"`
	// Index indicates whether the field should be quickly searchable.
	// Default(nil) is true.
	Index        *bool         `json:"index,omitempty"`
	IndexOptions *indexOptions `json:"index_options,omitempty"`
	// NullValue is substituted value for any explicit null (nil).
	// Defaults to null (nil), which means the field is treated as missing.
	NullValue *float64 `json:"null_value,omitempty"`
	// Defaults to "BM25".
	// Only "BM25" and "boolean" are available out-of-box.
	Similarity *string `json:"similarity,omitempty"`
	// Defaults to false.
	SplitQueriesOnWhitespace *bool `json:"split_queries_on_whitespace,omitempty"`
}

func (p *FlattenedParams) FillType() {
	if p.Type == "" {
		p.Type = Flattened
	}
}

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/index-options.html
type indexOptions string

const (
	Docs      indexOptions = "docs"
	Freqs     indexOptions = "freqs"
	Positions indexOptions = "positions" // Default
	Offsets   indexOptions = "offsets"
)
