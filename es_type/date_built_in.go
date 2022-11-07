package estype

import (
	"encoding/json"
	"strconv"
	"time"

	builtinformat "github.com/ngicks/elastic-type/es_type/builtin_format"
)

type StrictDateOptionalTimeEpochMillis time.Time

func (t StrictDateOptionalTimeEpochMillis) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictDateOptionalTimeEpochMillis) UnmarshalJSON(data []byte) error {
	bb, err := UnmarshalEsTime(data, builtinformat.Formatters[builtinformat.StrictDateOptionalTime].Parse, time.UnixMilli)
	if err != nil {
		return err
	}
	*t = StrictDateOptionalTimeEpochMillis(bb)
	return nil
}

func (t StrictDateOptionalTimeEpochMillis) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictDateOptionalTime])
}

type StrictDateOptionalTimeNanosEpochMillis time.Time

func (t StrictDateOptionalTimeNanosEpochMillis) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictDateOptionalTimeNanosEpochMillis) UnmarshalJSON(data []byte) error {
	bb, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictDateOptionalTimeNanos].Parse,
		time.UnixMilli,
	)
	if err != nil {
		return err
	}
	*t = StrictDateOptionalTimeNanosEpochMillis(bb)
	return nil
}

func (t StrictDateOptionalTimeNanosEpochMillis) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictDateOptionalTimeNanos])
}

type EpochMillis time.Time

func (t EpochMillis) MarshalJSON() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *EpochMillis) UnmarshalJSON(data []byte) error {
	bb, err := UnmarshalEsTime(data, nil, time.UnixMilli)
	if err != nil {
		return err
	}
	*t = EpochMillis(bb)
	return nil
}

func (t EpochMillis) String() string {
	return strconv.FormatInt(time.Time(t).UnixMilli(), 10)
}

type EpochSecond time.Time

func (t EpochSecond) MarshalJSON() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *EpochSecond) UnmarshalJSON(data []byte) error {
	bb, err := UnmarshalEsTime(data, nil, ParseUnixSec, `UnixSec`)
	if err != nil {
		return err
	}
	*t = EpochSecond(bb)
	return nil
}

func (t EpochSecond) String() string {
	return strconv.FormatInt(time.Time(t).Unix(), 10)
}

// generate_date:start

type DateOptionalTime time.Time

func (t DateOptionalTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *DateOptionalTime) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.DateOptionalTime].Parse,
		nil,
		`DateOptionalTime`,
	)
	if err != nil {
		return err
	}
	*t = DateOptionalTime(tt)
	return nil
}

func (t DateOptionalTime) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.DateOptionalTime])
}

type StrictDateOptionalTime time.Time

func (t StrictDateOptionalTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictDateOptionalTime) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictDateOptionalTime].Parse,
		nil,
		`StrictDateOptionalTime`,
	)
	if err != nil {
		return err
	}
	*t = StrictDateOptionalTime(tt)
	return nil
}

func (t StrictDateOptionalTime) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictDateOptionalTime])
}

type StrictDateOptionalTimeNanos time.Time

func (t StrictDateOptionalTimeNanos) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictDateOptionalTimeNanos) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictDateOptionalTimeNanos].Parse,
		nil,
		`StrictDateOptionalTimeNanos`,
	)
	if err != nil {
		return err
	}
	*t = StrictDateOptionalTimeNanos(tt)
	return nil
}

func (t StrictDateOptionalTimeNanos) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictDateOptionalTimeNanos])
}

type BasicDate time.Time

func (t BasicDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *BasicDate) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.BasicDate].Parse,
		nil,
		`BasicDate`,
	)
	if err != nil {
		return err
	}
	*t = BasicDate(tt)
	return nil
}

func (t BasicDate) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.BasicDate])
}

type BasicDateTime time.Time

func (t BasicDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *BasicDateTime) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.BasicDateTime].Parse,
		nil,
		`BasicDateTime`,
	)
	if err != nil {
		return err
	}
	*t = BasicDateTime(tt)
	return nil
}

func (t BasicDateTime) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.BasicDateTime])
}

type BasicDateTimeNoMillis time.Time

func (t BasicDateTimeNoMillis) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *BasicDateTimeNoMillis) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.BasicDateTimeNoMillis].Parse,
		nil,
		`BasicDateTimeNoMillis`,
	)
	if err != nil {
		return err
	}
	*t = BasicDateTimeNoMillis(tt)
	return nil
}

