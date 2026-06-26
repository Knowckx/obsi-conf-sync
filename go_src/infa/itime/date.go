package itime

import (
	"time"

	"github.com/cockroachdb/errors"
)

// 字符串日期 格式是 2006-01-02
type StrDate = string

// 返回今天的字符串 2026-01-01
func TodayStr() StrDate {
	today := time.Now().Format("2006-01-02")
	return today
}

// ShiftDate 接收 "YYYY-MM-DD" 格式的日期字符串和偏移天数 offset（可正可负），
// 返回计算后对应的日期字符串。
func ShiftDate(dateStr string, offset int) (string, error) {
	// 定义标准的日期格式模板
	const layout = "2006-01-02"

	// 1. 将输入的字符串解析为 time.Time 类型
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return "", errors.Newf("解析日期失败: %w", err)
	}

	// 2. 使用 AddDate 进行天数偏移
	// AddDate 的三个参数分别是：年、月、日
	shiftedTime := t.AddDate(0, 0, offset)

	// 3. 将计算后的时间格式化为字符串返回
	return shiftedTime.Format(layout), nil
}
