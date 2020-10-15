package extmath

import (
	"math"
)

// Pow is int64 type of math.Pow function.
func Pow(x, y int64) int64 {
	if y <= 0 {
		return 1
	}
	if y%2 == 0 {
		sqrt := Pow(x, y/2)
		return sqrt * sqrt
	}
	return Pow(x, y-1) * x
}

// Round 数值保留小数点,n为保留小数点位数
func Round(f float64, n int) float64 {
	pow10N := math.Pow10(n)
	return math.Trunc((f+0.5/pow10N)*pow10N) / pow10N
}
