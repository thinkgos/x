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
	letterString     = "QWERTYUIOPLKJHGFDSAZXCVBNMabcdefghijklmnopqrstuvwxyz0123456789"
	letterStrIdxBits = 6                       // 6 bits to represent a letter index
	letterStrIdxMask = 1<<letterStrIdxBits - 1 // All 1-bits, as many as letterStrIdxBits
	letterStrIdxMax  = 63 / letterStrIdxBits   // # of letter indices fitting in 63 bits

	letterInt        = "0123456789"
	letterIntIdxBits = 4                       // 4 bits to represent a letter index
	letterIntIdxMask = 1<<letterStrIdxBits - 1 // All 1-bits, as many as letterIntIdxBits
	letterIntIdxMax  = 63 / letterStrIdxBits   // # of letter indices fitting in 63 bits

)

var globalRand = rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63() + rand.Int63() + rand.Int63()))

// RandString rand string  with give length
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
