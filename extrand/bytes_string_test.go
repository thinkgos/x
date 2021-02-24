package extrand

import "testing"

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
}

func BenchmarkRandLetter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandLetter(20)
	}
}

func BenchmarkRandNumeric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandNumeric(20)
	}
}

func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandString(20)
	}
}

func BenchmarkRandSymbol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandSymbol(20)
	}
}

func BenchmarkRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rand(20)
	}
}
