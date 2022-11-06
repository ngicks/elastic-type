package mapping

// see https://www.elastic.co/guide/en/elasticsearch/reference/8.4/modules-scripting-using.html
type Script struct {
	// Defaults to "painless"
	Lang *string `json:"lang,omitempty"`
	// Use only either of source or id.
	Source *string         `json:"source,omitempty"`
	Id     *string         `json:"id,omitempty"`
	Params *map[string]any `json:"params,omitempty"`
}
