package mapping

// Nothing!
// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/histogram.html#histogram-params
type HistogramParams struct {
	// Type is type of this property. Automatically filled if zero.
	Type esType `json:"type,omitempty"`
}

func (p *HistogramParams) FillType() {
	if p.Type == "" {
		p.Type = Histogram
	}
}
