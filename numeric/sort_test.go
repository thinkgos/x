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
	"sort"
	"testing"
)

func TestSortUints(t *testing.T) {
	type args struct {
		a []uint
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sort",
			args: args{a: []uint{4, 2, 7, 9, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UintSlice(tt.args.a).Sort()

			if !sort.IsSorted(UintSlice(tt.args.a)) {
				t.Error("SortUints test failed")
			}
		})
	}
}

func TestSortInt8s(t *testing.T) {
	type args struct {
		a []int8
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sort",
			args: args{a: []int8{4, 2, 7, 9, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Int8Slice(tt.args.a).Sort()

			if !sort.IsSorted(Int8Slice(tt.args.a)) {
				t.Error("SortInt8s test failed")
			}
		})
	}
}

func TestSortUint8s(t *testing.T) {
	type args struct {
		a []uint8
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sort",
			args: args{a: []uint8{4, 2, 7, 9, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Uint8Slice(tt.args.a).Sort()

			if !sort.IsSorted(Uint8Slice(tt.args.a)) {
				t.Error("SortUint8s test failed")
			}
		})
	}
}

func TestSortInt16s(t *testing.T) {
	type args struct {
		a []int16
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sort",
			args: args{a: []int16{4, 2, 7, 9, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Int16Slice(tt.args.a).Sort()

			if !sort.IsSorted(Int16Slice(tt.args.a)) {
				t.Error("SortInt16s test failed")
			}
		})
	}
}

func TestSortUint16s(t *testing.T) {
	type args struct {
		a []uint16
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sort",
			args: args{a: []uint16{4, 2, 7, 9, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Uint16Slice(tt.args.a).Sort()

			if !sort.IsSorted(Uint16Slice(tt.args.a)) {
				t.Error("SortUint16s test failed")
			}
		})
	}
}

func TestSortInt32s(t *testing.T) {
	type args struct {
		a []int32
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sort",
			args: args{a: []int32{4, 2, 7, 9, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Int32Slice(tt.args.a).Sort()

			if !sort.IsSorted(Int32Slice(tt.args.a)) {
				t.Error("SortInt32s test failed")
			}
		})
	}
}

func TestSortUint32s(t *testing.T) {
	type args struct {
		a []uint32
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sort",
			args: args{a: []uint32{4, 2, 7, 9, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Uint32Slice(tt.args.a).Sort()

			if !sort.IsSorted(Uint32Slice(tt.args.a)) {
				t.Error("SortUint32s test failed")
			}
		})
	}
}

func TestSortInt64s(t *testing.T) {
	type args struct {
		a []int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sort",
			args: args{a: []int64{4, 2, 7, 9, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Int64Slice(tt.args.a).Sort()

			if !sort.IsSorted(Int64Slice(tt.args.a)) {
				t.Error("SortInt64s test failed")
			}
		})
	}
}

func TestSortUint64s(t *testing.T) {
	type args struct {
		a []uint64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sort",
			args: args{a: []uint64{4, 2, 7, 9, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Uint64Slice(tt.args.a).Sort()

			if !sort.IsSorted(Uint64Slice(tt.args.a)) {
				t.Error("SortUint64s test failed")
			}
		})
	}
}
