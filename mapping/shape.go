package mapping

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/shape.html#shape-mapping-options
type ShapeParams struct {
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