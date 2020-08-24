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

// Package extbase64 extend base64
package extbase64

import (
	"encoding/base64"

	"github.com/thinkgos/go-core-package/internal/bytesconv"
)

// Encode base64 encode bytes
func Encode(b []byte) []byte {
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(buf, b)
	return buf
}

// EncodeString base64 encode string
func EncodeString(str string) string {
	return base64.StdEncoding.EncodeToString(bytesconv.Str2Bytes(str))
}

// Decode base64 decode to bytes
func Decode(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}

// DecodeString base64 decode to string
func DecodeString(str string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(str)
	return bytesconv.Bytes2Str(b), err
}
