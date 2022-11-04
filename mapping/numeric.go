package mapping

// NumericProperty is properties of Numeric field types.
//
// see https://www.elastic.co/guide/en/elasticsearch/reference/8.4/number.html#number-params
type NumericParams struct {
	// Type is type of this property. Automatically filled if zero.
	// Default is Integer.
	Type esType `json:"type,omitempty"`
	// Coerce indicates whether it should try to convert other json value to number.
	// Default(nil) is true.
	// This is not applicable for unsigned_long.
	// Invariants: invalid to set Coerce to true if the script parameter is set.
	Coerce *bool `json:"coerce,omitempty"`
	// DocValues indicates whether it should save field on disk in a column-stride fashion,
	// so that it can later be used for sorting, aggregations, or scripting.
	// Default(nil) is true.
	DocValues *bool `json:"doc_values,omitempty"`
	// IgnoreMalformed indicates whether it should ignore malformed value rather than rejecting whole document.
	// Default(nil) is false.
	IgnoreMalformed *bool `json:"ignore_malformed,omitempty"`
	// Index indicates whether the field should be quickly searchable.
	// Numeric fields that only have doc_values enabled can also be queried, albeit slower.
	// Default(nil) is true.
	Index *bool `json:"index,omitempty"`
	// Meta is metadata about the field.
	Meta *Meta `json:"meta,omitempty"`
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
	// [preview]
	//
	// TimeSeriesDimension marks the field as a time series dimension.
	// Defaults to false.
	TimeSeriesDimension *bool `json:"time_series_dimension,omitempty"`
}

func (p *NumericParams) FillType() {
	if p.Type == "" {
		p.Type = Integer
	}
}

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/number.html#scaled-float-params
type ScaledFloatParams struct {
	NumericParams
	ScalingFactor float64 `json:"scaling_factor"`
}

func (p *ScaledFloatParams) FillType() {
	if p.Type == "" {
		p.Type = ScaledFloat
	}
}
