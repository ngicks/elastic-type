package mapping

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/aggregate-metric-double.html#aggregate-metric-double-params
type AggregateMetricDoubleParams struct {
	// Type is type of this property. Automatically filled if zero.
	Type          esType                             `json:"type"`
	Metrics       []aggregateMetricDoubleAggregation `json:"metrics"`
	DefaultMetric aggregateMetricDoubleAggregation   `json:"default_metric"`
}

func (p *AggregateMetricDoubleParams) FillType() {
	p.Type = AggregateMetricDouble
}

type aggregateMetricDoubleAggregation string

const (
	Min        aggregateMetricDoubleAggregation = "min"         // aggregation returns the minimum value of all min sub-fields.
	Max        aggregateMetricDoubleAggregation = "max"         // aggregation returns the maximum value of all max sub-fields.
	Sum        aggregateMetricDoubleAggregation = "sum"         // aggregation returns the sum of the values of all sum sub-fields.
	ValueCount aggregateMetricDoubleAggregation = "value_count" // aggregation returns the sum of the values of all value_count sub-fields.
)
