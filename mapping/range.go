package mapping

// Params for
//   - IntegerRange
//   - FloatRange
//   - LongRange
//   - DoubleRange
//   - DateRange
//   - IpRange
//
// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/range.html#range-params
type RangeParams struct {
	// Default(nil) is true.
	Coerce *bool `json:"coerce,omitempty"`
	// Index indicates whether the field should be quickly searchable.
	// Numeric fields that only have doc_values enabled can also be queried, albeit slower.
	// Default(nil) is true.
	Index *bool `json:"index,omitempty"`
	// Store indicates whether the field value should be stored and retrievable separately from the _source field.
	// Default is false.
	Store *bool `json:"store,omitempty"`
}
