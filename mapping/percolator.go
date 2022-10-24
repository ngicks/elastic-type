package mapping

// Seemes nothing?
// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/percolator.html
type PercolatorParams struct {
	// Type is type of this property. Automatically filled if zero.
	Type esType `json:"type,omitempty"`
}

func (p *PercolatorParams) FillType() {
	if p.Type == "" {
		p.Type = Percolator
	}
}
