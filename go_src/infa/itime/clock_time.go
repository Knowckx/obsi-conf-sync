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




// // NewClockTime 创建 00:00:00 的时间。
// func NewClockTime3() ClockTime {
// 	return ClockTime{}
// }

// // NewClockTimeInt 创建时分秒时间。
// func (t ClockTime) NewWithHMS(hour, minute, second int) {
// 	t.Hour,t.Minute,t.Second = hour, minute, second
// }

// // NowClockTime 返回当前本地时分秒。
// func (t ClockTime) NowClockTime() ClockTime {
// 	return t.ClockTimeFromTime(time.Now())
// }

// // ClockTimeFromTime 从 time.Time 提取时分秒。
// func (t ClockTime) ClockTimeFromTime(t time.Time) ClockTime {
// 	return t.NewClockTimeInt(t.Hour(), t.Minute(), t.Second())
// }

// // ParseClockTime 解析 HH:MM:SS 格式的时分秒。
// func (t ClockTime) ParseClockTime(value string) (ClockTime, error) {
// 	t, err := time.Parse(time.TimeOnly, value)
// 	if err != nil {
// 		return ClockTime{}, errors.WithStack(err)
// 	}
// 	return t.ClockTimeFromTime(t), nil
// }

// // String 返回 HH:MM:SS 格式。
// func (t ClockTime) String() string {
// 	return fmt.Sprintf("%02d:%02d:%02d", t.Hour, t.Minute, t.Second)
// }

// // SelfCheck 检查时分秒是否在正常范围内。
// func (t ClockTime) SelfCheck() error {
// 	if t.Hour < 0 || t.Hour > 23 {
// 		return errors.Errorf("hour out of range: %d", t.Hour)
// 	}
// 	if t.Minute < 0 || t.Minute > 59 {
// 		return errors.Errorf("minute out of range: %d", t.Minute)
// 	}
// 	if t.Second < 0 || t.Second > 59 {
// 		return errors.Errorf("second out of range: %d", t.Second)
// 	}
// 	return nil
// }

// // Seconds 返回当天已过秒数。
// func (t ClockTime) Seconds() int {
// 	return t.Hour*3600 + t.Minute*60 + t.Second
// }

// // Duration 返回当天零点到当前时分秒的时间段。
// func (t ClockTime) Duration() time.Duration {
// 	return time.Duration(t.Seconds()) * time.Second
// }

// // Compare 比较两个时分秒。
// func (t ClockTime) Compare(other ClockTime) int {
// 	return t.Seconds() - other.Seconds()
// }

// // Before 判断当前时分秒是否早于 other。
// func (t ClockTime) Before(other ClockTime) bool {
// 	return t.Compare(other) < 0
// }

// // After 判断当前时分秒是否晚于 other。
// func (t ClockTime) After(other ClockTime) bool {
// 	return t.Compare(other) > 0
// }

// // Equal 判断两个时分秒是否相同。
// func (t ClockTime) Equal(other ClockTime) bool {
// 	return t.Compare(other) == 0
// }
