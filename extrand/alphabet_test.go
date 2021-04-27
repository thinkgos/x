package extrand

import (
	cryptoRand "crypto/rand"
	"math/bits"
	"math/rand"
	"testing"
	"time"
)

func TestImproveCoverage(t *testing.T) {
	t.Log(Alphabet(16))
	t.Log(Number(16))
	t.Log(AlphaNumber(16))
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

func BenchmarkNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Number(6)
	}
}

func BenchmarkAlphaNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AlphaNumber(20)
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

func cryptoRandBytes(length int, alphabets []byte) []byte {
	b := make([]byte, length)
	if _, err := cryptoRand.Read(b); err != nil {
		panic(err)
	}
	for i, v := range b {
		b[i] = alphabets[v%byte(len(alphabets))]
	}
	return b
}

func comRandBytes(length int, alphabets []byte) []byte {
	b := make([]byte, length)
	bn := bits.Len(uint(len(alphabets)))
	mask := int64(1)<<bn - 1
	max := 63 / bn
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// A rand.Int63() generates 63 random bits, enough for alphabets letters!
	for i, cache, remain := 0, r.Int63(), max; i < length; {
		if remain == 0 {
			cache, remain = r.Int63(), max
		}
		if idx := int(cache & mask); idx < len(alphabets) {
			b[i] = alphabets[idx]
			i++
		}
		cache >>= bn
		remain--
	}
	return b
}

func Benchmark_cryptoRandBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cryptoRandBytes(10, DefaultSymbol)
	}
}

func Benchmark_comRandBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		comRandBytes(10, DefaultSymbol)
	}
}
