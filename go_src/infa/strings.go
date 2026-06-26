package infa

import (
	"fmt"
	"strings"
)

type StrJoiner struct {
	parts []string
}

func NewStrJoiner() *StrJoiner {
	return &StrJoiner{
		parts: make([]string, 0, 5),
	}
}

func (t *StrJoiner) Add(s string) {
	t.parts = append(t.parts, s)
}

func (t *StrJoiner) Addf(format string, args ...any) {
	t.parts = append(t.parts, fmt.Sprintf(format, args...))
}

// 默认换行
func (t *StrJoiner) String() string {
	return t.StringJoin("\n")
}

func (t *StrJoiner) StringJoin(sep string) string {
	return strings.Join(t.parts, sep)
}