func (t BasicDateTimeNoMillis) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.BasicDateTimeNoMillis])
}

type BasicOrdinalDate time.Time

func (t BasicOrdinalDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *BasicOrdinalDate) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.BasicOrdinalDate].Parse,
		nil,
		`BasicOrdinalDate`,
	)
	if err != nil {
		return err
	}
	*t = BasicOrdinalDate(tt)
	return nil
}

func (t BasicOrdinalDate) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.BasicOrdinalDate])
}

type BasicOrdinalDateTime time.Time

func (t BasicOrdinalDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *BasicOrdinalDateTime) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.BasicOrdinalDateTime].Parse,
		nil,
		`BasicOrdinalDateTime`,
	)
	if err != nil {
		return err
	}
	*t = BasicOrdinalDateTime(tt)
	return nil
}

func (t BasicOrdinalDateTime) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.BasicOrdinalDateTime])
}

type BasicOrdinalDateTimeNoMillis time.Time

func (t BasicOrdinalDateTimeNoMillis) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *BasicOrdinalDateTimeNoMillis) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.BasicOrdinalDateTimeNoMillis].Parse,
		nil,
		`BasicOrdinalDateTimeNoMillis`,
	)
	if err != nil {
		return err
	}
	*t = BasicOrdinalDateTimeNoMillis(tt)
	return nil
}

func (t BasicOrdinalDateTimeNoMillis) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.BasicOrdinalDateTimeNoMillis])
}

type BasicTime time.Time

func (t BasicTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *BasicTime) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.BasicTime].Parse,
		nil,
		`BasicTime`,
	)
	if err != nil {
		return err
	}
	*t = BasicTime(tt)
	return nil
}

func (t BasicTime) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.BasicTime])
}

type BasicTimeNoMillis time.Time

func (t BasicTimeNoMillis) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *BasicTimeNoMillis) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.BasicTimeNoMillis].Parse,
		nil,
		`BasicTimeNoMillis`,
	)
	if err != nil {
		return err
	}
	*t = BasicTimeNoMillis(tt)
	return nil
}

func (t BasicTimeNoMillis) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.BasicTimeNoMillis])
}

type BasicTTime time.Time

func (t BasicTTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *BasicTTime) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.BasicTTime].Parse,
		nil,
		`BasicTTime`,
	)
	if err != nil {
		return err
	}
	*t = BasicTTime(tt)
	return nil
}

func (t BasicTTime) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.BasicTTime])
}

type BasicTTimeNoMillis time.Time

func (t BasicTTimeNoMillis) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *BasicTTimeNoMillis) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.BasicTTimeNoMillis].Parse,
		nil,
		`BasicTTimeNoMillis`,
	)
	if err != nil {
		return err
	}
	*t = BasicTTimeNoMillis(tt)
	return nil
}

func (t BasicTTimeNoMillis) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.BasicTTimeNoMillis])
}

type BasicWeekDate time.Time

func (t BasicWeekDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *BasicWeekDate) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.BasicWeekDate].Parse,
		nil,
		`BasicWeekDate`,
	)
	if err != nil {
		return err
	}
	*t = BasicWeekDate(tt)
	return nil
}

func (t BasicWeekDate) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.BasicWeekDate])
}

type StrictBasicWeekDate time.Time

func (t StrictBasicWeekDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictBasicWeekDate) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictBasicWeekDate].Parse,
		nil,
		`StrictBasicWeekDate`,
	)
	if err != nil {
		return err
	}
	*t = StrictBasicWeekDate(tt)
	return nil
}

func (t StrictBasicWeekDate) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictBasicWeekDate])
}

type BasicWeekDateTime time.Time

func (t BasicWeekDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *BasicWeekDateTime) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.BasicWeekDateTime].Parse,
		nil,
		`BasicWeekDateTime`,
	)
	if err != nil {
		return err
	}
	*t = BasicWeekDateTime(tt)
	return nil
}

func (t BasicWeekDateTime) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.BasicWeekDateTime])
}

type StrictBasicWeekDateTime time.Time

func (t StrictBasicWeekDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictBasicWeekDateTime) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictBasicWeekDateTime].Parse,
		nil,
		`StrictBasicWeekDateTime`,
	)
	if err != nil {
		return err
	}
	*t = StrictBasicWeekDateTime(tt)
	return nil
}

func (t StrictBasicWeekDateTime) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictBasicWeekDateTime])
}

