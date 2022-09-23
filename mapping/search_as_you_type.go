package mapping

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/search-as-you-type.html#specific-params
// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/search-as-you-type.html#general-params
type SearchAsYouTypeParams struct {
	// Valid values are 2 (inclusive) to 4 (inclusive).
	// Defaults to 3.
	MaxShingleSize *int `json:"max_shingle_size,omitempty"`
	// Defaults to default index analyzer or "standard".
	Analyzer *string `json:"analyzer,omitempty"`
	// Index indicates whether the field should be quickly searchable.
	// Numeric fields that only have doc_values enabled can also be queried, albeit slower.
	// Default(nil) is true.
	Index *bool `json:"index,omitempty"`
	// Defaults to Positions.
	IndexOptions *indexOptions `json:"index_options,omitempty"`
	// Defaults to true.
	Norms *bool `json:"norms,omitempty"`
	// Default is false.
	Store *bool `json:"store,omitempty"`
	// Defaults to the Analyzer setting.
	SearchAnalyzer *string `json:"search_analyzer,omitempty"`
	// Defaults to the SearchAnalyzer setting.
	SearchQuoteAnalyzer *string `json:"search_quote_analyzer,omitempty"`
	// Defaults to "BM25".
	// Only "BM25" and "boolean" are available out-of-box.
	Similarity *string `json:"similarity,omitempty"`
	// Defaults to "no".
	TermVector *termVector `json:"term_vector,omitempty"`
}
