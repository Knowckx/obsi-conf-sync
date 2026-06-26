package itime

import (
	"errors"
	"fmt"
)

// DateRange 表示一个日期范围
type DateRange struct {
	Start StrDate
	End   StrDate
}

func NewDateRange (st, end StrDate) DateRange {
	return DateRange{
		Start: st,
		End: end,
	}
}

func (t DateRange) String() string {
	out := fmt.Sprintf("(%s -> %s)", t.Start, t.End)
	return out
}

func (t DateRange) SelfCheck() error {
	if t.Start == "" {
		return  errors.New("start date is required")
	}
	if t.End == "" {
		return  errors.New("end date is required")
	}
	if t.Start > t.End {
		return  errors.New("start date must be less than or equal to end date")
	}
	return nil
}
