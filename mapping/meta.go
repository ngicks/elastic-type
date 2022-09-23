package mapping

type metricType string

const (
	Gauge   metricType = "gauge"
	Counter metricType = "counter"
)

// see https://www.elastic.co/guide/en/elasticsearch/reference/8.4/mapping-field-meta.html
type Meta struct {
	Unit       *esUnit     `json:"unit,omitempty"`
	MetricType *metricType `json:"metric_type,omitempty"`
}
