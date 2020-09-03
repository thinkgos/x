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

// AppendInt append int to slice with no duplicates.
func AppendInt(s []int, e int) []int {
	for _, v := range s {
		if v == e {
			return s
		}
	}
	return append(s, e)
}

// AppendUint append uint to slice with no duplicates.
func AppendUint(s []uint, e uint) []uint {
	for _, v := range s {
		if v == e {
			return s
		}
	}
	return append(s, e)
}

// AppendInt8 append int8 to slice with no duplicates.
func AppendInt8(s []int8, e int8) []int8 {
	for _, v := range s {
		if v == e {
			return s
		}
	}
	return append(s, e)
}

// AppendUint8 append uint8 to slice with no duplicates.
func AppendUint8(s []uint8, e uint8) []uint8 {
	for _, v := range s {
		if v == e {
			return s
		}
	}
	return append(s, e)
}

// AppendInt16 appends int16 to slice with no duplicates.
func AppendInt16(s []int16, e int16) []int16 {
	for _, v := range s {
		if v == e {
			return s
		}
	}
	return append(s, e)
}

// AppendUint16 appends uint16 to slice with no duplicates.
func AppendUint16(s []uint16, e uint16) []uint16 {
	for _, v := range s {
		if v == e {
			return s
		}
	}
	return append(s, e)
}

// AppendInt32 appends int32 to slice with no duplicates.
func AppendInt32(s []int32, e int32) []int32 {
	for _, v := range s {
		if v == e {
			return s
		}
	}
	return append(s, e)
}

// AppendUint32 appends uint16 to slice with no duplicates.
func AppendUint32(s []uint32, e uint32) []uint32 {
	for _, v := range s {
		if v == e {
			return s
		}
	}
	return append(s, e)
}

// AppendInt64 appends int64 to slice with no duplicates.
func AppendInt64(s []int64, e int64) []int64 {
	for _, v := range s {
		if v == e {
			return s
		}
	}
	return append(s, e)
}

// AppendUint64 appends int64 to slice with no duplicates.
func AppendUint64(s []uint64, e uint64) []uint64 {
	for _, v := range s {
		if v == e {
			return s
		}
	}
	return append(s, e)
}

// DeleteUint delete an uint element from slice if it exist
func DeleteUint(s []uint, e uint) []uint {
	for i, v := range s {
		if v == e {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// DeleteInt delete an int element from slice if it exist
func DeleteInt(s []int, e int) []int {
	for i, v := range s {
		if v == e {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// DeleteUint8 delete an uint8 element from slice if it exist
func DeleteUint8(s []uint8, e uint8) []uint8 {
	for i, v := range s {
		if v == e {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// DeleteInt8 delete an int8 element from slice if it exist
func DeleteInt8(s []int8, e int8) []int8 {
	for i, v := range s {
		if v == e {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// DeleteUint16 delete an uint16 element from slice if it exist
func DeleteUint16(s []uint16, e uint16) []uint16 {
	for i, v := range s {
		if v == e {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// DeleteInt16 delete an int16 element from slice if it exist
func DeleteInt16(s []int16, e int16) []int16 {
	for i, v := range s {
		if v == e {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// DeleteUint32 delete an uint32 element from slice if it exist
func DeleteUint32(s []uint32, e uint32) []uint32 {
	for i, v := range s {
		if v == e {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// DeleteInt32 delete an int32 element from slice if it exist
func DeleteInt32(s []int32, e int32) []int32 {
	for i, v := range s {
		if v == e {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// DeleteUint64 delete an uint64 element from slice if it exist
func DeleteUint64(s []uint64, e uint64) []uint64 {
	for i, v := range s {
		if v == e {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// DeleteInt64 delete an int64 element from slice if it exist
func DeleteInt64(s []int64, e int64) []int64 {
	for i, v := range s {
		if v == e {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// DeleteIntAll delete all int element from slice if it exist
func DeleteIntAll(s []int, e int) []int {
	if len(s) == 0 {
		return s
	}
	tmpS := make([]int, 0, len(s))
	for _, v := range s {
		if v != e {
			tmpS = append(tmpS, v)
		}
	}
	return tmpS
}

// DeleteUintAll delete all uint element from slice if it exist
func DeleteUintAll(s []uint, e uint) []uint {
	if len(s) == 0 {
		return s
	}
	tmpS := make([]uint, 0, len(s))
	for _, v := range s {
		if v != e {
			tmpS = append(tmpS, v)
		}
	}
	return tmpS
}

// DeleteInt8All delete all int8 element from slice if it exist
func DeleteInt8All(s []int8, e int8) []int8 {
	if len(s) == 0 {
		return s
	}
	tmpS := make([]int8, 0, len(s))
	for _, v := range s {
		if v != e {
			tmpS = append(tmpS, v)
		}
	}
	return tmpS
}

// DeleteUint8All delete all uint8 element from slice if it exist
func DeleteUint8All(s []uint8, e uint8) []uint8 {
	if len(s) == 0 {
		return s
	}
	tmpS := make([]uint8, 0, len(s))
	for _, v := range s {
		if v != e {
			tmpS = append(tmpS, v)
		}
	}
	return tmpS
}

// DeleteInt16All delete all int16 element from slice if it exist
func DeleteInt16All(s []int16, e int16) []int16 {
	if len(s) == 0 {
		return s
	}
	tmpS := make([]int16, 0, len(s))
	for _, v := range s {
		if v != e {
			tmpS = append(tmpS, v)
		}
	}
	return tmpS
}

// DeleteUint16All delete all uint16 element from slice if it exist
func DeleteUint16All(s []uint16, e uint16) []uint16 {
	if len(s) == 0 {
		return s
	}
	tmpS := make([]uint16, 0, len(s))
	for _, v := range s {
		if v != e {
			tmpS = append(tmpS, v)
		}
	}
	return tmpS
}

// DeleteInt32All delete all int32 element from slice if it exist
func DeleteInt32All(s []int32, e int32) []int32 {
	if len(s) == 0 {
		return s
	}
	tmpS := make([]int32, 0, len(s))
	for _, v := range s {
		if v != e {
			tmpS = append(tmpS, v)
		}
	}
	return tmpS
}

// DeleteUint32All delete all uint32 element from slice if it exist
func DeleteUint32All(s []uint32, e uint32) []uint32 {
	if len(s) == 0 {
		return s
	}
	tmpS := make([]uint32, 0, len(s))
	for _, v := range s {
		if v != e {
			tmpS = append(tmpS, v)
		}
	}
	return tmpS
}

// DeleteInt64All delete all int64 element from slice if it exist
func DeleteInt64All(s []int64, e int64) []int64 {
	if len(s) == 0 {
		return s
	}
	tmpS := make([]int64, 0, len(s))
	for _, v := range s {
		if v != e {
			tmpS = append(tmpS, v)
		}
	}
	return tmpS
}

// DeleteUint64All delete all uint64 element from slice if it exist
func DeleteUint64All(s []uint64, e uint64) []uint64 {
	if len(s) == 0 {
		return s
	}
	tmpS := make([]uint64, 0, len(s))
	for _, v := range s {
		if v != e {
			tmpS = append(tmpS, v)
		}
	}
	return tmpS
}
