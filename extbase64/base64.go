package outil

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
