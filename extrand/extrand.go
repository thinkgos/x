// Copyright [2020] [thinkgos] thinkgo@aliyun.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package extrand extend rand
package extrand

import (
	"math/bits"
	"math/rand"
	"time"

	"github.com/thinkgos/x/internal/bytesconv"
)

const (
	LetterString     = "QWERTYUIOPLKJHGFDSAZXCVBNMabcdefghijklmnopqrstuvwxyz"
	letterStrIdxBits = 6                       // 6 bits to represent a letter index
	letterStrIdxMask = 1<<letterStrIdxBits - 1 // All 1-bits, as many as digitalLetterStrIdxBits
	letterStrIdxMax  = 63 / letterStrIdxBits   // # of letter indices fitting in 63 bits

	DigitalLetterString     = "QWERTYUIOPLKJHGFDSAZXCVBNMabcdefghijklmnopqrstuvwxyz0123456789"
	digitalLetterStrIdxBits = 6                              // 6 bits to represent a letter index
	digitalLetterStrIdxMask = 1<<digitalLetterStrIdxBits - 1 // All 1-bits, as many as digitalLetterStrIdxBits
	digitalLetterStrIdxMax  = 63 / digitalLetterStrIdxBits   // # of letter indices fitting in 63 bits

	DigitalString     = "0123456789"
	digitalStrIdxBits = 4                              // 4 bits to represent a letter index
	digitalStrIdxMask = 1<<digitalLetterStrIdxBits - 1 // All 1-bits, as many as digitalStrIdxBits
	digitalStrIdxMax  = 63 / digitalLetterStrIdxBits   // # of letter indices fitting in 63 bits

	SymbolString     = "QWERTYUIOPLKJHGFDSAZXCVBNMabcdefghijklmnopqrstuvwxyz0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_{|}~`"
	symbolStrIdxBits = 7                       // 7 bits to represent a letter index
	symbolStrIdxMask = 1<<symbolStrIdxBits - 1 // All 1-bits, as many as digitalLetterStrIdxBits
	symbolStrIdxMax  = 63 / symbolStrIdxBits   // # of letter indices fitting in 63 bits
)

var Letter = []byte(LetterString)
var DigitalLetter = []byte(DigitalLetterString)
var Digital = []byte(DigitalString)
var Symbol = []byte(SymbolString)

var globalRand = rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63() + rand.Int63() + rand.Int63()))

// RandLetterBytes rand alpha with give length(只包含字母)
func RandLetterBytes(length int) []byte {
	b := make([]byte, length)
	// A rand.Int63() generates 63 random bits, enough for digitalLetterStrIdxMax letters!
	for i, cache, remain := 0, globalRand.Int63(), letterStrIdxMax; i < length; {
		if remain == 0 {
			cache, remain = globalRand.Int63(), letterStrIdxMax
		}
		if idx := int(cache & letterStrIdxMask); idx < len(LetterString) {
			b[i] = LetterString[idx]
			i++
		}
		cache >>= letterStrIdxBits
		remain--
	}
	return b
}

// RandLetter rand alpha with give length(只包含字母)
func RandLetter(length int) string { return bytesconv.Bytes2Str(RandLetterBytes(length)) }

// RandNumericBytes rand string with give length(包含数字)
func RandNumericBytes(length int) []byte {
	b := make([]byte, length)
	// A rand.Int63() generates 63 random bits, enough for digitalStrIdxMax letters!
	for i, cache, remain := 0, globalRand.Int63(), digitalStrIdxMax; i < length; {
		if remain == 0 {
			cache, remain = globalRand.Int63(), digitalStrIdxMax
		}
		if idx := int(cache & digitalStrIdxMask); idx < len(DigitalString) && !(i == 0 && DigitalString[idx] == '0') {
			b[i] = DigitalString[idx]
			i++
		}
		cache >>= digitalStrIdxBits
		remain--
	}
	return b
}

// RandNumeric rand string with give length(包含数字)
func RandNumeric(length int) string { return bytesconv.Bytes2Str(RandNumericBytes(length)) }

