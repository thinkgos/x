package extrand

import (
	cryptoRand "crypto/rand"
	"math/bits"
	"math/rand"

	"github.com/thinkgos/x/internal/bytesconv"
)

var Letter = []byte("QWERTYUIOPLKJHGFDSAZXCVBNMabcdefghijklmnopqrstuvwxyz")
var Digital = []byte("0123456789")
var DigitalLetter = []byte("QWERTYUIOPLKJHGFDSAZXCVBNMabcdefghijklmnopqrstuvwxyz0123456789")
var Symbol = []byte("QWERTYUIOPLKJHGFDSAZXCVBNMabcdefghijklmnopqrstuvwxyz0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_{|}~`")

// RandLetter rand alpha with give length(只包含字母)
func RandLetter(length int) string { return bytesconv.Bytes2Str(RandLetterBytes(length)) }

// RandLetterBytes rand alpha with give length(只包含字母)
func RandLetterBytes(length int) []byte { return randBytes(length, Letter) }

// RandNumeric rand string with give length(包含数字)
func RandNumeric(length int) string { return bytesconv.Bytes2Str(RandNumericBytes(length)) }

// RandNumericBytes rand string with give length(包含数字)
func RandNumericBytes(length int) []byte { return randBytes(length, Digital) }

// RandString rand string with give length(包含字母, 数字)
func RandString(length int) string { return bytesconv.Bytes2Str(RandStringBytes(length)) }

// RandStringBytes rand string with give length(包含字母, 数字)
func RandStringBytes(length int) []byte { return randBytes(length, DigitalLetter) }

// RandSymbol rand symbol with give length(包含字母, 数字, 特殊符号)
func RandSymbol(length int) string { return bytesconv.Bytes2Str(RandSymbolBytes(length)) }

// RandSymbolBytes rand symbol with give length(包含字母, 数字, 特殊符号)
func RandSymbolBytes(length int) []byte { return randBytes(length, Symbol) }

// Rand rand bytes(如果没有给出alphabets,将使用DigitalLetter)
func Rand(length int, alphabets ...byte) string {
	return bytesconv.Bytes2Str(RandBytes(length, alphabets...))
}

// RandBytes rand bytes(如果没有给出alphabets,将使用DigitalLetter)
func RandBytes(length int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = DigitalLetter
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

	// A rand.Int63() generates 63 random bits, enough for alphabets letters!
	for i, cache, remain := 0, rand.Int63(), max; i < length; {
		if remain == 0 {
			cache, remain = rand.Int63(), max
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
