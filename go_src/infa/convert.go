package infa

import (
	"strconv"
	"strings"

	"github.com/cockroachdb/errors"
)

// StrToFloat 字段解析成浮点数
func StrToFloat(input string, field string) (float64, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return 0, nil
	}

	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, errors.Wrapf(err, "parse float field %s: %q", field, input)
	}

	return value, nil
}

// StrToInt 字段解析成整数
func StrToInt(input string, field string) (int, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return 0, nil
	}

	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.Wrapf(err, "parse int field %s: %q", field, input)
	}

	return value, nil
}

