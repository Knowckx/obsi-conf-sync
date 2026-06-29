package itime

import (
	"fmt"
)

// ClockTime 只有时分秒，等价于 Python 的 datetime.time 能力。
// 零值 var ti itime.ClockTime 表示 00:00:00
type ClockTime struct {
	Hour   int
	Minute int
	Second int
}

func NewClockTime(hour, minute, second int) ClockTime {
	return ClockTime{
		Hour:   hour,
		Minute: minute,
		Second: second,
	}
}

// AtDate 把时分秒应用到指定日期 返回"2006-01-02 15:04:05"
func (t ClockTime) AtDate(date StrDate) string {
	out := fmt.Sprintf("%s %02d:%02d:%02d", date, t.Hour, t.Minute, t.Second)
	return out
}


