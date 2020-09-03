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
	"testing"
)

func TestSearchUints(t *testing.T) {
	type args struct {
		a []uint
		x uint
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search",
			args: args{
				a: []uint{4, 1, 7, 2, 9},
				x: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UintSlice(tt.args.a).Search(tt.args.x); got != tt.want {
				t.Errorf("SearchUints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchInt8s(t *testing.T) {
	type args struct {
		a []int8
		x int8
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search",
			args: args{
				a: []int8{4, 1, 7, 2, 9},
				x: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8Slice(tt.args.a).Search(tt.args.x); got != tt.want {
				t.Errorf("SearchInt8s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchUint8s(t *testing.T) {
	type args struct {
		a []uint8
		x uint8
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search",
			args: args{
				a: []uint8{4, 1, 7, 2, 9},
				x: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint8Slice(tt.args.a).Search(tt.args.x); got != tt.want {
				t.Errorf("SearchUint8s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchInt16s(t *testing.T) {
	type args struct {
		a []int16
		x int16
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search",
			args: args{
				a: []int16{4, 1, 7, 2, 9},
				x: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16Slice(tt.args.a).Search(tt.args.x); got != tt.want {
				t.Errorf("SearchInt16s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchUint16s(t *testing.T) {
	type args struct {
		a []uint16
		x uint16
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search",
			args: args{
				a: []uint16{4, 1, 7, 2, 9},
				x: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint16Slice(tt.args.a).Search(tt.args.x); got != tt.want {
				t.Errorf("SearchUint16s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchInt32s(t *testing.T) {
	type args struct {
		a []int32
		x int32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search",
			args: args{
				a: []int32{4, 1, 7, 2, 9},
				x: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32Slice(tt.args.a).Search(tt.args.x); got != tt.want {
				t.Errorf("SearchInt32s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchUint32s(t *testing.T) {
	type args struct {
		a []uint32
		x uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search",
			args: args{
				a: []uint32{4, 1, 7, 2, 9},
				x: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint32Slice(tt.args.a).Search(tt.args.x); got != tt.want {
				t.Errorf("SearchUint32s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchInt64s(t *testing.T) {
	type args struct {
		a []int64
		x int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search",
			args: args{
				a: []int64{4, 1, 7, 2, 9},
				x: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64Slice(tt.args.a).Search(tt.args.x); got != tt.want {
				t.Errorf("SearchInt64s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchUint64s(t *testing.T) {
	type args struct {
		a []uint64
		x uint64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search",
			args: args{
				a: []uint64{4, 1, 7, 2, 9},
				x: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64Slice(tt.args.a).Search(tt.args.x); got != tt.want {
				t.Errorf("SearchUint64s() = %v, want %v", got, tt.want)
			}
		})
	}
}
