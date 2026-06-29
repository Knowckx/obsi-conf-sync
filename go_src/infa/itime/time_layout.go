package itime

import (
	"go_trade/src/infa/ops"
	"time"
)

// StrLayout 表示日志时间格式。
type StrLayout string

const (
	StrLayoutTimeMillis StrLayout = "15:04:05.000"
)

// String 返回时间格式字符串。
func (t StrLayout) String() string {
	return string(t)
}

var DefulatLoction = mustLoadLocation("Asia/Shanghai")


func mustLoadLocation(name string) *time.Location {
	loc, err := time.LoadLocation(name)
	ops.MustNoErr(err)
	return loc
}