type BasicWeekDateTimeNoMillis time.Time

func (t BasicWeekDateTimeNoMillis) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *BasicWeekDateTimeNoMillis) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.BasicWeekDateTimeNoMillis].Parse,
		nil,
		`BasicWeekDateTimeNoMillis`,
	)
	if err != nil {
		return err
	}
	*t = BasicWeekDateTimeNoMillis(tt)
	return nil
}

func (t BasicWeekDateTimeNoMillis) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.BasicWeekDateTimeNoMillis])
}

type StrictBasicWeekDateTimeNoMillis time.Time

func (t StrictBasicWeekDateTimeNoMillis) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictBasicWeekDateTimeNoMillis) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictBasicWeekDateTimeNoMillis].Parse,
		nil,
		`StrictBasicWeekDateTimeNoMillis`,
	)
	if err != nil {
		return err
	}
	*t = StrictBasicWeekDateTimeNoMillis(tt)
	return nil
}

func (t StrictBasicWeekDateTimeNoMillis) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictBasicWeekDateTimeNoMillis])
}

type Date time.Time

func (t Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *Date) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.Date].Parse,
		nil,
		`Date`,
	)
	if err != nil {
		return err
	}
	*t = Date(tt)
	return nil
}

func (t Date) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.Date])
}

type StrictDate time.Time

func (t StrictDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictDate) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictDate].Parse,
		nil,
		`StrictDate`,
	)
	if err != nil {
		return err
	}
	*t = StrictDate(tt)
	return nil
}

func (t StrictDate) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictDate])
}

type DateHour time.Time

func (t DateHour) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *DateHour) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.DateHour].Parse,
		nil,
		`DateHour`,
	)
	if err != nil {
		return err
	}
	*t = DateHour(tt)
	return nil
}

func (t DateHour) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.DateHour])
}

type StrictDateHour time.Time

func (t StrictDateHour) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictDateHour) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictDateHour].Parse,
		nil,
		`StrictDateHour`,
	)
	if err != nil {
		return err
	}
	*t = StrictDateHour(tt)
	return nil
}

func (t StrictDateHour) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictDateHour])
}

type DateHourMinute time.Time

func (t DateHourMinute) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *DateHourMinute) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.DateHourMinute].Parse,
		nil,
		`DateHourMinute`,
	)
	if err != nil {
		return err
	}
	*t = DateHourMinute(tt)
	return nil
}

func (t DateHourMinute) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.DateHourMinute])
}

type StrictDateHourMinute time.Time

func (t StrictDateHourMinute) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictDateHourMinute) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictDateHourMinute].Parse,
		nil,
		`StrictDateHourMinute`,
	)
	if err != nil {
		return err
	}
	*t = StrictDateHourMinute(tt)
	return nil
}

func (t StrictDateHourMinute) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictDateHourMinute])
}

type DateHourMinuteSecond time.Time

func (t DateHourMinuteSecond) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *DateHourMinuteSecond) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.DateHourMinuteSecond].Parse,
		nil,
		`DateHourMinuteSecond`,
	)
	if err != nil {
		return err
	}
	*t = DateHourMinuteSecond(tt)
	return nil
}

func (t DateHourMinuteSecond) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.DateHourMinuteSecond])
}

type StrictDateHourMinuteSecond time.Time

func (t StrictDateHourMinuteSecond) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictDateHourMinuteSecond) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictDateHourMinuteSecond].Parse,
		nil,
		`StrictDateHourMinuteSecond`,
	)
	if err != nil {
		return err
	}
	*t = StrictDateHourMinuteSecond(tt)
	return nil
}

func (t StrictDateHourMinuteSecond) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictDateHourMinuteSecond])
}

type DateHourMinuteSecondFraction time.Time

func (t DateHourMinuteSecondFraction) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *DateHourMinuteSecondFraction) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.DateHourMinuteSecondFraction].Parse,
		nil,
		`DateHourMinuteSecondFraction`,
	)
	if err != nil {
		return err
	}
	*t = DateHourMinuteSecondFraction(tt)
	return nil
}

func (t DateHourMinuteSecondFraction) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.DateHourMinuteSecondFraction])
}

type StrictDateHourMinuteSecondFraction time.Time

