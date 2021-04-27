package extrand

import "testing"

func TestImproveCoverage(t *testing.T) {
	t.Log(Alphabet(16))
	t.Log(Digit(16))
	t.Log(AlphaDigit(16))
	t.Log(Symbol(16))
	t.Log(String(16))
	t.Log(String(16, DefaultAlphabet...))
	t.Log(String(16, DefaultAlphaDigit...))
	t.Log(String(16, DefaultDigit...))
	t.Log(String(16, DefaultSymbol...))
}

func BenchmarkAlphabet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Alphabet(20)
	}
}

func BenchmarkDigit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Digit(6)
	}
}

func BenchmarkAlphaDigit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AlphaDigit(20)
	}
}

func BenchmarkSymbol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Symbol(20)
	}
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(20)
	}
}
