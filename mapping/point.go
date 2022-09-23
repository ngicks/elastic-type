package mapping

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/point.html#point-params
type PointParams struct {
	// Type is type of this property. Automatically filled if zero.
	Type esType `json:"type,omitempty"`
	// IgnoreMalformed indicates whether it should ignore malformed value rather than rejecting whole document.
	// Default(nil) is false.
	IgnoreMalformed *bool `json:"ignore_malformed,omitempty"`
	// Default(nil) is true.
	IgnoreZValue *bool `json:"ignore_z_value,omitempty"`
	// NullValue is substituted value for any explicit null (nil).
	// Defaults to null (nil), which means the field is treated as missing.
	// Invariants: invalid to set NullValue to true if the script parameter is set.
	NullValue *float64 `json:"null_value,omitempty"`
}

func (p *PointParams) FillType() {
	p.Type = Point
}
