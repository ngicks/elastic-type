package mapping

// Prams for RankFeature and RankFeatures
//
// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/rank-feature.html
type RankFeatureParams struct {
	// Type is type of this property. Automatically filled if zero.
	// Default is RankFeature.
	Type esType `json:"type,omitempty"`
	// Defaults to true
	PositiveScoreImpact *bool `json:"positive_score_impact,omitempty"`
}

func (p *RankFeatureParams) FillType() {
	if p.Type == "" {
		p.Type = RankFeature
	}
}
