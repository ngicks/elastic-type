package mapping

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/geo-point.html#geo-point-params
type GeopointParams struct {
	// Type is type of this property. Automatically filled if zero.
	Type EsType `json:"type,omitempty"`
	// IgnoreMalformed indicates whether it should ignore malformed value rather than rejecting whole document.
	// Default(nil) is false.
	IgnoreMalformed *bool `json:"ignore_malformed,omitempty"`
	// Default(nil) is true.
	IgnoreZValue *bool `json:"ignore_z_value,omitempty"`
	// Index indicates whether the field should be quickly searchable.
	// Numeric fields that only have doc_values enabled can also be queried, albeit slower.
	// Default(nil) is true.
	Index *bool `json:"index,omitempty"`
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
}

func (p *GeopointParams) FillType() {
	if p.Type == "" {
		p.Type = Geopoint
	}
}
