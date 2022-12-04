package generate

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func AggregateMetricDoubleParams(agg types.AggregateMetricDoubleProperty) GeneratedType {
	var min, max, sum, valueCount bool

	for _, v := range agg.Metrics {
		switch v {
		case "min":
			min = true
		case "max":
			max = true
		case "sum":
			sum = true
		case "value_count":
			valueCount = true
		}
	}

	var suffix string
	if min {
		suffix += "Min"
	}
	if max {
		suffix += "Max"
	}
	if sum {
		suffix += "Sum"
	}
	if valueCount {
		suffix += "ValueCount"
	}

	if min && max && sum && valueCount {
		suffix = ""
	}

	return GeneratedType{
		TyName:  "estype.AggregateMetricDouble" + suffix,
		Imports: estypeImport,
	}
}