// RandStringBytes rand string with give length(包含字母与数字)
func RandStringBytes(length int) []byte {
	b := make([]byte, length)
	// A rand.Int63() generates 63 random bits, enough for digitalLetterStrIdxMax letters!
	for i, cache, remain := 0, globalRand.Int63(), digitalLetterStrIdxMax; i < length; {
		if remain == 0 {
			cache, remain = globalRand.Int63(), digitalLetterStrIdxMax
		}
		if idx := int(cache & digitalLetterStrIdxMask); idx < len(DigitalLetterString) {
			b[i] = DigitalLetterString[idx]
			i++
		}
		cache >>= digitalLetterStrIdxBits
		remain--
	}
	return b
}

// RandString rand string with give length(包含字母与数字)
func RandString(length int) string { return bytesconv.Bytes2Str(RandStringBytes(length)) }

// RandSymbolBytes rand symbol with give length(包含字母数字和特殊符号)
func RandSymbolBytes(length int) []byte {
	b := make([]byte, length)
	// A rand.Int63() generates 63 random bits, enough for symbolStrIdxMax letters!
	for i, cache, remain := 0, globalRand.Int63(), symbolStrIdxMax; i < length; {
		if remain == 0 {
			cache, remain = globalRand.Int63(), symbolStrIdxMax
		}
		if idx := int(cache & symbolStrIdxMask); idx < len(SymbolString) {
			b[i] = SymbolString[idx]
			i++
		}
		cache >>= symbolStrIdxBits
		remain--
	}
	return b
}

// RandSymbol rand symbol with give length(包含字母数字和特殊符号)
func RandSymbol(length int) string { return bytesconv.Bytes2Str(RandSymbolBytes(length)) }

// RandBytes rand bytes(如果没有给出alphabets,将使用DigitalLetter)
func RandBytes(length int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = DigitalLetter
	}

	bn := bits.Len(uint(len(alphabets)))
	mask := int64(1)<<bn - 1
	max := 63 / bn

	b := make([]byte, length)
	// A rand.Int63() generates 63 random bits, enough for alphabets letters!
	for i, cache, remain := 0, globalRand.Int63(), max; i < length; {
		if remain == 0 {
			cache, remain = globalRand.Int63(), max
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

// Rand rand bytes(如果没有给出alphabets,将使用DigitalLetter)
func Rand(length int, alphabets ...byte) string {
	return bytesconv.Bytes2Str(RandBytes(length, alphabets...))
}

// RandInt64 rand int64 with give length
func RandInt64(length int) int64 {
	var val int64

	// A rand.Int63() generates 63 random bits, enough for digitalStrIdxMax letters!
	for i, cache, remain := 0, globalRand.Int63(), digitalStrIdxMax; i < length; {
		if remain == 0 {
			cache, remain = globalRand.Int63(), digitalStrIdxMax
		}
		if idx := int(cache & digitalStrIdxMask); idx < len(DigitalString) && !(i == 0 && DigitalString[idx] == '0') {
			val = val*10 + int64(DigitalString[idx]-'0')
			i++
		}
		cache >>= digitalStrIdxBits
		remain--
	}
	return val
}

// Int 随机[min,max)中的值
func Int(min, max int) int {
	if min > max {
		panic("invalid argument to Int")
	}
	if min == max {
		return min
	}
	return globalRand.Intn(max-min) + min
}

// Int31 随机[min,max)中的值
func Int31(min, max int32) int32 {
	if min > max {
		panic("invalid argument to Int31")
	}
	if min == max {
		return min
	}
	return globalRand.Int31n(max-min) + min
}

// Int63 随机[min,max)中的值
func Int63(min, max int64) int64 {
	if min > max {
		panic("invalid argument to Int63")
	}
	if min == max {
		return min
	}
	return globalRand.Int63n(max-min) + min
}

// Float64 随机min,max中的值
func Float64(min, max float64) float64 {
	if min > max {
		panic("invalid argument to Float64")
	}
	return min + (max-min)*globalRand.Float64()
}

// Shuffle pseudo-randomizes the order of elements using the default Source.
func Shuffle(str string) string {
	runes := []rune(str)
	globalRand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})
	return string(runes)
}
