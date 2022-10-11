package mapping

// IndexSettings is main body for [Create index API.](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/indices-create-index.html)
type IndexSettings struct {
	// Aliases is irrevant for this package's goal.
	// But allowing it to be stored in here is maybe useful for some user.
	Aliases any `json:"aliases,omitempty"`
	// Settings is irrevant for this package's goal.
	// But allowing it to be stored in here is maybe useful for some user.
	Settings any       `json:"settings,omitempty"`
	Mappings *Mappings `json:"mappings,omitempty"`
}

// MappingSettings is response body of GET /<index_name>/_mapping.
//
// Map may only contain <index_name> as a key, and the contained IndexSettings may only have Mappings.
type MappingSettings map[string]IndexSettings

// Mappings is main body for [updating mapping](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/explicit-mapping.html#add-field-mapping),
// or a part of [Create index API](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/indices-create-index.html).
type Mappings struct {
	Properties Properties `json:"properties"`
}

type Properties map[string]any

func (p *Properties) FillType() {
	for _, v := range *p {
		if filler, ok := v.(FillTyper); ok {
			filler.FillType()
		}
	}
}

type FillTyper interface {
	// FillType fills Type field if it is zero value.
	FillType()
}

type onScriptError string

const (
	Continue onScriptError = "continue"
	Fail     onScriptError = "fail"
)
