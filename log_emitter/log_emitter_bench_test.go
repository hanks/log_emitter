package le

import (
	"testing"
	"time"
)

func BenchmarkLogEmitter_loadConf(b *testing.B) {
	b.StopTimer()
	var l LogEmitter
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		l.LoadConf("conf_test.yml")
	}
}

func BenchmarkLogEmitter_updateTimestamp(b *testing.B) {
	b.StopTimer()
	var l LogEmitter
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		raw := "timestamp: 2016-12-29T18:21:42+09:00"
		ts_re := "[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}\\+[0-9]{2}:[0-9]{2}"
		ts_format := "2006-01-02T15:04:05Z07:00"
		t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+10:00")

		l.updateTimestamp(raw, ts_re, ts_format, t1)
	}
}
