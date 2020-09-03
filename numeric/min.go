package numeric

import (
	"math"
)

// MinByte returns the min value
func MinByte(x, y byte) byte {
	if x < y {
		return x
	}
	return y
}

// MinInt returns the min value
func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// MinUint returns the min value
func MinUint(x, y uint) uint {
	if x < y {
		return x
	}
	return y
}

// MinInt8 returns the min value
func MinInt8(x, y int8) int8 {
	if x < y {
		return x
	}
	return y
}

// MinUint8 returns the min value
func MinUint8(x, y uint8) uint8 {
	if x < y {
		return x
	}
	return y
}

// MinInt16 returns the min value
func MinInt16(x, y int16) int16 {
	if x < y {
		return x
	}
	return y
}

// MinUint16 returns the min value
func MinUint16(x, y uint16) uint16 {
	if x < y {
		return x
	}
	return y
}

// MinInt32 returns the min value
func MinInt32(x, y int32) int32 {
	if x < y {
		return x
	}
	return y
}

// MinUint32 returns the min value
func MinUint32(x, y uint32) uint32 {
	if x < y {
		return x
	}
	return y
}

// MinInt64 returns the min value
func MinInt64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

// MinUint64 returns the min value
func MinUint64(x, y uint64) uint64 {
	if x < y {
		return x
	}
	return y
}

// MinFloat64 returns the min value
func MinFloat64(x, y float64) float64 {
	return math.Min(x, y)
}
