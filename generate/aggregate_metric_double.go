package generate

import (
	"bytes"
	"html/template"

	"github.com/ngicks/elastic-type/mapping"
)

func GenerateAggregateMetricDoubleParams(agg mapping.AggregateMetricDoubleParams, tyeNamePrefix string) (tyName string, tyDef string) {
	buf := bytes.NewBuffer(make([]byte, 0))

	var min, max, sum, valueCount bool

	for _, v := range agg.Metrics {
		switch v {
		case mapping.Min:
			min = true
		case mapping.Max:
			max = true
		case mapping.Sum:
			sum = true
		case mapping.ValueCount:
			valueCount = true
		}
	}

	err := aggregateMetricDoubleType.Execute(buf, aggregateMetricDoubleTypeParam{
		Prefix:     tyeNamePrefix,
		Min:        min,
		Max:        max,
		Sum:        sum,
		ValueCount: valueCount,
	})
	if err != nil {
		panic(err)
	}

	return tyeNamePrefix + "AggregateMetricDouble", buf.String()
}

type aggregateMetricDoubleTypeParam struct {
	Prefix     string
	Min        bool
	Max        bool
	Sum        bool
	ValueCount bool
}

var aggregateMetricDoubleType = template.Must(template.New("aggregateMetricDoubleType").Parse(`
type {{.Prefix}}AggregateMetricDouble struct {
{{ if .Min}}	Min        float64 ` + "`" + `json:"min"` + "`" + `{{ end }}
{{ if .Max}}	Max        float64 ` + "`" + `json:"max"` + "`" + `{{ end }}
{{ if .Sum}}	Sum        float64 ` + "`" + `json:"sum"` + "`" + `{{ end }}
{{ if .ValueCount}}	ValueCount uint64  ` + "`" + `json:"value_count"` + "`" + `{{ end }}
}
`))
