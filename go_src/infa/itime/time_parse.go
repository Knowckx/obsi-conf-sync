package itime

import (
	"time"

	"github.com/cockroachdb/errors"
)

// Parse 把时间字符串转换成 time.Time，默认东8区。
//
// 支持 2006-01-02 15:04:05 || 15.04.05
func Parse(input string) (time.Time, error) {
	out, err := time.ParseInLocation(time.DateTime, input, DefulatLoction)
	if err == nil {
		return out, nil
	}

	out, err = parseTodayAt(input)
	if err == nil {
		return out, nil
	}

	return time.Time{}, errors.Errorf("unsupported time format: %s", input)
}

// 特定的格式 "15.04.05"
func parseTodayAt(input string) (time.Time, error) {
	t, err := time.ParseInLocation("15.04.05", input, DefulatLoction)
	if err != nil {
		return time.Time{}, errors.WithStack(err)
	}

	now := time.Now().In(DefulatLoction)
	out := time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), t.Second(), 0, DefulatLoction)
	return out, nil
}

// time.Time → string
func TimeToStr(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}