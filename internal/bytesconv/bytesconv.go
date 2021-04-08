package bytesconv

import (
	"unsafe"
)

// Str2Bytes Convert different types to byte slice using types and functions in
// unsafe and reflect package(see reflect.SliceHeader and reflect.StringHeader).
// It has higher performance, but notice that it may be not safe when garbage
// collection happens.Use it when you need to temporary convert a long string
// to a byte slice and won't keep it for long time.
// 理论上与 []byte("string") 速度几乎是一致的
func Str2Bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// Bytes2Str Zero-copy convert from byte slice to a string
// see reflect.SliceHeader and reflect.StringHeader
// 理论上是 string(byte{"a","b"}) 的20倍速率
func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
