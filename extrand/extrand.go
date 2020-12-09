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
	alphaString     = "QWERTYUIOPLKJHGFDSAZXCVBNMabcdefghijklmnopqrstuvwxyz"
	alphaStrIdxBits = 6                      // 6 bits to represent a letter index
	alphaStrIdxMask = 1<<alphaStrIdxBits - 1 // All 1-bits, as many as letterStrIdxBits
	alphaStrIdxMax  = 63 / alphaStrIdxBits   // # of letter indices fitting in 63 bits

	letterString     = "QWERTYUIOPLKJHGFDSAZXCVBNMabcdefghijklmnopqrstuvwxyz0123456789"
	letterStrIdxBits = 6                       // 6 bits to represent a letter index
	letterStrIdxMask = 1<<letterStrIdxBits - 1 // All 1-bits, as many as letterStrIdxBits
	letterStrIdxMax  = 63 / letterStrIdxBits   // # of letter indices fitting in 63 bits

	letterInt        = "0123456789"
	letterIntIdxBits = 4                       // 4 bits to represent a letter index
	letterIntIdxMask = 1<<letterStrIdxBits - 1 // All 1-bits, as many as letterIntIdxBits
	letterIntIdxMax  = 63 / letterStrIdxBits   // # of letter indices fitting in 63 bits

	letterSymbol        = "QWERTYUIOPLKJHGFDSAZXCVBNMabcdefghijklmnopqrstuvwxyz0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_{|}~`"
	letterSymbolIdxBits = 7                          // 7 bits to represent a letter index
	letterSymbolIdxMask = 1<<letterSymbolIdxBits - 1 // All 1-bits, as many as letterStrIdxBits
	letterSymbolIdxMax  = 63 / letterSymbolIdxBits   // # of letter indices fitting in 63 bits
)

var globalRand = rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63() + rand.Int63() + rand.Int63()))

// RandAlpha rand alpha with give length(只包含字母)
func RandAlpha(length int) string {
	b := make([]byte, length)
	// A rand.Int63() generates 63 random bits, enough for letterStrIdxMax letters!
	for i, cache, remain := 0, globalRand.Int63(), alphaStrIdxMax; i < length; {
		if remain == 0 {
			cache, remain = globalRand.Int63(), alphaStrIdxMax
		}
		if idx := int(cache & alphaStrIdxMask); idx < len(alphaString) {
			b[i] = alphaString[idx]
			i++
		}
		cache >>= alphaStrIdxBits
		remain--
	}
	return bytesconv.Bytes2Str(b)
}

// RandString rand string with give length(包含字母与数字)
func RandString(length int) string {
	b := make([]byte, length)
	// A rand.Int63() generates 63 random bits, enough for letterStrIdxMax letters!
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

// RandInt64 rand int64 with give length
func RandInt64(length int) int64 {
	var val int64

	// A rand.Int63() generates 63 random bits, enough for letterIntIdxMax letters!
	for i, cache, remain := 0, globalRand.Int63(), letterIntIdxMax; i < length; {
		if remain == 0 {
			cache, remain = globalRand.Int63(), letterIntIdxMax
		}
		if idx := int(cache & letterIntIdxMask); idx < len(letterInt) && !(i == 0 && letterInt[idx] == '0') {
			val = val*10 + int64(letterInt[idx]-'0')
			i++
		}
		cache >>= letterIntIdxBits
		remain--
	}
	return val
}

// RandSymbol rand symbol with give length(包含字母数字和特殊符号)
func RandSymbol(length int) string {
	b := make([]byte, length)
	// A rand.Int63() generates 63 random bits, enough for letterSymbolIdxMax letters!
	for i, cache, remain := 0, globalRand.Int63(), letterSymbolIdxMax; i < length; {
		if remain == 0 {
			cache, remain = globalRand.Int63(), letterSymbolIdxMax
		}
		if idx := int(cache & letterSymbolIdxMask); idx < len(letterSymbol) {
			b[i] = letterSymbol[idx]
			i++
		}
		cache >>= letterSymbolIdxBits
		remain--
	}
	return bytesconv.Bytes2Str(b)
}

func Int(min, max int) int {
	if min > max {
		panic("invalid argument to Int")
	}
	if min == max {
		return min
	}
	return globalRand.Intn(max-min) + min
}

func Int31(min, max int32) int32 {
	if min > max {
		panic("invalid argument to Int31")
	}
	if min == max {
		return min
	}
	return globalRand.Int31n(max-min) + min
}

func Int63(min, max int64) int64 {
	if min > max {
		panic("invalid argument to Int63")
	}
	if min == max {
		return min
	}
	return globalRand.Int63n(max-min) + min
}

func Float64(min, max float64) float64 {
	if min > max {
		panic("invalid argument to Float64")
	}
	return min + (max-min)*globalRand.Float64()
}
