package mapping

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/nested.html#nested-params
type NestedParams struct {
	// Type is type of this property. Automatically filled if zero.
	Type EsType `json:"type,omitempty"`
	// Dynamic can be bool(true/false/"true"/"false") or "strict".
	// Defaults to true.
	Dynamic    *any        `json:"dynamic,omitempty"`
	Properties *Properties `json:"properties,omitempty"`
	// Defaults to false.
	IncludeInParent *bool `json:"include_in_parent,omitempty"`
	// Defaults to false.
	IncludeInRoot *bool `json:"include_in_root,omitempty"`
}

func (p *NestedParams) FillType() {
	if p.Type == "" {
		p.Type = Nested
	}
}