func (t StrictDateHourMinuteSecondFraction) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictDateHourMinuteSecondFraction) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictDateHourMinuteSecondFraction].Parse,
		nil,
		`StrictDateHourMinuteSecondFraction`,
	)
	if err != nil {
		return err
	}
	*t = StrictDateHourMinuteSecondFraction(tt)
	return nil
}

func (t StrictDateHourMinuteSecondFraction) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictDateHourMinuteSecondFraction])
}

type DateHourMinuteSecondMillis time.Time

func (t DateHourMinuteSecondMillis) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *DateHourMinuteSecondMillis) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.DateHourMinuteSecondMillis].Parse,
		nil,
		`DateHourMinuteSecondMillis`,
	)
	if err != nil {
		return err
	}
	*t = DateHourMinuteSecondMillis(tt)
	return nil
}

func (t DateHourMinuteSecondMillis) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.DateHourMinuteSecondMillis])
}

type StrictDateHourMinuteSecondMillis time.Time

func (t StrictDateHourMinuteSecondMillis) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictDateHourMinuteSecondMillis) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictDateHourMinuteSecondMillis].Parse,
		nil,
		`StrictDateHourMinuteSecondMillis`,
	)
	if err != nil {
		return err
	}
	*t = StrictDateHourMinuteSecondMillis(tt)
	return nil
}

func (t StrictDateHourMinuteSecondMillis) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictDateHourMinuteSecondMillis])
}

type DateTime time.Time

func (t DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *DateTime) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.DateTime].Parse,
		nil,
		`DateTime`,
	)
	if err != nil {
		return err
	}
	*t = DateTime(tt)
	return nil
}

func (t DateTime) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.DateTime])
}

type StrictDateTime time.Time

func (t StrictDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictDateTime) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictDateTime].Parse,
		nil,
		`StrictDateTime`,
	)
	if err != nil {
		return err
	}
	*t = StrictDateTime(tt)
	return nil
}

func (t StrictDateTime) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictDateTime])
}

type DateTimeNoMillis time.Time

func (t DateTimeNoMillis) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *DateTimeNoMillis) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.DateTimeNoMillis].Parse,
		nil,
		`DateTimeNoMillis`,
	)
	if err != nil {
		return err
	}
	*t = DateTimeNoMillis(tt)
	return nil
}

func (t DateTimeNoMillis) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.DateTimeNoMillis])
}

type StrictDateTimeNoMillis time.Time

func (t StrictDateTimeNoMillis) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictDateTimeNoMillis) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictDateTimeNoMillis].Parse,
		nil,
		`StrictDateTimeNoMillis`,
	)
	if err != nil {
		return err
	}
	*t = StrictDateTimeNoMillis(tt)
	return nil
}

func (t StrictDateTimeNoMillis) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictDateTimeNoMillis])
}

type Hour time.Time

func (t Hour) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *Hour) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.Hour].Parse,
		nil,
		`Hour`,
	)
	if err != nil {
		return err
	}
	*t = Hour(tt)
	return nil
}

func (t Hour) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.Hour])
}

type StrictHour time.Time

func (t StrictHour) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictHour) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictHour].Parse,
		nil,
		`StrictHour`,
	)
	if err != nil {
		return err
	}
	*t = StrictHour(tt)
	return nil
}

func (t StrictHour) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictHour])
}

type HourMinute time.Time

func (t HourMinute) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *HourMinute) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.HourMinute].Parse,
		nil,
		`HourMinute`,
	)
	if err != nil {
		return err
	}
	*t = HourMinute(tt)
	return nil
}

func (t HourMinute) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.HourMinute])
}

type StrictHourMinute time.Time

func (t StrictHourMinute) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictHourMinute) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictHourMinute].Parse,
		nil,
		`StrictHourMinute`,
	)
	if err != nil {
		return err
	}
	*t = StrictHourMinute(tt)
	return nil
}

func (t StrictHourMinute) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictHourMinute])
}

type HourMinuteSecond time.Time

func (t HourMinuteSecond) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *HourMinuteSecond) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.HourMinuteSecond].Parse,
		nil,
		`HourMinuteSecond`,
	)
	if err != nil {
		return err
	}
	*t = HourMinuteSecond(tt)
	return nil
}

func (t HourMinuteSecond) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.HourMinuteSecond])
}

type StrictHourMinuteSecond time.Time

func (t StrictHourMinuteSecond) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictHourMinuteSecond) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictHourMinuteSecond].Parse,
		nil,
		`StrictHourMinuteSecond`,
	)
	if err != nil {
		return err
	}
	*t = StrictHourMinuteSecond(tt)
	return nil
}

