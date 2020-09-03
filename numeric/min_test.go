package numeric

import (
	"testing"
)

func TestMin(t *testing.T) {
	var x = 1
	var y = 2

	if got := MinByte(byte(x), byte(y)); got != byte(x) {
		t.Errorf("MinByte got %d, but want %d", got, byte(x))
	}
	if got := MinByte(byte(y), byte(x)); got != byte(x) {
		t.Errorf("MinByte got %d, but want %d", got, byte(x))
	}

	if got := MinInt(x, y); got != x {
		t.Errorf("MinInt got %d, but want %d", got, x)
	}
	if got := MinInt(y, x); got != x {
		t.Errorf("MinInt got %d, but want %d", got, byte(x))
	}

	if got := MinUint(uint(x), uint(y)); got != uint(x) {
		t.Errorf("MinUint got %d, but want %d", got, uint(x))
	}
	if got := MinUint(uint(y), uint(x)); got != uint(x) {
		t.Errorf("MinUint got %d, but want %d", got, uint(x))
	}

	if got := MinInt8(int8(x), int8(y)); got != int8(x) {
		t.Errorf("MinInt8 got %d, but want %d", got, int8(x))
	}
	if got := MinInt8(int8(y), int8(x)); got != int8(x) {
		t.Errorf("MinInt8 got %d, but want %d", got, int8(x))
	}

	if got := MinUint8(uint8(x), uint8(y)); got != uint8(x) {
		t.Errorf("MinUint8 got %d, but want %d", got, uint8(x))
	}
	if got := MinUint8(uint8(y), uint8(x)); got != uint8(x) {
		t.Errorf("MinUint8 got %d, but want %d", got, uint8(x))
	}

	if got := MinInt16(int16(x), int16(y)); got != int16(x) {
		t.Errorf("MinInt16 got %d, but want %d", got, int16(x))
	}
	if got := MinInt16(int16(y), int16(x)); got != int16(x) {
		t.Errorf("MinInt16 got %d, but want %d", got, int16(x))
	}

	if got := MinUint16(uint16(x), uint16(y)); got != uint16(x) {
		t.Errorf("MinUint16 got %d, but want %d", got, uint16(x))
	}
	if got := MinUint16(uint16(y), uint16(x)); got != uint16(x) {
		t.Errorf("MinUint16 got %d, but want %d", got, uint16(x))
	}

	if got := MinInt32(int32(x), int32(y)); got != int32(x) {
		t.Errorf("MinInt32 got %d, but want %d", got, int32(x))
	}
	if got := MinInt32(int32(y), int32(x)); got != int32(x) {
		t.Errorf("MinInt32 got %d, but want %d", got, int32(x))
	}

	if got := MinUint32(uint32(x), uint32(y)); got != uint32(x) {
		t.Errorf("MinUint32 got %d, but want %d", got, uint32(x))
	}
	if got := MinUint32(uint32(y), uint32(x)); got != uint32(x) {
		t.Errorf("MinUint32 got %d, but want %d", got, uint32(x))
	}

	if got := MinInt64(int64(x), int64(y)); got != int64(x) {
		t.Errorf("MinInt64 got %d, but want %d", got, int64(x))
	}
	if got := MinInt64(int64(y), int64(x)); got != int64(x) {
		t.Errorf("MinInt64 got %d, but want %d", got, int64(x))
	}

	if got := MinUint64(uint64(x), uint64(y)); got != uint64(x) {
		t.Errorf("MinUint64 got %d, but want %d", got, uint64(x))
	}
	if got := MinUint64(uint64(y), uint64(x)); got != uint64(x) {
		t.Errorf("MinUint64 got %d, but want %d", got, uint64(x))
	}

	if got := MinFloat64(float64(y), float64(x)); got != float64(x) {
		t.Errorf("MinFloat64 got %f, but want %f", got, float64(x))
	}
	if got := MinFloat64(float64(x), float64(y)); got != float64(x) {
		t.Errorf("MinFloat64 got %f, but want %f", got, float64(x))
	}
}
