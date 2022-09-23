package mapping

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/completion.html#_parameters_for_completion_fields
type CompletionParams struct {
	// Defaults to simple.
	Analyzer *string `json:"analyzer,omitempty"`
	// Defaults to value of analyzer.
	SearchAnalyzer *string `json:"search_analyzer,omitempty"`
	// Defaults to true.
	PreserveSeparators *bool `json:"preserve_separators,omitempty"`
	// Defaults to true.
	PreservePositionIncrements *bool `json:"preserve_position_increments,omitempty"`
	// Defaults to 50 UTF-16 code points.
	MaxInputLength uint `json:"max_input_length,omitempty"`
}
