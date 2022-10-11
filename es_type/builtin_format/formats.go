package builtinformat

import (
	"fmt"
	"os"

	"github.com/ngicks/flextime"
)

const DefaultFormat = "strict_date_optional_time||epoch_millis"

// Built-in formats
//
// Formats that contains weekyear is not supported by this module, although they are listed.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/mapping-date-format.html#strict-date-time
const (
	EpochMillis                        = "epoch_millis"
	EpochSecond                        = "epoch_second"
	DateOptionalTime                   = "date_optional_time"
	StrictDateOptionalTime             = "strict_date_optional_time"
	StrictDateOptionalTimeNanos        = "strict_date_optional_time_nanos"
	BasicDate                          = "basic_date"
	BasicDateTime                      = "basic_date_time"
	BasicDateTimeNoMillis              = "basic_date_time_no_millis"
	BasicOrdinalDate                   = "basic_ordinal_date"
	BasicOrdinalDateTime               = "basic_ordinal_date_time"
	BasicOrdinalDateTimeNoMillis       = "basic_ordinal_date_time_no_millis"
	BasicTime                          = "basic_time"
	BasicTimeNoMillis                  = "basic_time_no_millis"
	BasicTTime                         = "basic_t_time"
	BasicTTimeNoMillis                 = "basic_t_time_no_millis"
	BasicWeekDate                      = "basic_week_date"
	StrictBasicWeekDate                = "strict_basic_week_date"
	BasicWeekDateTime                  = "basic_week_date_time"
	StrictBasicWeekDateTime            = "strict_basic_week_date_time"
	BasicWeekDateTimeNoMillis          = "basic_week_date_time_no_millis"
	StrictBasicWeekDateTimeNoMillis    = "strict_basic_week_date_time_no_millis"
	Date                               = "date"
	StrictDate                         = "strict_date"
	DateHour                           = "date_hour"
	StrictDateHour                     = "strict_date_hour"
	DateHourMinute                     = "date_hour_minute"
	StrictDateHourMinute               = "strict_date_hour_minute"
	DateHourMinuteSecond               = "date_hour_minute_second"
	StrictDateHourMinuteSecond         = "strict_date_hour_minute_second"
	DateHourMinuteSecondFraction       = "date_hour_minute_second_fraction"
	StrictDateHourMinuteSecondFraction = "strict_date_hour_minute_second_fraction"
	DateHourMinuteSecondMillis         = "date_hour_minute_second_millis"
	StrictDateHourMinuteSecondMillis   = "strict_date_hour_minute_second_millis"
	DateTime                           = "date_time"
	StrictDateTime                     = "strict_date_time"
	DateTimeNoMillis                   = "date_time_no_millis"
	StrictDateTimeNoMillis             = "strict_date_time_no_millis"
	Hour                               = "hour"
	StrictHour                         = "strict_hour"
	HourMinute                         = "hour_minute"
	StrictHourMinute                   = "strict_hour_minute"
	HourMinuteSecond                   = "hour_minute_second"
	StrictHourMinuteSecond             = "strict_hour_minute_second"
	HourMinuteSecondFraction           = "hour_minute_second_fraction"
	StrictHourMinuteSecondFraction     = "strict_hour_minute_second_fraction"
	HourMinuteSecondMillis             = "hour_minute_second_millis"
	StrictHourMinuteSecondMillis       = "strict_hour_minute_second_millis"
	OrdinalDate                        = "ordinal_date"
	StrictOrdinalDate                  = "strict_ordinal_date"
	OrdinalDateTime                    = "ordinal_date_time"
	StrictOrdinalDateTime              = "strict_ordinal_date_time"
	OrdinalDateTimeNoMillis            = "ordinal_date_time_no_millis"
	StrictOrdinalDateTimeNoMillis      = "strict_ordinal_date_time_no_millis"
	Time                               = "time"
	StrictTime                         = "strict_time"
	TimeNoMillis                       = "time_no_millis"
	StrictTimeNoMillis                 = "strict_time_no_millis"
	TTime                              = "t_time"
	StrictTTime                        = "strict_t_time"
	TTimeNoMillis                      = "t_time_no_millis"
	StrictTTimeNoMillis                = "strict_t_time_no_millis"
	WeekDate                           = "week_date"
	StrictWeekDate                     = "strict_week_date"
	WeekDateTime                       = "week_date_time"
	StrictWeekDateTime                 = "strict_week_date_time"
	WeekDateTimeNoMillis               = "week_date_time_no_millis"
	StrictWeekDateTimeNoMillis         = "strict_week_date_time_no_millis"
	Weekyear                           = "weekyear"
	StrictWeekyear                     = "strict_weekyear"
	WeekyearWeek                       = "weekyear_week"
	StrictWeekyearWeek                 = "strict_weekyear_week"
	WeekyearWeekDay                    = "weekyear_week_day"
	StrictWeekyearWeekDay              = "strict_weekyear_week_day"
	Year                               = "year"
	StrictYear                         = "strict_year"
	YearMonth                          = "year_month"
	StrictYearMonth                    = "strict_year_month"
	YearMonthDay                       = "year_month_day"
	StrictYearMonthDay                 = "strict_year_month_day"
)

