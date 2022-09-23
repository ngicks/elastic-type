package mapping

// DateParams is property for date and date_nanos.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/date.html#date-params
type DateParams struct {
	// DocValues indicates whether it should save field on disk in a column-stride fashion,
	// so that it can later be used for sorting, aggregations, or scripting.
	// Default(nil) is true.
	DocValues *bool `json:"doc_values,omitempty"`
	// Format is time token layout used to parse string. Format is double vertical line (`||`) delimitted string.
	// Defaults to  "strict_date_optional_time||epoch_millis" for date,
	// "strict_date_optional_time_nanos||epoch_millis" for date_nanos.
	Format *string `json:"format,omitempty"`
	Locale *string `json:"locale,omitempty"`
	// IgnoreMalformed indicates whether it should ignore malformed value rather than rejecting whole document.
	// Default(nil) is false.
	IgnoreMalformed *bool `json:"ignore_malformed,omitempty"`
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
	// Store indicates whether the field value should be stored and retrievable separately from the _source field.
	// Default is false.
	Store *bool `json:"store,omitempty"`
	// Meta is metadata about the field.
	Meta *Meta `json:"meta,omitempty"`
}
