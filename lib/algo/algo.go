// Copyright 2020 thinkgos.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package algo implement common api
package algo

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

// algo method
const (
	MethodMD5    Method = "md5"
	MethodSha1   Method = "sha1"
	MethodSha224 Method = "sha224"
	MethodSha256 Method = "sha256"
	MethodSha384 Method = "sha384"
	MethodSha512 Method = "sha512"
)

// Method algo method
type Method string

// MD5 calculate the md5 hash of a hex string.
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s)) // nolint: errCheck
	return hex.EncodeToString(h.Sum(nil))
}

// SHA1 calculate the sha1 hash of a hex string.
func SHA1(s string) string {
	h := sha1.New()
	h.Write([]byte(s)) // nolint: errCheck
	return hex.EncodeToString(h.Sum(nil))
}

// Hash Generate a hex hash value, expects: MD5, SHA1, SHA224, SHA256, SHA384, SHA512.
func Hash(method Method, s string) string {
	var h hash.Hash

	switch method {
	case MethodMD5:
		h = md5.New()
	case MethodSha1:
		h = sha1.New()
	case MethodSha224:
		h = sha256.New224()
	case MethodSha256:
		h = sha256.New()
	case MethodSha384:
		h = sha512.New384()
	case MethodSha512:
		h = sha512.New()
	default:
		return s
	}
	h.Write([]byte(s)) // nolint: errCheck
	return hex.EncodeToString(h.Sum(nil))
}

// Hmac Generate a hex hash value with the key, expects: MD5, SHA1, SHA224, SHA256, SHA384, SHA512.
func Hmac(method Method, s, key string) string {
	var mac hash.Hash

	switch method {
	case MethodMD5:
		mac = hmac.New(md5.New, []byte(key))
	case MethodSha1:
		mac = hmac.New(sha1.New, []byte(key))
	case MethodSha224:
		mac = hmac.New(sha256.New224, []byte(key))
	case MethodSha256:
		mac = hmac.New(sha256.New, []byte(key))
	case MethodSha384:
		mac = hmac.New(sha512.New384, []byte(key))
	case MethodSha512:
		mac = hmac.New(sha512.New, []byte(key))
	default:
		return s
	}
	mac.Write([]byte(s)) // nolint: errCheck
	return hex.EncodeToString(mac.Sum(nil))
}

// AddSlashes returns a string with backslashes added before characters that need to be escaped.
func AddSlashes(s string) string {
	var buf bytes.Buffer

	for _, ch := range s {
		if ch == '\'' || ch == '"' || ch == '\\' {
			buf.WriteRune('\\')
		}
		buf.WriteRune(ch)
	}
	return buf.String()
}

// StripSlashes returns a string with backslashes stripped off. (\' becomes ' and so on.) Double backslashes (\\) are made into a single backslash (\).
func StripSlashes(s string) string {
	var buf bytes.Buffer

	l, skip := len(s), false
	for i, ch := range s {
		if skip {
			buf.WriteRune(ch)
			skip = false
			continue
		}

		if ch == '\\' {
			if i+1 < l && s[i+1] == '\\' {
				skip = true
			}
			continue
		}

		buf.WriteRune(ch)
	}
	return buf.String()
}

// QuoteMeta returns a version of str with a backslash character (\) before every character that is among these: . \ + * ? [ ^ ] ( $ )
func QuoteMeta(s string) string {
	var buf bytes.Buffer

	for _, ch := range s {
		switch ch {
		case '.', '+', '\\', '(', '$', ')', '[', '^', ']', '*', '?':
			buf.WriteRune('\\')
		}

		buf.WriteRune(ch)
	}

	return buf.String()
}
