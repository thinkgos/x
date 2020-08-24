package extrand

import (
	"testing"
)

func TestImproveCoverage(t *testing.T) {
	t.Log(RandInt64(16))
	t.Log(RandString(16))
}

func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandString(10)
	}
}

func BenchmarkRandInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandInt64(10)
	}
}
