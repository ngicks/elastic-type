package main

import (
	"bytes"
	"fmt"
	"text/template"
)

type aggregateMetricDoubleTypeParam struct {
	Suffix     string
	Min        bool
	Max        bool
	Sum        bool
	ValueCount bool
}

var aggregateMetricDoubleType = template.Must(template.New("aggregateMetricDoubleType").Parse(`
type AggregateMetricDouble{{.Suffix}} struct {
{{ if .Min}}	Min        float64 ` + "`" + `json:"min"` + "`" + `{{ end }}
{{ if .Max}}	Max        float64 ` + "`" + `json:"max"` + "`" + `{{ end }}
{{ if .Sum}}	Sum        float64 ` + "`" + `json:"sum"` + "`" + `{{ end }}
{{ if .ValueCount}}	ValueCount uint64  ` + "`" + `json:"value_count"` + "`" + `{{ end }}
}
`))

func main() {
	buf := bytes.NewBuffer(make([]byte, 0))

	for i := 0b0001; i <= 0b1111; i++ {
		var suffix string
		var min, max, sum, valueCount bool

		if i&0b0001 > 0 {
			suffix += "Min"
			min = true
		}
		if i&0b0010 > 0 {
			suffix += "Max"
			max = true
		}
		if i&0b0100 > 0 {
			suffix += "Sum"
			sum = true
		}
		if i&0b1000 > 0 {
			suffix += "ValueCount"
			valueCount = true
		}

		if min && max && sum && valueCount {
			suffix = ""
		}

		err := aggregateMetricDoubleType.Execute(buf, aggregateMetricDoubleTypeParam{
			Min:        min,
			Max:        max,
			Sum:        sum,
			ValueCount: valueCount,
			Suffix:     suffix,
		})
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(buf.String())
}
