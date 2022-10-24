package mapping

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/field-alias.html
type AliasParams struct {
	// Type is type of this property. Automatically filled if zero.
	Type esType `json:"type,omitempty"`
	Path string `json:"path"`
}

func (p *AliasParams) FillType() {
	if p.Type == "" {
		p.Type = Alias
	}
}