// ParsingLayout is layout conversion table of es built-in date format to github.com/ngicks/flextime optional string format.
// Formats that contains weekyear is not supported by this module. It is listed though.
var ParsingLayout map[string]string = map[string]string{
	// It accepts 9 digits of fration-of-time anyway.
	// The document says that non-strict format can accept strings like yyyy, yyy, yy or y, while Go has no equivalent.
	// yyyy or yy is best effort here. I do not personally agree with this _too elastic_ formats.
	DateOptionalTime:                   "[yy]yy-M-d['T'HH:m:s.999999999Z]",
	StrictDateOptionalTime:             "yyyy-MM-dd['T'HH:mm:ss.999999999Z]",
	StrictDateOptionalTimeNanos:        "yyyy-MM-dd['T'HH:mm:ss.999999999Z]",
	BasicDate:                          "yyyyMMdd",
	BasicDateTime:                      "yyyyMMdd'T'HHmmss.999999999Z",
	BasicDateTimeNoMillis:              "yyyyMMdd'T'HHmmssZ",
	BasicOrdinalDate:                   "yyyyDDD",
	BasicOrdinalDateTime:               "yyyyDDD'T'HHmmss.999999999",
	BasicOrdinalDateTimeNoMillis:       "yyyyDDD'T'HHmmssZ",
	BasicTime:                          "HHmmss.999999999Z",
	BasicTimeNoMillis:                  "HHmmssZ",
	BasicTTime:                         "'T'HHmmss.999999999Z",
	BasicTTimeNoMillis:                 "'T'HHmmssZ",
	BasicWeekDate:                      "xxxx'W'wwe",
	StrictBasicWeekDate:                "xxxx'W'wwe",
	BasicWeekDateTime:                  "xxxx'W'wwe'T'HHmmss.999999999Z",
	StrictBasicWeekDateTime:            "xxxx'W'wwe'T'HHmmss.999999999Z",
	BasicWeekDateTimeNoMillis:          "xxxx'W'wwe'T'HHmmssZ",
	StrictBasicWeekDateTimeNoMillis:    "xxxx'W'wwe'T'HHmmssZ",
	Date:                               "[yy]yy-M-d",
	StrictDate:                         "yyyy-MM-dd",
	DateHour:                           "[yy]yy-M-d'T'HH",
	StrictDateHour:                     "yyyy-MM-dd'T'HH",
	DateHourMinute:                     "[yy]yy-M-d'T'HH:mm",
	StrictDateHourMinute:               "yyyy-MM-dd'T'HH:mm",
	DateHourMinuteSecond:               "[yy]yy-M-d'T'HH:m:s",
	StrictDateHourMinuteSecond:         "yyyy-MM-dd'T'HH:mm:ss",
	DateHourMinuteSecondFraction:       "[yy]yy-M-d'T'HH:m:s.999999999",
	StrictDateHourMinuteSecondFraction: "yyyy-MM-dd'T'HH:mm:ss.999999999",
	DateHourMinuteSecondMillis:         "[yy]yy-M-d'T'HH:m:s.999999999",
	StrictDateHourMinuteSecondMillis:   "yyyy-MM-dd'T'HH:mm:ss.999999999",
	DateTime:                           "[yy]yy-M-d'T'HH:m:s.999999999Z",
	StrictDateTime:                     "yyyy-MM-dd'T'HH:mm:ss.999999999Z",
	DateTimeNoMillis:                   "[yy]yy-M-d'T'HH:m:sZ",
	StrictDateTimeNoMillis:             "yyyy-MM-dd'T'HH:mm:ssZ",
	Hour:                               "HH",
	StrictHour:                         "HH",
	HourMinute:                         "HH:m",
	StrictHourMinute:                   "HH:mm",
	HourMinuteSecond:                   "HH:m:s",
	StrictHourMinuteSecond:             "HH:mm:ss",
	HourMinuteSecondFraction:           "HH:m:s.999999999",
	StrictHourMinuteSecondFraction:     "HH:mm:ss.999999999",
	HourMinuteSecondMillis:             "HH:m:s.999999999",
	StrictHourMinuteSecondMillis:       "HH:mm:ss.999999999",
	OrdinalDate:                        "[yy]yy-DDD",
	StrictOrdinalDate:                  "yyyy-DDD",
	OrdinalDateTime:                    "[yy]yy-DDD'T'HH:m:s.999999999Z",
	StrictOrdinalDateTime:              "yyyy-DDD'T'HH:mm:ss.999999999Z",
	OrdinalDateTimeNoMillis:            "[yy]yy-DDD'T'HH:m:sZ",
	StrictOrdinalDateTimeNoMillis:      "yyyy-DDD'T'HH:mm:ssZ",
	Time:                               "HH:m:s.999999999Z",
	StrictTime:                         "HH:mm:ss.999999999Z",
	TimeNoMillis:                       "HH:m:sZ",
	StrictTimeNoMillis:                 "HH:mm:ssZ",
	TTime:                              "'T'HH:m:s.999999999Z",
	StrictTTime:                        "'T'HH:mm:ss.999999999Z",
	TTimeNoMillis:                      "'T'HH:m:sZ",
	StrictTTimeNoMillis:                "'T'HH:mm:ssZ",
	WeekDate:                           "xxxx-'W'ww-e",
	StrictWeekDate:                     "xxxx-'W'ww-e",
	WeekDateTime:                       "xxxx-'W'ww-e'T'HH:mm:ss.SSSZ",
	StrictWeekDateTime:                 "xxxx-'W'ww-e'T'HH:mm:ss.SSSZ",
	WeekDateTimeNoMillis:               "xxxx-'W'ww-e'T'HH:mm:ssZ",
	StrictWeekDateTimeNoMillis:         "xxxx-'W'ww-e'T'HH:mm:ssZ",
	Weekyear:                           "xxxx",
	StrictWeekyear:                     "xxxx",
	WeekyearWeek:                       "xxxx-'W'ww",
	StrictWeekyearWeek:                 "xxxx-'W'ww",
	WeekyearWeekDay:                    "xxxx-'W'ww-e",
	StrictWeekyearWeekDay:              "xxxx-'W'ww-e",
	Year:                               "[yy]yy",
	StrictYear:                         "yyyy",
	YearMonth:                          "[yy]yy-M",
	StrictYearMonth:                    "yyyy-MM",
	YearMonthDay:                       "[yy]yy-M-d",
	StrictYearMonthDay:                 "yyyy-MM-dd",
}

// FormatLayout contains same keys of ParsingLayout.
// The value is Go format layout string.
// If corresponding parse layout is optional string, longest format will be used as format layout.
var FormatLayout map[string]string = make(map[string]string)
var Formatters map[string]*flextime.Flextime = make(map[string]*flextime.Flextime)

func init() {
	for k, v := range ParsingLayout {
		layouts, err := flextime.NewLayoutSet(v)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: layout = %s\n", v)
			panic(err)
		}
		Formatters[k] = flextime.NewFlextime(layouts)
		FormatLayout[k] = layouts.Layout()[0]
	}
}
