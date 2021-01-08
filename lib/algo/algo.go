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

// Package algo implement common api
package algo

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"

	"github.com/thinkgos/x/internal/bytesconv"
)

// MD5 calculate the md5 hash of a hex string.
func MD5(s string) string {
	h := md5.New()
	h.Write(bytesconv.Str2Bytes(s)) // nolint: errCheck
	return hex.EncodeToString(h.Sum(nil))
}

// SHA1 calculate the sha1 hash of a hex string.
func SHA1(s string) string {
	h := sha1.New()
	h.Write(bytesconv.Str2Bytes(s)) // nolint: errCheck
	return hex.EncodeToString(h.Sum(nil))
}

// SHA256 calculate the sha256 hash of a hex string.
func SHA256(s string) string {
	h := sha256.New()
	h.Write(bytesconv.Str2Bytes(s)) // nolint: errCheck
	return hex.EncodeToString(h.Sum(nil))
}

// SHA512 calculate the sha512 hash of a hex string.
func SHA512(s string) string {
	h := sha512.New()
	h.Write(bytesconv.Str2Bytes(s)) // nolint: errCheck
	return hex.EncodeToString(h.Sum(nil))
}

// Hash Generate a hex hash value,
// expects: md5, sha1, sha224, sha256, sha384, sha512.
func Hash(method, val string) string {
	var h hash.Hash

	switch method {
	case "md5":
		h = md5.New()
	case "sha1":
		h = sha1.New()
	case "sha224":
		h = sha256.New224()
	case "sha256":
		h = sha256.New()
	case "sha384":
		h = sha512.New384()
	case "sha512":
		h = sha512.New()
	default:
		return val
	}
	h.Write(bytesconv.Str2Bytes(val)) // nolint: errCheck
	return hex.EncodeToString(h.Sum(nil))
}

// Hmac Generate a hex hash value with the key,
// expects: hmacmd5, hmacsha1, hmacsha224, hmacsha256, hmacsha384, hmacsha512.
func Hmac(method, key, val string) string {
	var f func() hash.Hash

	switch method {
	case "hmacmd5":
		f = md5.New
	case "hmacsha1":
		f = sha1.New
	case "hmacsha224":
		f = sha256.New224
	case "hmacsha256":
		f = sha256.New
	case "hmacsha384":
		f = sha512.New384
	case "hmacsha512":
		f = sha512.New
	default:
		return val
	}
	h := hmac.New(f, bytesconv.Str2Bytes(key))
	h.Write(bytesconv.Str2Bytes(val)) // nolint: errCheck
	return hex.EncodeToString(h.Sum(nil))
}
