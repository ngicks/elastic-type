package mapping

// Prams for RankFeature and RankFeatures
//
// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/rank-feature.html
type RankFeatureParams struct {
	// Defaults to true
	PositiveScoreImpact *bool `json:"positive_score_impact,omitempty"`
}
