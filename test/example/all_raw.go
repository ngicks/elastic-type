package example

import (
	estype "github.com/ngicks/elastic-type/es_type"
)

type AllRaw struct {
	Agg         estype.Field[estype.AggregateMetricDouble] `json:"agg"`
	Alias       estype.Field[any]                          `json:"alias"`
	Blob        estype.Field[estype.Binary]                `json:"blob"`
	Bool        estype.Field[estype.Boolean]               `json:"bool"`
	Comp        estype.Field[string]                       `json:"comp"`
	Date        estype.Field[AllDate]                      `json:"date"`
	DateNano    estype.Field[AllDateNano]                  `json:"dateNano"`
	DenseVector estype.Field[float64]                      `json:"dense_vector"`
	Flattened   estype.Field[map[string]interface{}]       `json:"flattened"`
}
