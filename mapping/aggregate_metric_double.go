package mapping

type aggregateMetricDoubleAggregation string

const (
	Min        aggregateMetricDoubleAggregation = "min"         // aggregation returns the minimum value of all min sub-fields.
	Max        aggregateMetricDoubleAggregation = "max"         // aggregation returns the maximum value of all max sub-fields.
	Sum        aggregateMetricDoubleAggregation = "sum"         // aggregation returns the sum of the values of all sum sub-fields.
	ValueCount aggregateMetricDoubleAggregation = "value_count" // aggregation returns the sum of the values of all value_count sub-fields.
)

type AggregateMetricDoubleParams struct {
	Metrics       []aggregateMetricDoubleAggregation `json:"metrics"`
	DefaultMetric aggregateMetricDoubleAggregation   `json:"default_metric"`
}
