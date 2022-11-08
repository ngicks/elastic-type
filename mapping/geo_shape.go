package mapping

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/geo-shape.html#geo-shape-mapping-options
type GeoshapeParams struct {
	// Type is type of this property. Automatically filled if zero.
	Type EsType `json:"type,omitempty"`
	// Defaults to "RIGHT".
	Orientation *orientation `json:"orientation,omitempty"`
	// IgnoreMalformed indicates whether it should ignore malformed value rather than rejecting whole document.
	// Default(nil) is false.
	IgnoreMalformed *bool `json:"ignore_malformed,omitempty"`
	// Default(nil) is true.
	IgnoreZValue *bool `json:"ignore_z_value,omitempty"`
	// If true, automatically close an unclosed polygon loop.
	// Default(nil) is false.
	Coerce *bool `json:"coerce,omitempty"`
}

func (p *GeoshapeParams) FillType() {
	if p.Type == "" {
		p.Type = Geoshape
	}
}

type orientation string

const (
	// RIGHT
	Right orientation = "right"
	// RIGHT
	Counterclockwise orientation = "counterclockwise"
	// RIGHT
	Ccw orientation = "ccw"
	// LEFT
	Left orientation = "left"
	// LEFT
	Clockwise orientation = "clockwise"
	// LEFT
	Cw orientation = "cw"
	// RIGHT
	RIGHT orientation = "RIGHT"
	// RIGHT
	COUNTERCLOCKWISE orientation = "COUNTERCLOCKWISE"
	// RIGHT
	CCW orientation = "CCW"
	// LEFT
	LEFT orientation = "LEFT"
	// LEFT
	CLOCKWISE orientation = "CLOCKWISE"
	// LEFT
	CW orientation = "CW"
)
