package numeric

import (
	"testing"
)

func TestMax(t *testing.T) {
	var x = 2
	var y = 1

	if got := MaxByte(byte(x), byte(y)); got != byte(x) {
		t.Errorf("MaxByte got %d, but want %d", got, byte(x))
	}
	if got := MaxByte(byte(y), byte(x)); got != byte(x) {
		t.Errorf("MaxByte got %d, but want %d", got, byte(x))
	}

	if got := MaxInt(x, y); got != x {
		t.Errorf("MaxInt got %d, but want %d", got, x)
	}
	if got := MaxInt(y, x); got != x {
		t.Errorf("MaxInt got %d, but want %d", got, byte(x))
	}

	if got := MaxUint(uint(x), uint(y)); got != uint(x) {
		t.Errorf("MaxUint got %d, but want %d", got, uint(x))
	}
	if got := MaxUint(uint(y), uint(x)); got != uint(x) {
		t.Errorf("MaxUint got %d, but want %d", got, uint(x))
	}

	if got := MaxInt8(int8(x), int8(y)); got != int8(x) {
		t.Errorf("MaxInt8 got %d, but want %d", got, int8(x))
	}
	if got := MaxInt8(int8(y), int8(x)); got != int8(x) {
		t.Errorf("MaxInt8 got %d, but want %d", got, int8(x))
	}

	if got := MaxUint8(uint8(x), uint8(y)); got != uint8(x) {
		t.Errorf("MaxUint8 got %d, but want %d", got, uint8(x))
	}
	if got := MaxUint8(uint8(y), uint8(x)); got != uint8(x) {
		t.Errorf("MaxUint8 got %d, but want %d", got, uint8(x))
	}

	if got := MaxInt16(int16(x), int16(y)); got != int16(x) {
		t.Errorf("MaxInt16 got %d, but want %d", got, int16(x))
	}
	if got := MaxInt16(int16(y), int16(x)); got != int16(x) {
		t.Errorf("MaxInt16 got %d, but want %d", got, int16(x))
	}

	if got := MaxUint16(uint16(x), uint16(y)); got != uint16(x) {
		t.Errorf("MaxUint16 got %d, but want %d", got, uint16(x))
	}
	if got := MaxUint16(uint16(y), uint16(x)); got != uint16(x) {
		t.Errorf("MaxUint16 got %d, but want %d", got, uint16(x))
	}

	if got := MaxInt32(int32(x), int32(y)); got != int32(x) {
		t.Errorf("MaxInt32 got %d, but want %d", got, int32(x))
	}
	if got := MaxInt32(int32(y), int32(x)); got != int32(x) {
		t.Errorf("MaxInt32 got %d, but want %d", got, int32(x))
	}

	if got := MaxUint32(uint32(x), uint32(y)); got != uint32(x) {
		t.Errorf("MaxUint32 got %d, but want %d", got, uint32(x))
	}
	if got := MaxUint32(uint32(y), uint32(x)); got != uint32(x) {
		t.Errorf("MaxUint32 got %d, but want %d", got, uint32(x))
	}

	if got := MaxInt64(int64(x), int64(y)); got != int64(x) {
		t.Errorf("MaxInt64 got %d, but want %d", got, int64(x))
	}
	if got := MaxInt64(int64(y), int64(x)); got != int64(x) {
		t.Errorf("MaxInt64 got %d, but want %d", got, int64(x))
	}

	if got := MaxUint64(uint64(x), uint64(y)); got != uint64(x) {
		t.Errorf("MaxUint64 got %d, but want %d", got, uint64(x))
	}
	if got := MaxUint64(uint64(y), uint64(x)); got != uint64(x) {
		t.Errorf("MaxUint64 got %d, but want %d", got, uint64(x))
	}

	if got := MaxFloat64(float64(x), float64(y)); got != float64(x) {
		t.Errorf("MaxUint64 got %f, but want %f", got, float64(x))
	}
	if got := MaxFloat64(float64(y), float64(x)); got != float64(x) {
		t.Errorf("MaxUint64 got %f, but want %f", got, float64(x))
	}
}
