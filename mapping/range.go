package mapping

// Params for
//   - IntegerRange
//   - FloatRange
//   - LongRange
//   - DoubleRange
//   - IpRange
//
// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/range.html#range-params
type RangeParams struct {
	// Type is type of this property. Automatically filled if zero.
	Type esType `json:"type,omitempty"`
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

func (p *RangeParams) FillType() {
	if p.Type == "" {
		p.Type = IntegerRange
	}
}

type DateRangeParams struct {
	Format string `json:"type,omitempty"`
	RangeParams
}

func (p *DateRangeParams) FillType() {
	if p.Type == "" {
		p.Type = DateRange
	}
}
