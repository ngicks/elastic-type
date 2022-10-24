package mapping

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/shape.html#shape-mapping-options
type ShapeParams struct {
	// Type is type of this property. Automatically filled if zero.
	Type esType `json:"type,omitempty"`
	// Defaults to "ccw".
	Orientation *orientation `json:"orientation,omitempty"`
	// IgnoreMalformed indicates whether it should ignore malformed value rather than rejecting whole document.
	// Default(nil) is false.
	IgnoreMalformed *bool `json:"ignore_malformed,omitempty"`
	// Default(nil) is true.
	IgnoreZValue *bool `json:"ignore_z_value,omitempty"`
	// Default(nil) is false.
	Coerce *bool `json:"coerce,omitempty"`
}

func (p *ShapeParams) FillType() {
	if p.Type == "" {
		p.Type = Shape
	}
}
