package mapping

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/field-alias.html
type AliasParams struct {
	Path string `json:"path"`
}
