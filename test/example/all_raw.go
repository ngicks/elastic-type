package example

import (
	"net/netip"

	estype "github.com/ngicks/elastic-type/es_type"
)

type AllRaw struct {
	Agg             estype.Field[estype.AggregateMetricDouble] `json:"agg"`
	Alias           estype.Field[any]                          `json:"alias"`
	Blob            estype.Field[estype.Binary]                `json:"blob"`
	Bool            estype.Field[estype.Boolean]               `json:"bool"`
	Byte            estype.Field[int8]                         `json:"byte"`
	Comp            estype.Field[string]                       `json:"comp"`
	ConstantKwd     estype.Field[string]                       `json:"constant_kwd"`
	Date            estype.Field[AllDate]                      `json:"date"`
	DateNano        estype.Field[AllDateNano]                  `json:"dateNano"`
	DateRange       estype.Field[map[string]interface{}]       `json:"date_range"`
	DenseVector     estype.Field[float64]                      `json:"dense_vector"`
	Double          estype.Field[float64]                      `json:"double"`
	DoubleRange     estype.Field[map[string]interface{}]       `json:"double_range"`
	Flattened       estype.Field[map[string]interface{}]       `json:"flattened"`
	Float           estype.Field[float32]                      `json:"float"`
	FloatRange      estype.Field[map[string]interface{}]       `json:"float_range"`
	Geopoint        estype.Field[estype.Geopoint]              `json:"geopoint"`
	Geoshape        estype.Field[estype.Geoshape]              `json:"geoshape"`
	HalfFloat       estype.Field[float32]                      `json:"half_float"`
	Histogram       estype.Field[map[string]interface{}]       `json:"histogram"`
	Integer         estype.Field[int32]                        `json:"integer"`
	IntegerRange    estype.Field[map[string]interface{}]       `json:"integer_range"`
	IpAddr          estype.Field[netip.Addr]                   `json:"ip_addr"`
	IpRange         estype.Field[map[string]interface{}]       `json:"ip_range"`
	Join            estype.Field[map[string]interface{}]       `json:"join"`
	Kwd             estype.Field[string]                       `json:"kwd"`
	Long            estype.Field[int64]                        `json:"long"`
	LongRange       estype.Field[map[string]interface{}]       `json:"long_range"`
	Nested          estype.Field[AllNestedRaw]                 `json:"nested"`
	Object          estype.Field[AllObjectRaw]                 `json:"object"`
	Point           estype.Field[map[string]interface{}]       `json:"point"`
	Query           estype.Field[map[string]interface{}]       `json:"query"`
	RankFeature     estype.Field[float64]                      `json:"rank_feature"`
	RankFeatures    estype.Field[map[string]float64]           `json:"rank_features"`
	ScaledFloat     estype.Field[float64]                      `json:"scaled_float"`
	SearchAsYouType estype.Field[string]                       `json:"search_as_you_type"`
	Shape           estype.Field[estype.Geoshape]              `json:"shape"`
	Short           estype.Field[int16]                        `json:"short"`
	Text            estype.Field[string]                       `json:"text"`
	TextWTokenCount estype.Field[string]                       `json:"text_w_token_count"`
	UnsignedLong    estype.Field[uint64]                       `json:"unsigned_long"`
	Version         estype.Field[string]                       `json:"version"`
	Wildcard        estype.Field[string]                       `json:"wildcard"`
}

type AllObjectRaw struct {
	Age  estype.Field[int32]      `json:"age"`
	Name estype.Field[AllNameRaw] `json:"name"`
}

type AllNameRaw struct {
	First estype.Field[string] `json:"first"`
	Last  estype.Field[string] `json:"last"`
}

type AllNestedRaw struct {
	Age  estype.Field[int32]            `json:"age"`
	Name estype.Field[AllNestedNameRaw] `json:"name"`
}

type AllNestedNameRaw struct {
	First estype.Field[string] `json:"first"`
	Last  estype.Field[string] `json:"last"`
}
