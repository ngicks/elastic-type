package mapping

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/version.html#version-params
type VersionParams struct {
	// Type is type of this property. Automatically filled if zero.
	Type EsType `json:"type,omitempty"`
	// Meta is metadata about the field.
	Meta *Meta `json:"meta,omitempty"`
}

func (p *VersionParams) FillType() {
	if p.Type == "" {
		p.Type = Version
	}
}
