// Copyright [2020] [thinkgos]
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

package numeric

// BreakUint16 break into bytes
func BreakUint16(v uint16) (loByte, hiByte byte) {
	return byte(v), byte(v >> 8)
}

// BreakUint32  break into bytes
func BreakUint32(v uint32) (byte0, byte1, byte2, byte3 byte) {
	return byte(v), byte(v >> 8), byte(v >> 16), byte(v >> 24)
}

// BreakUint64 break into uint32
func BreakUint64(v uint64) (lo32, hi32 uint32) {
	return uint32(v), uint32(v >> 32)
}

// BuildUint16 combine into uint16
func BuildUint16(loByte, hiByte byte) uint16 {
	return uint16(loByte) | uint16(hiByte)<<8
}

// BuildUint32 combine into uint32
func BuildUint32(byte0, byte1, byte2, byte3 byte) uint32 {
	return uint32(byte0) | (uint32(byte1) << 8) | (uint32(byte2) << 16) | (uint32(byte3) << 24)
}

// BuildUint64 combine into uint64
func BuildUint64(lo32, hi32 uint32) uint64 {
	return uint64(lo32) | uint64(hi32)<<32
}

// ReverseBytes reverse []byte
func ReverseBytes(b []byte) []byte {
	for from, to := 0, len(b)-1; from < to; from, to = from+1, to-1 {
		b[from], b[to] = b[to], b[from]
	}

	return b
}
