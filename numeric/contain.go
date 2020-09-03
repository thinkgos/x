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

import (
	"reflect"
)

// ContainInt checks if x exists in []ints and returns TRUE if x is found.
func ContainInt(vs []int, x int) bool {
	for _, v := range vs {
		if x == v {
			return true
		}
	}
	return false
}

// ContainUint checks if x exists in []uints and returns TRUE if x is found.
func ContainUint(vs []uint, x uint) bool {
	for _, v := range vs {
		if x == v {
			return true
		}
	}
	return false
}

// ContainInt8 checks if x exists in []int8s and returns TRUE if x is found.
func ContainInt8(vs []int8, x int8) bool {
	for _, v := range vs {
		if x == v {
			return true
		}
	}
	return false
}

// ContainUint8 checks if x exists in []uint8s and returns TRUE if x is found.
func ContainUint8(vs []uint8, x uint8) bool {
	for _, v := range vs {
		if x == v {
			return true
		}
	}
	return false
}

// ContainInt16 checks if x exists in []int16s and returns TRUE if x is found.
func ContainInt16(vs []int16, x int16) bool {
	for _, v := range vs {
		if x == v {
			return true
		}
	}
	return false
}

// ContainUint16 checks if x exists in []uint16s and returns TRUE if x is found.
func ContainUint16(vs []uint16, x uint16) bool {
	for _, v := range vs {
		if x == v {
			return true
		}
	}
	return false
}

// ContainInt32 checks if x exists in []int32s and returns TRUE if x is found.
func ContainInt32(vs []int32, x int32) bool {
	for _, v := range vs {
		if x == v {
			return true
		}
	}
	return false
}

// ContainUint32 checks if x exists in []uint32s and returns TRUE if x is found.
func ContainUint32(vs []uint32, x uint32) bool {
	for _, v := range vs {
		if x == v {
			return true
		}
	}
	return false
}

// ContainInt64 checks if x exists in []int64s and returns TRUE if x is found.
func ContainInt64(vs []int64, x int64) bool {
	for _, v := range vs {
		if x == v {
			return true
		}
	}
	return false
}

// ContainUint64 checks if x exists in []uint64s and returns TRUE if x is found.
func ContainUint64(vs []uint64, x uint64) bool {
	for _, v := range vs {
		if x == v {
			return true
		}
	}
	return false
}

// ContainFloat64 checks if x exists in []float64s and returns TRUE if x is found.
func ContainFloat64(vs []float64, x float64) bool {
	for _, v := range vs {
		if x == v {
			return true
		}
	}
	return false
}

// Contain checks if x exists in a slice and returns TRUE if x is found.
func Contain(vs []interface{}, x interface{}) bool {
	for _, v := range vs {
		if reflect.DeepEqual(x, v) {
			return true
		}
	}
	return false
}
