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

// UniqueInts takes an input slice of ints and
// returns a new slice of ints without duplicate values.
func UniqueInts(a []int) []int {
	l := len(a)
	if l <= 1 {
		return a
	}

	m := make(map[int]struct{}, l)
	r := make([]int, 0, l)
	for _, v := range a {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			r = append(r, v)
		}
	}
	return r
}

// UniqueUints takes an input slice of uints and
// returns a new slice of uints without duplicate values.
func UniqueUints(a []uint) []uint {
	l := len(a)
	if l <= 1 {
		return a
	}

	m := make(map[uint]struct{}, l)
	r := make([]uint, 0, l)
	for _, v := range a {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			r = append(r, v)
		}
	}
	return r
}

// UniqueInt8s takes an input slice of int8s and
// returns a new slice of int8s without duplicate values.
func UniqueInt8s(a []int8) []int8 {
	l := len(a)
	if l <= 1 {
		return a
	}

	m := make(map[int8]struct{}, l)
	r := make([]int8, 0, l)
	for _, v := range a {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			r = append(r, v)
		}
	}
	return r
}

// UniqueUint8s takes an input slice of uint8s and
// returns a new slice of uint8s without duplicate values.
func UniqueUint8s(a []uint8) []uint8 {
	l := len(a)
	if l <= 1 {
		return a
	}

	m := make(map[uint8]struct{}, l)
	r := make([]uint8, 0, l)
	for _, v := range a {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			r = append(r, v)
		}
	}
	return r
}

// UniqueInt16s takes an input slice of int16s and
// returns a new slice of int16s without duplicate values.
func UniqueInt16s(a []int16) []int16 {
	l := len(a)
	if l <= 1 {
		return a
	}

	m := make(map[int16]struct{}, l)
	r := make([]int16, 0, l)
	for _, v := range a {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			r = append(r, v)
		}
	}
	return r
}

// UniqueUint16s takes an input slice of uint16s and
// returns a new slice of uint16s without duplicate values.
func UniqueUint16s(a []uint16) []uint16 {
	l := len(a)
	if l <= 1 {
		return a
	}

	m := make(map[uint16]struct{}, l)
	r := make([]uint16, 0, l)
	for _, v := range a {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			r = append(r, v)
		}
	}
	return r
}

// UniqueInt32s takes an input slice of int32s and
// returns a new slice of int32s without duplicate values.
func UniqueInt32s(a []int32) []int32 {
	l := len(a)
	if l <= 1 {
		return a
	}

	m := make(map[int32]struct{}, l)
	r := make([]int32, 0, l)
	for _, v := range a {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			r = append(r, v)
		}
	}
	return r
}

// UniqueUint32s takes an input slice of uint32s and
// returns a new slice of uint32s without duplicate values.
func UniqueUint32s(a []uint32) []uint32 {
	l := len(a)
	if l <= 1 {
		return a
	}

	m := make(map[uint32]struct{}, l)
	r := make([]uint32, 0, l)
	for _, v := range a {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			r = append(r, v)
		}
	}
	return r
}

// UniqueInt64s takes an input slice of int64s and
// returns a new slice of int64s without duplicate values.
func UniqueInt64s(a []int64) []int64 {
	l := len(a)
	if l <= 1 {
		return a
	}

	m := make(map[int64]struct{}, l)
	r := make([]int64, 0, l)
	for _, v := range a {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			r = append(r, v)
		}
	}
	return r
}

// UniqueUint64s takes an input slice of uint64s and
// returns a new slice of uint64s without duplicate values.
func UniqueUint64s(a []uint64) []uint64 {
	l := len(a)
	if l <= 1 {
		return a
	}

	m := make(map[uint64]struct{}, l)
	r := make([]uint64, 0, l)
	for _, v := range a {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			r = append(r, v)
		}
	}
	return r
}

// UniqueFloat64s takes an input slice of float64s and
// returns a new slice of float64s without duplicate values.
func UniqueFloat64s(a []float64) []float64 {
	l := len(a)
	if l <= 1 {
		return a
	}

	m := make(map[float64]struct{}, l)
	r := make([]float64, 0, l)
	for _, v := range a {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			r = append(r, v)
		}
	}
	return r
}
