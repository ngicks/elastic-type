package mapping

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/token-count.html#token-count-params
type TokenCountParams struct {
	// Type is type of this property. Automatically filled if zero.
	Type                     esType  `json:"type"`
	Analyzer                 *string `json:"analyzer,omitempty"`
	EnablePositionIncrements *bool   `json:"enable_position_increments,omitempty"`
	// Default(nil) is true.
	DocValues *bool `json:"doc_values,omitempty"`
	// Default(nil) is true.
	Index *bool `json:"index,omitempty"`
	// NullValue is substituted value for any explicit null (nil).
	// Defaults to null (nil), which means the field is treated as missing.
	// Invariants: invalid to set NullValue to true if the script parameter is set.
	NullValue *float64 `json:"null_value,omitempty"`
	// Store indicates whether the field value should be stored and retrievable separately from the _source field.
	// Default is false.
	Store *bool `json:"store,omitempty"`
}

func (p *TokenCountParams) FillType() {
	p.Type = TokenCount
}
