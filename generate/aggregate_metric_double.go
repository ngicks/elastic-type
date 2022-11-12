package generate

import (
	"github.com/ngicks/elastic-type/mapping"
)

func AggregateMetricDoubleParams(agg mapping.AggregateMetricDoubleParams) GeneratedType {
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
