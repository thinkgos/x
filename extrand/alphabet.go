package extrand

import (
	cryptoRand "crypto/rand"
	"math/bits"
	"math/rand"
	"time"

	"github.com/thinkgos/x/internal/bytesconv"
)

// previous defined bytes
var (
	DefaultAlphabet   = []byte("QWERTYUIOPLKJHGFDSAZXCVBNMabcdefghijklmnopqrstuvwxyz")
	DefaultDigit      = []byte("0123456789")
	DefaultAlphaDigit = []byte("QWERTYUIOPLKJHGFDSAZXCVBNMabcdefghijklmnopqrstuvwxyz0123456789")
	DefaultSymbol     = []byte("QWERTYUIOPLKJHGFDSAZXCVBNMabcdefghijklmnopqrstuvwxyz0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_{|}~`")
)

// Alphabet rand alpha with give length(只包含字母)
func Alphabet(length int) string { return bytesconv.Bytes2Str(AlphabetBytes(length)) }

// AlphabetBytes rand alpha with give length(只包含字母)
func AlphabetBytes(length int) []byte { return randBytes(length, DefaultAlphabet) }

// Number rand string with give length(只包含数字)
func Number(length int) string { return bytesconv.Bytes2Str(NumberBytes(length)) }

// NumberBytes rand string with give length(只包含数字)
func NumberBytes(length int) []byte { return randBytes(length, DefaultDigit) }

// AlphaNumber rand string with give length(只包含字母, 数字)
func AlphaNumber(length int) string { return bytesconv.Bytes2Str(AlphaNumberBytes(length)) }

// AlphaNumberBytes rand string with give length(只包含字母, 数字)
func AlphaNumberBytes(length int) []byte { return randBytes(length, DefaultAlphaDigit) }

// Symbol rand symbol with give length(只包含字母, 数字, 特殊符号)
func Symbol(length int) string { return bytesconv.Bytes2Str(SymbolBytes(length)) }

// SymbolBytes rand symbol with give length(只包含字母, 数字, 特殊符号)
func SymbolBytes(length int) []byte { return randBytes(length, DefaultSymbol) }

// String rand bytes(如果没有给出alphabets,将使用 DefaultAlphabet)
func String(length int, alphabets ...byte) string {
	return bytesconv.Bytes2Str(Bytes(length, alphabets...))
}

// Bytes rand bytes(如果没有给出alphabets,将使用 DefaultAlphabet)
func Bytes(length int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = DefaultAlphaDigit
	}
	return randBytes(length, alphabets)
}

func randBytes(length int, alphabets []byte) []byte {
	b := make([]byte, length)
	if _, err := cryptoRand.Read(b); err == nil {
		for i, v := range b {
			b[i] = alphabets[v%byte(len(alphabets))]
		}
		return b
	}

	bn := bits.Len(uint(len(alphabets)))
	mask := int64(1)<<bn - 1
	max := 63 / bn
	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63() + rand.Int63()))

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
