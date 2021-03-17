package extrand

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt63(t *testing.T) {
	got := Int63()
	require.LessOrEqual(t, got, int64(math.MaxInt64))
	require.LessOrEqual(t, int64(0), got)
}

func TestInt63n(t *testing.T) {
	require.Panics(t, func() {
		Int63n(-1)
	})
	got := Int63n(math.MaxInt64 / 2)
	require.LessOrEqual(t, got, int64(math.MaxInt64/2))
	require.LessOrEqual(t, int64(0), got)
}

func TestInt31(t *testing.T) {
	got := Int31()
	require.LessOrEqual(t, got, int32(math.MaxInt32))
	require.LessOrEqual(t, int32(0), got)
}

func TestInt31n(t *testing.T) {
	require.Panics(t, func() {
		Int31n(-1)
	})
	got := Int31n(math.MaxInt32 / 2)
	require.LessOrEqual(t, got, int32(math.MaxInt32/2))
	require.LessOrEqual(t, int32(0), got)
}

func TestInt(t *testing.T) {
	got := Int()

	require.LessOrEqual(t, got, math.MaxInt64)
	require.LessOrEqual(t, 0, got)
}

func TestIntn(t *testing.T) {
	require.Panics(t, func() {
		Intn(-1)
	})
	got := Intn(math.MaxInt64 / 2)
	require.LessOrEqual(t, got, math.MaxInt64/2)
	require.LessOrEqual(t, 0, got)
}

func TestFloat64(t *testing.T) {
	got := Float64()
	require.Less(t, got, 1.0)
	require.LessOrEqual(t, 0.0, got)
}

func TestFloat32(t *testing.T) {
	got := Float32()
	require.Less(t, got, float32(1.0))
	require.LessOrEqual(t, float32(0.0), got)
}

func TestPerm(t *testing.T) {
	got := Perm(100)
	for _, v := range got {
		require.Less(t, v, 100)
	}
}

func BenchmarkUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint64()
	}
}

func BenchmarkUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint32()
	}
}
