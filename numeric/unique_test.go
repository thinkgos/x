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
	"testing"
)

func TestUniqueInts(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "smaller than 1",
			args: args{a: []int{2}},
			want: []int{2},
		},
		{
			name: "unique",
			args: args{a: []int{2, 4, 6, 7, 1, 3, 4, 9, 7}},
			want: []int{2, 4, 6, 7, 1, 3, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueInts(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueUints(t *testing.T) {
	type args struct {
		a []uint
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{
			name: "smaller than 1",
			args: args{a: []uint{2}},
			want: []uint{2},
		},
		{
			name: "unique",
			args: args{a: []uint{2, 4, 6, 7, 1, 3, 4, 9, 7}},
			want: []uint{2, 4, 6, 7, 1, 3, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueUints(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueUints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueInt8s(t *testing.T) {
	type args struct {
		a []int8
	}
	tests := []struct {
		name string
		args args
		want []int8
	}{
		{
			name: "smaller than 1",
			args: args{a: []int8{2}},
			want: []int8{2},
		},
		{
			name: "unique",
			args: args{a: []int8{2, 4, 6, 7, 1, 3, 4, 9, 7}},
			want: []int8{2, 4, 6, 7, 1, 3, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueInt8s(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueInt8s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueUint8s(t *testing.T) {
	type args struct {
		a []uint8
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{
			name: "smaller than 1",
			args: args{a: []uint8{2}},
			want: []uint8{2},
		},
		{
			name: "unique",
			args: args{a: []uint8{2, 4, 6, 7, 1, 3, 4, 9, 7}},
			want: []uint8{2, 4, 6, 7, 1, 3, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueUint8s(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueUint8s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueInt16s(t *testing.T) {
	type args struct {
		a []int16
	}
	tests := []struct {
		name string
		args args
		want []int16
	}{
		{
			name: "smaller than 1",
			args: args{a: []int16{2}},
			want: []int16{2},
		},
		{
			name: "unique",
			args: args{a: []int16{2, 4, 6, 7, 1, 3, 4, 9, 7}},
			want: []int16{2, 4, 6, 7, 1, 3, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueInt16s(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueInt16s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueUint16s(t *testing.T) {
	type args struct {
		a []uint16
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{
			name: "smaller than 1",
			args: args{a: []uint16{2}},
			want: []uint16{2},
		},
		{
			name: "unique",
			args: args{a: []uint16{2, 4, 6, 7, 1, 3, 4, 9, 7}},
			want: []uint16{2, 4, 6, 7, 1, 3, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueUint16s(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueUint16s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueInt32s(t *testing.T) {
	type args struct {
		a []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "smaller than 1",
			args: args{a: []int32{2}},
			want: []int32{2},
		},
		{
			name: "unique",
			args: args{a: []int32{3, 4, 5, 7, 5, 4, 3, 2}},
			want: []int32{3, 4, 5, 7, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueInt32s(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueInt32s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueUint32s(t *testing.T) {
	type args struct {
		a []uint32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "smaller than 1",
			args: args{a: []uint32{2}},
			want: []uint32{2},
		},
		{
			name: "unique",
			args: args{a: []uint32{3, 4, 5, 7, 5, 4, 3, 2}},
			want: []uint32{3, 4, 5, 7, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueUint32s(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueUint32s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueInt64s(t *testing.T) {
	type args struct {
		a []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "smaller than 1",
			args: args{a: []int64{2}},
			want: []int64{2},
		},
		{
			name: "unique",
			args: args{a: []int64{3, 4, 5, 7, 5, 4, 3, 2}},
			want: []int64{3, 4, 5, 7, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueInt64s(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueInt64s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueUint64s(t *testing.T) {
	type args struct {
		a []uint64
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "smaller than 1",
			args: args{a: []uint64{2}},
			want: []uint64{2},
		},
		{
			name: "unique",
			args: args{a: []uint64{3, 4, 5, 7, 5, 4, 3, 2}},
			want: []uint64{3, 4, 5, 7, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueUint64s(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueUint64s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueFloat64s(t *testing.T) {
	type args struct {
		a []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "smaller than 1",
			args: args{a: []float64{2}},
			want: []float64{2},
		},
		{
			name: "unique",
			args: args{a: []float64{2.2, 4.2, 6.2, 7.2, 1.2, 3.2, 4.2, 9.2, 7.2}},
			want: []float64{2.2, 4.2, 6.2, 7.2, 1.2, 3.2, 9.2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueFloat64s(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueFloat64s() = %v, want %v", got, tt.want)
			}
		})
	}
}
