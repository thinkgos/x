package numeric

import (
	"math"
)

// MaxByte returns the max value
func MaxByte(x, y byte) byte {
	if x < y {
		return y
	}
	return x
}

// MaxInt returns the max value
func MaxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// MaxUint returns the max value
func MaxUint(x, y uint) uint {
	if x < y {
		return y
	}
	return x
}

// MaxInt8 returns the max value
func MaxInt8(x, y int8) int8 {
	if x < y {
		return y
	}
	return x
}

// MaxUint8 returns the max value
func MaxUint8(x, y uint8) uint8 {
	if x < y {
		return y
	}
	return x
}

// MaxInt16 returns the max value
func MaxInt16(x, y int16) int16 {
	if x < y {
		return y
	}
	return x
}

// MaxUint16 returns the max value
func MaxUint16(x, y uint16) uint16 {
	if x < y {
		return y
	}
	return x
}

// MaxInt32 returns the max value
func MaxInt32(x, y int32) int32 {
	if x < y {
		return y
	}
	return x
}

// MaxUint32 returns the max value
func MaxUint32(x, y uint32) uint32 {
	if x < y {
		return y
	}
	return x
}

// MaxInt64 returns the max value
func MaxInt64(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

// MaxUint64 returns the max value
func MaxUint64(x, y uint64) uint64 {
	if x < y {
		return y
	}
	return x
}

// MaxFloat64 returns the max value
func MaxFloat64(x, y float64) float64 {
	return math.Max(x, y)
}
