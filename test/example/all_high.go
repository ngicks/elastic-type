package example

import (
	"encoding/json"
	"net/netip"
	"time"

	estype "github.com/ngicks/elastic-type/es_type"
	"github.com/ngicks/flextime"
	typeparamcommon "github.com/ngicks/type-param-common"
)

type All struct {
	Agg             *[]estype.AggregateMetricDouble `json:"agg"`
	Alias           *[]any                          `json:"alias"`
	Blob            *[]estype.Binary                `json:"blob"`
	Bool            *[]estype.Boolean               `json:"bool"`
	Byte            *[]int8                         `json:"byte"`
	Comp            *[]string                       `json:"comp"`
	ConstantKwd     *[]string                       `json:"constant_kwd"`
	Date            *[]AllDate                      `json:"date"`
	DateNano        *[]AllDateNano                  `json:"dateNano"`
	DateRange       *[]map[string]interface{}       `json:"date_range"`
	DenseVector     *[]float64                      `json:"dense_vector"`
	Double          *[]float64                      `json:"double"`
	DoubleRange     *[]map[string]interface{}       `json:"double_range"`
	Flattened       *[]map[string]interface{}       `json:"flattened"`
	Float           *[]float32                      `json:"float"`
	FloatRange      *[]map[string]interface{}       `json:"float_range"`
	Geopoint        *[]estype.Geopoint              `json:"geopoint"`
	Geoshape        *[]estype.Geoshape              `json:"geoshape"`
	HalfFloat       *[]float32                      `json:"half_float"`
	Histogram       *[]map[string]interface{}       `json:"histogram"`
	Integer         *[]int32                        `json:"integer"`
	IntegerRange    *[]map[string]interface{}       `json:"integer_range"`
	IpAddr          *[]netip.Addr                   `json:"ip_addr"`
	IpRange         *[]map[string]interface{}       `json:"ip_range"`
	Join            *[]map[string]interface{}       `json:"join"`
	Kwd             *[]string                       `json:"kwd"`
	Long            *[]int64                        `json:"long"`
	LongRange       *[]map[string]interface{}       `json:"long_range"`
	Nested          *[]AllNested                    `json:"nested"`
	Object          *[]AllObject                    `json:"object"`
	Point           *[]map[string]interface{}       `json:"point"`
	Query           *[]map[string]interface{}       `json:"query"`
	RankFeature     *[]float64                      `json:"rank_feature"`
	RankFeatures    *[]map[string]float64           `json:"rank_features"`
	ScaledFloat     *[]float64                      `json:"scaled_float"`
	SearchAsYouType *[]string                       `json:"search_as_you_type"`
	Shape           *[]estype.Geoshape              `json:"shape"`
	Short           *[]int16                        `json:"short"`
	Text            *[]string                       `json:"text"`
	TextWTokenCount *[]string                       `json:"text_w_token_count"`
	UnsignedLong    *[]uint64                       `json:"unsigned_long"`
	Version         *[]string                       `json:"version"`
	Wildcard        *[]string                       `json:"wildcard"`
}

type AllObject struct {
	Age  *[]int32   `json:"age"`
	Name *[]AllName `json:"name"`
}

type AllName struct {
	First *[]string `json:"first"`
	Last  *[]string `json:"last"`
}

type AllNested struct {
	Age  *[]int32         `json:"age"`
	Name *[]AllNestedName `json:"name"`
}

type AllNestedName struct {
	First *[]string `json:"first"`
	Last  *[]string `json:"last"`
}

// AllDateNano represents elasticsearch date.
type AllDateNano time.Time

func (t AllDateNano) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

var parserAllDateNano = flextime.NewFlextime(
	typeparamcommon.Must(flextime.NewLayoutSet(`2006-01-02T15:04:05.999999999Z07:00`)).
		AddLayout(typeparamcommon.Must(flextime.NewLayoutSet(`2006-01-02`))),
)

func (t *AllDateNano) UnmarshalJSON(data []byte) error {
	tt, err := estype.UnmarshalEsTime(
		data,
		parserAllDateNano.Parse,
		func(v int64) time.Time { return time.Unix(v, 0) },
	)
	if err != nil {
		return err
	}
	*t = AllDateNano(tt)
	return nil
}

func (t AllDateNano) String() string {
	return time.Time(t).Format(`2006-01-02T15:04:05.999999999Z07:00`)
}

// AllDate represents elasticsearch date.
type AllDate time.Time

func (t AllDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

var parserAllDate = flextime.NewFlextime(
	typeparamcommon.Must(flextime.NewLayoutSet(`2006-01-02 15:04:05`)).
		AddLayout(typeparamcommon.Must(flextime.NewLayoutSet(`2006-01-02`))),
)

func (t *AllDate) UnmarshalJSON(data []byte) error {
	tt, err := estype.UnmarshalEsTime(
		data,
		parserAllDate.Parse,
		time.UnixMilli,
	)
	if err != nil {
		return err
	}
	*t = AllDate(tt)
	return nil
}

func (t AllDate) String() string {
	return time.Time(t).Format(`2006-01-02 15:04:05`)
}
