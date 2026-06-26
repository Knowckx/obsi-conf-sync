package infa

import (
	"math"
)

// 0.001 → 3位小数  假如转换不成功 返回0
func CalcDecimalPlaces(size float64) int {
	if size <= 0 || math.IsNaN(size) || math.IsInf(size, 0) {
		return -1
	}

	places := int(math.Round(-math.Log10(size)))
	if places < 0 {
		return -1
	}
	return places
}



// 取符号
func Sign(n int) int {
	if n > 0 {
		return 1
	}
	return -1
}