func (t StrictHourMinuteSecond) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictHourMinuteSecond])
}

type HourMinuteSecondFraction time.Time

func (t HourMinuteSecondFraction) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *HourMinuteSecondFraction) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.HourMinuteSecondFraction].Parse,
		nil,
		`HourMinuteSecondFraction`,
	)
	if err != nil {
		return err
	}
	*t = HourMinuteSecondFraction(tt)
	return nil
}

func (t HourMinuteSecondFraction) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.HourMinuteSecondFraction])
}

type StrictHourMinuteSecondFraction time.Time

func (t StrictHourMinuteSecondFraction) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictHourMinuteSecondFraction) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictHourMinuteSecondFraction].Parse,
		nil,
		`StrictHourMinuteSecondFraction`,
	)
	if err != nil {
		return err
	}
	*t = StrictHourMinuteSecondFraction(tt)
	return nil
}

func (t StrictHourMinuteSecondFraction) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictHourMinuteSecondFraction])
}

type HourMinuteSecondMillis time.Time

func (t HourMinuteSecondMillis) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *HourMinuteSecondMillis) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.HourMinuteSecondMillis].Parse,
		nil,
		`HourMinuteSecondMillis`,
	)
	if err != nil {
		return err
	}
	*t = HourMinuteSecondMillis(tt)
	return nil
}

func (t HourMinuteSecondMillis) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.HourMinuteSecondMillis])
}

type StrictHourMinuteSecondMillis time.Time

func (t StrictHourMinuteSecondMillis) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictHourMinuteSecondMillis) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictHourMinuteSecondMillis].Parse,
		nil,
		`StrictHourMinuteSecondMillis`,
	)
	if err != nil {
		return err
	}
	*t = StrictHourMinuteSecondMillis(tt)
	return nil
}

func (t StrictHourMinuteSecondMillis) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictHourMinuteSecondMillis])
}

type OrdinalDate time.Time

func (t OrdinalDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *OrdinalDate) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.OrdinalDate].Parse,
		nil,
		`OrdinalDate`,
	)
	if err != nil {
		return err
	}
	*t = OrdinalDate(tt)
	return nil
}

func (t OrdinalDate) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.OrdinalDate])
}

type StrictOrdinalDate time.Time

func (t StrictOrdinalDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictOrdinalDate) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictOrdinalDate].Parse,
		nil,
		`StrictOrdinalDate`,
	)
	if err != nil {
		return err
	}
	*t = StrictOrdinalDate(tt)
	return nil
}

func (t StrictOrdinalDate) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictOrdinalDate])
}

type OrdinalDateTime time.Time

func (t OrdinalDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *OrdinalDateTime) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.OrdinalDateTime].Parse,
		nil,
		`OrdinalDateTime`,
	)
	if err != nil {
		return err
	}
	*t = OrdinalDateTime(tt)
	return nil
}

func (t OrdinalDateTime) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.OrdinalDateTime])
}

type StrictOrdinalDateTime time.Time

func (t StrictOrdinalDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictOrdinalDateTime) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictOrdinalDateTime].Parse,
		nil,
		`StrictOrdinalDateTime`,
	)
	if err != nil {
		return err
	}
	*t = StrictOrdinalDateTime(tt)
	return nil
}

func (t StrictOrdinalDateTime) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictOrdinalDateTime])
}

type OrdinalDateTimeNoMillis time.Time

func (t OrdinalDateTimeNoMillis) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *OrdinalDateTimeNoMillis) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.OrdinalDateTimeNoMillis].Parse,
		nil,
		`OrdinalDateTimeNoMillis`,
	)
	if err != nil {
		return err
	}
	*t = OrdinalDateTimeNoMillis(tt)
	return nil
}

func (t OrdinalDateTimeNoMillis) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.OrdinalDateTimeNoMillis])
}

type StrictOrdinalDateTimeNoMillis time.Time

func (t StrictOrdinalDateTimeNoMillis) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictOrdinalDateTimeNoMillis) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictOrdinalDateTimeNoMillis].Parse,
		nil,
		`StrictOrdinalDateTimeNoMillis`,
	)
	if err != nil {
		return err
	}
	*t = StrictOrdinalDateTimeNoMillis(tt)
	return nil
}

