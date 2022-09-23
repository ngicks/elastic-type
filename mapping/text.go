package mapping

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/text.html#text-params
type TextParams struct {
	// Type is type of this property. Automatically filled if zero.
	Type esType `json:"type,omitempty"`
	// Defaults to default index analyzer or "standard".
	Analyzer *string `json:"analyzer,omitempty"`
	// Defaults to false.
	EagerGlobalOrdinals *bool `json:"eager_global_ordinals,omitempty"`
	// Defaults to false.
	Fielddata *bool `json:"fielddata,omitempty"`
	// FielddataFrequencyFilter decides a range that it loads into in-memoery when Fielddata is enabled.
	// By default it loads all values.
	FielddataFrequencyFilter *FielddataFrequencyFilter `json:"fielddata_frequency_filter,omitempty"`
	Fields                   *Fields                   `json:"fields,omitempty"`
	// Index indicates whether the field should be quickly searchable.
	// Numeric fields that only have doc_values enabled can also be queried, albeit slower.
	// Default(nil) is true.
	Index *bool `json:"index,omitempty"`
	// Defaults to Positions.
	IndexOptions  *indexOptions  `json:"index_options,omitempty"`
	IndexPrefixes *IndexPrefixes `json:"index_prefixes,omitempty"`
	// Defaults to false.
	IndexPhrases *bool `json:"index_phrases,omitempty"`
	// Defaults to true.
	Norms                *bool `json:"norms,omitempty"`
	PositionIncrementGap *int  `json:"position_increment_gap,omitempty"`
	// Default is false.
	Store *bool `json:"store,omitempty"`
	// Defaults to the Analyzer setting.
	SearchAnalyzer *string `json:"search_analyzer,omitempty"`
	// Defaults to the SearchAnalyzer setting.
	SearchQuoteAnalyzer *string `json:"search_quote_analyzer,omitempty"`
	// Defaults to "BM25".
	// Only "BM25" and "boolean" are available out-of-box.
	Similarity *string     `json:"similarity,omitempty"`
	TermVector *termVector `json:"term_vector,omitempty"`
	// Meta is metadata about the field.
	Meta *Meta `json:"meta,omitempty"`
}

func (p *TextParams) FillType() {
	p.Type = Text
}

type FielddataFrequencyFilter struct {
	Min            float64 `json:"min"`
	Max            float64 `json:"max"`
	MinSegmentSize float64 `json:"min_segment_size"`
}

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/index-prefixes.html
type IndexPrefixes struct {
	// Defaults to 2.
	// MinChars must be greater than 0.
	MinChars uint `json:"min_chars,omitempty"`
	// Defaults to 5.
	// MaxChars must be less than 20
	MaxChars uint `json:"max_chars,omitempty"`
}

type termVector string

const (
	No                           termVector = "no"
	Yes                          termVector = "yes"
	WithPositions                termVector = "with_positions"
	WithOffsets                  termVector = "with_offsets"
	WithPositionsOffsets         termVector = "with_positions_offsets"
	WithPositionsPayloads        termVector = "with_positions_payloads"
	WithPositionsOffsetspayloads termVector = "with_positions_offsets_payloads"
)
