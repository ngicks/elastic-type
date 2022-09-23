package mapping

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/version.html#version-params
type VersionParams struct {
	// Meta is metadata about the field.
	Meta *Meta `json:"meta,omitempty"`
}
