package mapping

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/parent-join.html
type JoinParams struct {
	// Type is type of this property. Automatically filled if zero.
	Type esType `json:"type,omitempty"`
	// Relations is mapping of parent (key) to child (value).
	// Value can be single string, or []string if the parent has multiple children.
	// Key can also be same string as one of child,
	// which mean there can be multiple level of child/parent relation.
	// Multi-level relation is not recommended however.
	Relations map[string]any `json:"relations"`
	// Defaults to true? (document is implicit)
	EagerGlobalOrdinals *bool `json:"eager_global_ordinals,omitempty"`
}

func (p *JoinParams) FillType() {
	if p.Type == "" {
		p.Type = Join
	}
}
