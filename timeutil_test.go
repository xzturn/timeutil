package timeutil

import "testing"

func TestParseTime(t *testing.T) {
    ts := "20151018120135"
    t.Log("ORIGIN: " + ts)
    t1, _ := ParseTimeDayHourMinSec(ts)
    t.Log(t1)
    t.Log(FormatAsDayHourMinSec(t1))
    t2, _ := ParseTimeDayHourMin(ts[0:12])
    t.Log(t2)
    t.Log(FormatAsDayHourMin(t2))
    t3, _ := ParseTimeDayHour(ts[0:10])
    t.Log(t3)
    t.Log(FormatAsDayHour(t3))
    t4, _ := ParseTimeDay(ts[0:8])
    t.Log(t4)
    t.Log(FormatAsDay(t4))
}
