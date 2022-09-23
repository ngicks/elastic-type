package mapping

type esUnit string

// for time units: https://www.elastic.co/guide/en/elasticsearch/reference/8.4/api-conventions.html#time-units
// for other units: https://www.elastic.co/guide/en/elasticsearch/reference/8.4/mapping-field-meta.html
const (
	Days         esUnit = "d"
	Hours        esUnit = "h"
	Minutes      esUnit = "m"
	Seconds      esUnit = "s"
	Milliseconds esUnit = "ms"
	Microseconds esUnit = "micros"
	Nanoseconds  esUnit = "nanos"
	Percent      esUnit = "percent"
	UnitByte     esUnit = "byte"
)
