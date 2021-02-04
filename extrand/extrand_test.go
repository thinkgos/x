package extrand

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestImproveCoverage(t *testing.T) {
	t.Log(RandLetter(16))
	t.Log(RandNumeric(16))
	t.Log(RandString(16))
	t.Log(RandSymbol(16))
	t.Log(Rand(16))
	t.Log(Rand(16, Letter...))
	t.Log(Rand(16, DigitalLetter...))
	t.Log(Rand(16, Digital...))
	t.Log(Rand(16, Symbol...))
	t.Log(RandInt64(16))
}

func BenchmarkRandLetter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandLetter(10)
	}
}

func BenchmarkRandNumeric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandNumeric(10)
	}
}

func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandString(10)
	}
}

func BenchmarkRandSymbol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandSymbol(10)
	}
}

func BenchmarkRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rand(10)
	}
}

func BenchmarkRandInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandInt64(10)
	}
}

func TestInt(t *testing.T) {
	require.Panics(t, func() {
		Int(5, 1)
	})
	got := Int(2, 20)
	require.LessOrEqual(t, 2, got)
	require.LessOrEqual(t, got, 20)
	require.Equal(t, 10, Int(10, 10))
}

func TestInt31(t *testing.T) {
	require.Panics(t, func() {
		Int31(5, 1)
	})
	got := Int31(2, 20)
	require.LessOrEqual(t, int32(2), got)
	require.LessOrEqual(t, got, int32(20))
	require.Equal(t, int32(10), Int31(10, 10))
}

func TestInt63(t *testing.T) {
	require.Panics(t, func() {
		Int63(5, 1)
	})
	got := Int63(2, 20)
	require.LessOrEqual(t, int64(2), got)
	require.LessOrEqual(t, got, int64(20))
	require.Equal(t, int64(10), Int63(10, 10))
}

func TestFloat64(t *testing.T) {
	require.Panics(t, func() {
		Float64(5.1, 1.1)
	})
	got := Float64(2.1, 20.1)
	require.LessOrEqual(t, 2.1, got)
	require.LessOrEqual(t, got, 20.1)
	require.Equal(t, 10.1, Float64(10.1, 10.1))
}

func TestShuffle(t *testing.T) {
	t.Log(Shuffle("hello world"))
	t.Log(Shuffle("hello world"))
}