func (t StrictOrdinalDateTimeNoMillis) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictOrdinalDateTimeNoMillis])
}

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *Time) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.Time].Parse,
		nil,
		`Time`,
	)
	if err != nil {
		return err
	}
	*t = Time(tt)
	return nil
}

func (t Time) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.Time])
}

type StrictTime time.Time

func (t StrictTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictTime) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictTime].Parse,
		nil,
		`StrictTime`,
	)
	if err != nil {
		return err
	}
	*t = StrictTime(tt)
	return nil
}

func (t StrictTime) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictTime])
}

type TimeNoMillis time.Time

func (t TimeNoMillis) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *TimeNoMillis) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.TimeNoMillis].Parse,
		nil,
		`TimeNoMillis`,
	)
	if err != nil {
		return err
	}
	*t = TimeNoMillis(tt)
	return nil
}

func (t TimeNoMillis) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.TimeNoMillis])
}

type StrictTimeNoMillis time.Time

func (t StrictTimeNoMillis) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictTimeNoMillis) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictTimeNoMillis].Parse,
		nil,
		`StrictTimeNoMillis`,
	)
	if err != nil {
		return err
	}
	*t = StrictTimeNoMillis(tt)
	return nil
}

func (t StrictTimeNoMillis) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictTimeNoMillis])
}

type TTime time.Time

func (t TTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *TTime) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.TTime].Parse,
		nil,
		`TTime`,
	)
	if err != nil {
		return err
	}
	*t = TTime(tt)
	return nil
}

func (t TTime) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.TTime])
}

type StrictTTime time.Time

func (t StrictTTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictTTime) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictTTime].Parse,
		nil,
		`StrictTTime`,
	)
	if err != nil {
		return err
	}
	*t = StrictTTime(tt)
	return nil
}

func (t StrictTTime) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictTTime])
}

type TTimeNoMillis time.Time

func (t TTimeNoMillis) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *TTimeNoMillis) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.TTimeNoMillis].Parse,
		nil,
		`TTimeNoMillis`,
	)
	if err != nil {
		return err
	}
	*t = TTimeNoMillis(tt)
	return nil
}

func (t TTimeNoMillis) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.TTimeNoMillis])
}

type StrictTTimeNoMillis time.Time

func (t StrictTTimeNoMillis) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictTTimeNoMillis) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictTTimeNoMillis].Parse,
		nil,
		`StrictTTimeNoMillis`,
	)
	if err != nil {
		return err
	}
	*t = StrictTTimeNoMillis(tt)
	return nil
}

func (t StrictTTimeNoMillis) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictTTimeNoMillis])
}

type Year time.Time

func (t Year) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *Year) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.Year].Parse,
		nil,
		`Year`,
	)
	if err != nil {
		return err
	}
	*t = Year(tt)
	return nil
}

func (t Year) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.Year])
}

type StrictYear time.Time

func (t StrictYear) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictYear) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictYear].Parse,
		nil,
		`StrictYear`,
	)
	if err != nil {
		return err
	}
	*t = StrictYear(tt)
	return nil
}

func (t StrictYear) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictYear])
}

type YearMonth time.Time

func (t YearMonth) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *YearMonth) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.YearMonth].Parse,
		nil,
		`YearMonth`,
	)
	if err != nil {
		return err
	}
	*t = YearMonth(tt)
	return nil
}

func (t YearMonth) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.YearMonth])
}

type StrictYearMonth time.Time

func (t StrictYearMonth) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictYearMonth) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictYearMonth].Parse,
		nil,
		`StrictYearMonth`,
	)
	if err != nil {
		return err
	}
	*t = StrictYearMonth(tt)
	return nil
}

func (t StrictYearMonth) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictYearMonth])
}

type YearMonthDay time.Time

func (t YearMonthDay) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *YearMonthDay) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.YearMonthDay].Parse,
		nil,
		`YearMonthDay`,
	)
	if err != nil {
		return err
	}
	*t = YearMonthDay(tt)
	return nil
}

func (t YearMonthDay) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.YearMonthDay])
}

type StrictYearMonthDay time.Time

func (t StrictYearMonthDay) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *StrictYearMonthDay) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.StrictYearMonthDay].Parse,
		nil,
		`StrictYearMonthDay`,
	)
	if err != nil {
		return err
	}
	*t = StrictYearMonthDay(tt)
	return nil
}

func (t StrictYearMonthDay) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.StrictYearMonthDay])
}

// generate_date:end
