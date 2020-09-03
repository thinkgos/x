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

import "testing"

func TestContainInt(t *testing.T) {
	type args struct {
		x  int
		vs []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				x:  4,
				vs: []int{},
			},
			want: false,
		},
		{
			name: "contain",
			args: args{
				x:  4,
				vs: []int{2, 4, 6, 7, 1, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainInt(tt.args.vs, tt.args.x); got != tt.want {
				t.Errorf("ContainInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainUint(t *testing.T) {
	type args struct {
		x  uint
		vs []uint
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				x:  4,
				vs: []uint{},
			},
			want: false,
		},
		{
			name: "contain",
			args: args{
				x:  4,
				vs: []uint{2, 4, 6, 7, 1, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainUint(tt.args.vs, tt.args.x); got != tt.want {
				t.Errorf("ContainUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainInt8(t *testing.T) {
	type args struct {
		x  int8
		vs []int8
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				x:  4,
				vs: []int8{},
			},
			want: false,
		},
		{
			name: "contain",
			args: args{
				x:  4,
				vs: []int8{2, 4, 6, 7, 1, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainInt8(tt.args.vs, tt.args.x); got != tt.want {
				t.Errorf("ContainInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainUint8(t *testing.T) {
	type args struct {
		x  uint8
		vs []uint8
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				x:  4,
				vs: []uint8{},
			},
			want: false,
		},
		{
			name: "contain",
			args: args{
				x:  4,
				vs: []uint8{2, 4, 6, 7, 1, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainUint8(tt.args.vs, tt.args.x); got != tt.want {
				t.Errorf("ContainUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainInt16(t *testing.T) {
	type args struct {
		x  int16
		vs []int16
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				x:  4,
				vs: []int16{},
			},
			want: false,
		},
		{
			name: "contain",
			args: args{
				x:  4,
				vs: []int16{2, 4, 6, 7, 1, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainInt16(tt.args.vs, tt.args.x); got != tt.want {
				t.Errorf("ContainInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainUint16(t *testing.T) {
	type args struct {
		x  uint16
		vs []uint16
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				x:  4,
				vs: []uint16{},
			},
			want: false,
		},
		{
			name: "contain",
			args: args{
				x:  4,
				vs: []uint16{2, 4, 6, 7, 1, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainUint16(tt.args.vs, tt.args.x); got != tt.want {
				t.Errorf("ContainUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainInt32(t *testing.T) {
	type args struct {
		x  int32
		vs []int32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				x:  4,
				vs: []int32{},
			},
			want: false,
		},
		{
			name: "contain",
			args: args{
				x:  4,
				vs: []int32{2, 4, 6, 7, 1, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainInt32(tt.args.vs, tt.args.x); got != tt.want {
				t.Errorf("ContainInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainUint32(t *testing.T) {
	type args struct {
		x  uint32
		vs []uint32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				x:  4,
				vs: []uint32{},
			},
			want: false,
		},
		{
			name: "contain",
			args: args{
				x:  4,
				vs: []uint32{2, 4, 6, 7, 1, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainUint32(tt.args.vs, tt.args.x); got != tt.want {
				t.Errorf("ContainUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainInt64(t *testing.T) {
	type args struct {
		x  int64
		vs []int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				x:  4,
				vs: []int64{},
			},
			want: false,
		},
		{
			name: "contain",
			args: args{
				x:  4,
				vs: []int64{5, 2, 4, 7, 6, 1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainInt64(tt.args.vs, tt.args.x); got != tt.want {
				t.Errorf("ContainInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainUint64(t *testing.T) {
	type args struct {
		x  uint64
		vs []uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				x:  4,
				vs: []uint64{},
			},
			want: false,
		},
		{
			name: "contain",
			args: args{
				x:  4,
				vs: []uint64{5, 2, 4, 7, 6, 1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainUint64(tt.args.vs, tt.args.x); got != tt.want {
				t.Errorf("ContainUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainFloat64(t *testing.T) {
	type args struct {
		x  float64
		vs []float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				x:  4,
				vs: []float64{},
			},
			want: false,
		},
		{
			name: "contain",
			args: args{
				x:  4.4,
				vs: []float64{2.3, 4.4, 6.7, 7.2, 1.9, 3.5},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainFloat64(tt.args.vs, tt.args.x); got != tt.want {
				t.Errorf("ContainFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContain(t *testing.T) {
	type args struct {
		x  interface{}
		vs []interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				x:  4,
				vs: []interface{}{},
			},
			want: false,
		},
		{
			name: "contain",
			args: args{
				x:  "iiinsomnia",
				vs: []interface{}{1, "test", "iiinsomnia", 2.9, true},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contain(tt.args.vs, tt.args.x); got != tt.want {
				t.Errorf("Contain() = %v, want %v", got, tt.want)
			}
		})
	}
}
