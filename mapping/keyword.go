package mapping

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/keyword.html#keyword-params
type KeywordParams struct {
	// Type is type of this property. Automatically filled if zero.
	Type esType `json:"type,omitempty"`
	// DocValues indicates whether it should save field on disk in a column-stride fashion,
	// so that it can later be used for sorting, aggregations, or scripting.
	// Default(nil) is true.
	DocValues *bool `json:"doc_values,omitempty"`
	// Defaults to false.
	EagerGlobalOrdinals *bool   `json:"eager_global_ordinals,omitempty"`
	Fields              *Fields `json:"fields,omitempty"`
	// Defaults to 256
	IgnoreAbove *int `json:"ignore_above,omitempty"`
	// Index indicates whether the field should be quickly searchable.
	// Numeric fields that only have doc_values enabled can also be queried, albeit slower.
	// Default(nil) is true.
	Index *bool `json:"index,omitempty"`
	// Defaults to docs.
	IndexOptions *indexOptions `json:"index_options,omitempty"`
	// Meta is metadata about the field.
	Meta *Meta `json:"meta,omitempty"`
	// Defaults to false.
	Norms *bool `json:"norms,omitempty"`
	// NullValue is substituted value for any explicit null (nil).
	// Defaults to null (nil), which means the field is treated as missing.
	// Invariants: invalid to set NullValue to true if the script parameter is set.
	NullValue *float64 `json:"null_value,omitempty"`
	// OnScriptError indicates whether it should continue or fail when script defined for this field throws.
	// Defaults to "fail",  which will cause the entire document to be rejected.
	// If OnScriptError is "continue", which will register the field in the document’s _ignored metadata field and continue indexing.
	// OnScriptError can only be set if the script field is also set.
	OnScriptError *onScriptError `json:"on_script_error,omitempty"`
	// Script defines script to generate value for this field, executed at idexing time.
	// If a value is set for this field on the input document, then the document will be rejected with an error.
	// Scripts can only be configured on long and double field types.
	Script *Script `json:"script,omitempty"`
	// Store indicates whether the field value should be stored and retrievable separately from the _source field.
	// Default is false.
	Store *bool `json:"store,omitempty"`
	// Defaults to "BM25".
	// Only "BM25" and "boolean" are available out-of-box.
	Similarity *string `json:"similarity,omitempty"`
	// Defaults to null.
	Normalizer *string `json:"normalizer,omitempty"`
	// Defaults to false.
	SplitQueriesOnWhitespace *bool `json:"split_queries_on_whitespace,omitempty"`
	// [preview]
	//
	// TimeSeriesDimension marks the field as a time series dimension.
	// Defaults to false.
	TimeSeriesDimension *bool `json:"time_series_dimension,omitempty"`
}

func (p *KeywordParams) FillType() {
	p.Type = Keyword
}

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/multi-fields.html
type Fields map[string]SubFieldType

type SubFieldType struct {
	// Type is text or keyword
	Type     esType  `json:"type"`
	Analyzer *string `json:"analyzer,omitempty"`
}

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/keyword.html#constant-keyword-field-type
type ConstantKeywordParams struct {
	// Type is type of this property. Automatically filled if zero.
	Type esType `json:"type,omitempty"`
	// Meta is metadata about the field.
	Meta *Meta `json:"meta,omitempty"`
	// If value is not set, first indexed value will be used.
	Value *string `json:"value,omitempty"`
}

func (p *ConstantKeywordParams) FillType() {
	p.Type = ConstantKeyword
}

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/keyword.html#constant-keyword-field-type
type WildcardParams struct {
	// Type is type of this property. Automatically filled if zero.
	Type esType `json:"type,omitempty"`
	// NullValue is substituted value for any explicit null (nil).
	// Defaults to null (nil), which means the field is treated as missing.
	// Invariants: invalid to set NullValue to true if the script parameter is set.
	NullValue *float64 `json:"null_value,omitempty"`
	// Defaults to 2147483647
	IgnoreAbove *int `json:"ignore_above,omitempty"`
}

func (p *WildcardParams) FillType() {
	p.Type = Wildcard
}
