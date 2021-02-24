package extrand

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntx(t *testing.T) {
	require.Panics(t, func() {
		Intx(5, 1)
	})
	got := Intx(2, 20)
	require.LessOrEqual(t, 2, got)
	require.LessOrEqual(t, got, 20)
	require.Equal(t, 10, Intx(10, 10))
}

func TestInt31x(t *testing.T) {
	require.Panics(t, func() {
		Int31x(5, 1)
	})
	got := Int31x(2, 20)
	require.LessOrEqual(t, int32(2), got)
	require.LessOrEqual(t, got, int32(20))
	require.Equal(t, int32(10), Int31x(10, 10))
}

func TestInt63x(t *testing.T) {
	require.Panics(t, func() {
		Int63x(5, 1)
	})
	got := Int63x(2, 20)
	require.LessOrEqual(t, int64(2), got)
	require.LessOrEqual(t, got, int64(20))
	require.Equal(t, int64(10), Int63x(10, 10))
}

func TestFloat64x(t *testing.T) {
	require.Panics(t, func() {
		Float64x(5.1, 1.1)
	})
	got := Float64x(2.1, 20.1)
	require.LessOrEqual(t, 2.1, got)
	require.LessOrEqual(t, got, 20.1)
	require.Equal(t, 10.1, Float64x(10.1, 10.1))
}
