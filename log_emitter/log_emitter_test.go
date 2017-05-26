package le

import (
	"testing"
	"time"
)

func TestLogEmitter_loadConf(t *testing.T) {
	var l LogEmitter
	l.LoadConf("conf_test.yml")

	result1 := l.LOGS[0].INTERVAL
	expect1 := 3
	if result1 != expect1 {
		t.Error("interval is not consistent.", "result:", result1, "expect", expect1)
	}

	result2 := l.LOGS[1].CATEGORY
	expect2 := "nginx-error"
	if result2 != expect2 {
		t.Error("category is not consistent.", "result:", result2, "expect", expect2)
	}
}

func TestLogEmitter_updateTimestamp(t *testing.T) {
	var l LogEmitter

	raw := "timestamp: 2016-12-29T18:21:42+09:00"
	ts_re := "[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}\\+[0-9]{2}:[0-9]{2}"
	ts_format := "2006-01-02T15:04:05Z07:00"
	t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+10:00")

	result := l.updateTimestamp(raw, ts_re, ts_format, t1)
	expect := "timestamp: 2012-11-01T22:08:41+10:00"

	if result != expect {
		t.Error("updateTimestamp does not work as expect.", "result:", result, "expect:", expect)
	}
}
