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
	"math/rand"
	"time"

	"github.com/thinkgos/go-core-package/internal/bytesconv"
)

const (
	letterString     = "QWERTYUIOPLKJHGFDSAZXCVBNMabcdefghijklmnopqrstuvwxyz"
	letterStrIdxBits = 6                       // 6 bits to represent a letter index
	letterStrIdxMask = 1<<letterStrIdxBits - 1 // All 1-bits, as many as strDigitalStrIdxBits
	letterStrIdxMax  = 63 / letterStrIdxBits   // # of letter indices fitting in 63 bits

	strDigitalString     = "QWERTYUIOPLKJHGFDSAZXCVBNMabcdefghijklmnopqrstuvwxyz0123456789"
	strDigitalStrIdxBits = 6                           // 6 bits to represent a letter index
	strDigitalStrIdxMask = 1<<strDigitalStrIdxBits - 1 // All 1-bits, as many as strDigitalStrIdxBits
	strDigitalStrIdxMax  = 63 / strDigitalStrIdxBits   // # of letter indices fitting in 63 bits

	digitalString     = "0123456789"
	digitalStrIdxBits = 4                           // 4 bits to represent a letter index
	digitalStrIdxMask = 1<<strDigitalStrIdxBits - 1 // All 1-bits, as many as digitalStrIdxBits
	digitalStrIdxMax  = 63 / strDigitalStrIdxBits   // # of letter indices fitting in 63 bits

	symbolString     = "QWERTYUIOPLKJHGFDSAZXCVBNMabcdefghijklmnopqrstuvwxyz0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_{|}~`"
	symbolStrIdxBits = 7                       // 7 bits to represent a letter index
	symbolStrIdxMask = 1<<symbolStrIdxBits - 1 // All 1-bits, as many as strDigitalStrIdxBits
	symbolStrIdxMax  = 63 / symbolStrIdxBits   // # of letter indices fitting in 63 bits
)

var globalRand = rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63() + rand.Int63() + rand.Int63()))

// RandLetter rand alpha with give length(只包含字母)
func RandLetter(length int) string {
	b := make([]byte, length)
	// A rand.Int63() generates 63 random bits, enough for strDigitalStrIdxMax letters!
	for i, cache, remain := 0, globalRand.Int63(), letterStrIdxMax; i < length; {
		if remain == 0 {
			cache, remain = globalRand.Int63(), letterStrIdxMax
		}
		if idx := int(cache & letterStrIdxMask); idx < len(letterString) {
			b[i] = letterString[idx]
			i++
		}
		cache >>= letterStrIdxBits
		remain--
	}
	return bytesconv.Bytes2Str(b)
}

// RandString rand string with give length(包含字母与数字)
func RandString(length int) string {
	b := make([]byte, length)
	// A rand.Int63() generates 63 random bits, enough for strDigitalStrIdxMax letters!
	for i, cache, remain := 0, globalRand.Int63(), strDigitalStrIdxMax; i < length; {
		if remain == 0 {
			cache, remain = globalRand.Int63(), strDigitalStrIdxMax
		}
		if idx := int(cache & strDigitalStrIdxMask); idx < len(strDigitalString) {
			b[i] = strDigitalString[idx]
			i++
		}
		cache >>= strDigitalStrIdxBits
		remain--
	}
	return bytesconv.Bytes2Str(b)
}

// RandInt64 rand int64 with give length
func RandInt64(length int) int64 {
	var val int64

	// A rand.Int63() generates 63 random bits, enough for digitalStrIdxMax letters!
	for i, cache, remain := 0, globalRand.Int63(), digitalStrIdxMax; i < length; {
		if remain == 0 {
			cache, remain = globalRand.Int63(), digitalStrIdxMax
		}
		if idx := int(cache & digitalStrIdxMask); idx < len(digitalString) && !(i == 0 && digitalString[idx] == '0') {
			val = val*10 + int64(digitalString[idx]-'0')
			i++
		}
		cache >>= digitalStrIdxBits
		remain--
	}
	return val
}

// RandSymbol rand symbol with give length(包含字母数字和特殊符号)
func RandSymbol(length int) string {
	b := make([]byte, length)
	// A rand.Int63() generates 63 random bits, enough for symbolStrIdxMax letters!
	for i, cache, remain := 0, globalRand.Int63(), symbolStrIdxMax; i < length; {
		if remain == 0 {
			cache, remain = globalRand.Int63(), symbolStrIdxMax
		}
		if idx := int(cache & symbolStrIdxMask); idx < len(symbolString) {
			b[i] = symbolString[idx]
			i++
		}
		cache >>= symbolStrIdxBits
		remain--
	}
	return bytesconv.Bytes2Str(b)
}

// Int 随机min,max中的值
func Int(min, max int) int {
	if min > max {
		panic("invalid argument to Int")
	}
	if min == max {
		return min
	}
	return globalRand.Intn(max-min) + min
}

// Int31 随机min,max中的值
func Int31(min, max int32) int32 {
	if min > max {
		panic("invalid argument to Int31")
	}
	if min == max {
		return min
	}
	return globalRand.Int31n(max-min) + min
}

// Int63 随机min,max中的值
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
