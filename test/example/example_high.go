package example

import (
	"encoding/json"
	"time"

	estype "github.com/ngicks/elastic-type/es_type"
	"github.com/ngicks/flextime"
	typeparamcommon "github.com/ngicks/type-param-common"
)

type Example struct {
	Blob [][]byte        `json:"blob"`
	Bool *estype.Boolean `json:"bool"`
	Date ExampleDate     `json:"date"`
}

func (t Example) ToRaw() ExampleRaw {
	return ExampleRaw{
		Blob: estype.NewFieldSlice(t.Blob, false),
		Bool: estype.NewFieldSinglePointer(t.Bool, false),
		Date: estype.NewFieldSingleValue(t.Date),
	}
}

// ExampleDate represents elasticsearch date.
type ExampleDate time.Time

func (t ExampleDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

var parserExampleDate = flextime.NewFlextime(
	typeparamcommon.Must(flextime.NewLayoutSet(`2006-01-02TT15:04:05`)).
		AddLayout(typeparamcommon.Must(flextime.NewLayoutSet(`2006-01-02`))),
)

func (t *ExampleDate) UnmarshalJSON(data []byte) error {
	tt, err := estype.UnmarshalEsTime(
		data,
		parserExampleDate.Parse,
		time.UnixMilli,
	)
	if err != nil {
		return err
	}
	*t = ExampleDate(tt)
	return nil
}

func (t ExampleDate) String() string {
	return time.Time(t).Format(`2006-01-02TT15:04:05`)
}
