package mapping

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/object.html#object-params
type ObjectParams struct {
	// Dynamic can be bool(true/false/"true"/"false"), "runtime" or "strict".
	// Defaults to true.
	Dynamic *any `json:"dynamic,omitempty"`
	// Defaults to true.
	Enabled *bool `json:"enabled,omitempty"`
	// Defaults to true.
	Subobjects *bool       `json:"subobjects,omitempty"`
	Properties *Properties `json:"properties,omitempty"`
}
