package sharp

import (
	"unsafe"
)

// IsMachineLittleEndian 判断系统大小端
func IsMachineLittleEndian() bool {
	var i int16 = 0x0001

	u := unsafe.Pointer(&i)
	b := *((*byte)(u))

	return b == 0x01
}
