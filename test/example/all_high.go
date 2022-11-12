package example

import (
	"encoding/json"
	"time"

	estype "github.com/ngicks/elastic-type/es_type"
	"github.com/ngicks/flextime"
	typeparamcommon "github.com/ngicks/type-param-common"
)

type All struct {
	Agg         *[]estype.AggregateMetricDouble `json:"agg"`
	Alias       *[]any                          `json:"alias"`
	Blob        *[]estype.Binary                `json:"blob"`
	Bool        *[]estype.Boolean               `json:"bool"`
	Comp        *[]string                       `json:"comp"`
	Date        *[]AllDate                      `json:"date"`
	DateNano    *[]AllDateNano                  `json:"dateNano"`
	DenseVector *[]float64                      `json:"dense_vector"`
	Flattened   *[]map[string]interface{}       `json:"flattened"`
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
