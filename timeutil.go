package timeutil

import (
    "errors"
    "fmt"
    "time"
)

const defaultZoneTag = "+08:00"

func ParseTimeDayHourMinSec(t string) (time.Time, error) {
    // parse from yyyyMMddHHmmss
    if len(t) != 14 { return time.Time{}, errors.New("Expect yyyyMMddHHmmss") }
    return time.Parse(time.RFC3339, t[0:4] + "-" + t[4:6] + "-" + t[6:8] + "T" + t[8:10] + ":" + t[10:12] + ":" + t[12:14] + defaultZoneTag)
}

func ParseTimeDayHourMin(t string) (time.Time, error) {
    // parse from yyyyMMddHHmm
    if len(t) != 12 { return time.Time{}, errors.New("Expect yyyyMMddHHmm") }
    return time.Parse(time.RFC3339, t[0:4] + "-" + t[4:6] + "-" + t[6:8] + "T" + t[8:10] + ":" + t[10:12] + ":00" + defaultZoneTag)
}

func ParseTimeDayHour(t string) (time.Time, error) {
    // parse from yyyyMMddHH
    if len(t) != 10 { return time.Time{}, errors.New("Expect yyyyMMddHH") }
    return time.Parse(time.RFC3339, t[0:4] + "-" + t[4:6] + "-" + t[6:8] + "T" + t[8:10] + ":00:00" + defaultZoneTag)
}

func ParseTimeDay(t string) (time.Time, error) {
    // parse from yyyyMMdd
    if len(t) != 8 { return time.Time{}, errors.New("Expect yyyyMMdd") }
    return time.Parse(time.RFC3339, t[0:4] + "-" + t[4:6] + "-" + t[6:8] + "T00:00:00" + defaultZoneTag)
}


func FormatAsDayHourMinSec(t time.Time) string {
    // format as yyyyMMddHHmmss
    return fmt.Sprintf("%04d%02d%02d%02d%02d%02d", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func FormatAsDayHourMin(t time.Time) string {
    // format as yyyyMMddHHmm
    return fmt.Sprintf("%04d%02d%02d%02d%02d", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute())
}

func FormatAsDayHour(t time.Time) string {
    // format as yyyyMMddHH
    return fmt.Sprintf("%04d%02d%02d%02d", t.Year(), int(t.Month()), t.Day(), t.Hour())
}

func FormatAsDay(t time.Time) string {
    // format as yyyyMMdd
    return fmt.Sprintf("%04d%02d%02d", t.Year(), int(t.Month()), t.Day())
}
