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

func TestAppendInt(t *testing.T) {
	type args struct {
		s []int
		e int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Append a int that does not exist in slice",
			args{[]int{1}, 2}, []int{1, 2}},
		{"Append a int that does exist in slice",
			args{[]int{1}, 1}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendInt(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendUint(t *testing.T) {
	type args struct {
		s []uint
		e uint
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{"Append a uint that does not exist in slice",
			args{[]uint{1}, 2}, []uint{1, 2}},
		{"Append a uint that does exist in slice",
			args{[]uint{1}, 1}, []uint{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendUint(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendInt8(t *testing.T) {
	type args struct {
		s []int8
		e int8
	}
	tests := []struct {
		name string
		args args
		want []int8
	}{
		{"Append a int8 that does not exist in slice",
			args{[]int8{1}, 2}, []int8{1, 2}},
		{"Append a int8 that does exist in slice",
			args{[]int8{1}, 1}, []int8{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendInt8(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendUint8(t *testing.T) {
	type args struct {
		s []uint8
		e uint8
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{"Append a uint8 that does not exist in slice",
			args{[]uint8{1}, 2}, []uint8{1, 2}},
		{"Append a uint8 that does exist in slice",
			args{[]uint8{1}, 1}, []uint8{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendUint8(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendInt16(t *testing.T) {
	type args struct {
		s []int16
		e int16
	}
	tests := []struct {
		name string
		args args
		want []int16
	}{
		{"Append a int16 that does not exist in slice",
			args{[]int16{1}, 2}, []int16{1, 2}},
		{"Append a int16 that does exist in slice",
			args{[]int16{1}, 1}, []int16{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendInt16(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendUint16(t *testing.T) {
	type args struct {
		s []uint16
		e uint16
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{"Append a uint16 that does not exist in slice",
			args{[]uint16{1}, 2}, []uint16{1, 2}},
		{"Append a uint16 that does exist in slice",
			args{[]uint16{1}, 1}, []uint16{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendUint16(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendInt32(t *testing.T) {
	type args struct {
		s []int32
		e int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{"Append a int32 that does not exist in slice",
			args{[]int32{1}, 2}, []int32{1, 2}},
		{"Append a int32 that does exist in slice",
			args{[]int32{1}, 1}, []int32{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendInt32(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendUint32(t *testing.T) {
	type args struct {
		s []uint32
		e uint32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{"Append a uint that does not exist in slice",
			args{[]uint32{1}, 2}, []uint32{1, 2}},
		{"Append a uint that does exist in slice",
			args{[]uint32{1}, 1}, []uint32{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendUint32(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendUint64(t *testing.T) {
	type args struct {
		s []uint64
		e uint64
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{"Append a uint64 that does not exist in slice",
			args{[]uint64{1, 2}, uint64(3)}, []uint64{1, 2, 3}},
		{"Append a uint64 that does exist in slice",
			args{[]uint64{1, 2}, 2}, []uint64{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendUint64(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendInt64(t *testing.T) {
	type args struct {
		s []int64
		e int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{"Append a int64 that does not exist in slice",
			args{[]int64{1, 2}, int64(3)}, []int64{1, 2, 3}},
		{"Append a int64 that does exist in slice",
			args{[]int64{1, 2}, 2}, []int64{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendInt64(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteUint(t *testing.T) {
	type args struct {
		s []uint
		e uint
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{"从uint的切片中删除第一个指定元素, 无指定元素值",
			args{[]uint{1, 2, 2, 3}, 4}, []uint{1, 2, 2, 3}},
		{"从uint的切片中删除第一个指定元素, 有指定元素值",
			args{[]uint{1, 2, 2, 3}, 2}, []uint{1, 2, 3}},
		{"从uint的切片中删除第一个指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteUint(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteInt(t *testing.T) {
	type args struct {
		s []int
		e int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"从int的切片中删除第一个指定元素, 无指定元素值",
			args{[]int{1, 2, 2, 3}, 4}, []int{1, 2, 2, 3}},
		{"从int的切片中删除第一个指定元素, 有指定元素值",
			args{[]int{1, 2, 2, 3}, 2}, []int{1, 2, 3}},
		{"从int的切片中删除第一个指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteInt(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteUint8(t *testing.T) {
	type args struct {
		s []uint8
		e uint8
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{"从uint8的切片中删除第一个指定元素, 无指定元素值",
			args{[]uint8{1, 2, 2, 3}, 4}, []uint8{1, 2, 2, 3}},
		{"从uint8的切片中删除第一个指定元素, 有指定元素值",
			args{[]uint8{1, 2, 2, 3}, 2}, []uint8{1, 2, 3}},
		{"从uint8的切片中删除第一个指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteUint8(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteInt8(t *testing.T) {
	type args struct {
		s []int8
		e int8
	}
	tests := []struct {
		name string
		args args
		want []int8
	}{
		{"从int8的切片中删除第一个指定元素, 无指定元素值",
			args{[]int8{1, 2, 2, 3}, 4}, []int8{1, 2, 2, 3}},
		{"从int8的切片中删除第一个指定元素, 有指定元素值",
			args{[]int8{1, 2, 2, 3}, 2}, []int8{1, 2, 3}},
		{"从int8的切片中删除第一个指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteInt8(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteInt16(t *testing.T) {
	type args struct {
		s []int16
		e int16
	}
	tests := []struct {
		name string
		args args
		want []int16
	}{
		{"从int16的切片中删除第一个指定元素, 无指定元素值",
			args{[]int16{1, 2, 2, 3}, 4}, []int16{1, 2, 2, 3}},
		{"从int16的切片中删除第一个指定元素, 有指定元素值",
			args{[]int16{1, 2, 2, 3}, 2}, []int16{1, 2, 3}},
		{"从int16的切片中删除第一个指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteInt16(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteUint16(t *testing.T) {
	type args struct {
		s []uint16
		e uint16
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{"从uint16的切片中删除第一个指定元素, 无指定元素值",
			args{[]uint16{1, 2, 2, 3}, 4}, []uint16{1, 2, 2, 3}},
		{"从uint16的切片中删除第一个指定元素, 有指定元素值",
			args{[]uint16{1, 2, 2, 3}, 2}, []uint16{1, 2, 3}},
		{"从uint16的切片中删除第一个指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteUint16(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteInt32(t *testing.T) {
	type args struct {
		s []int32
		e int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{"从int32的切片中删除第一个指定元素, 无指定元素值",
			args{[]int32{1, 2, 2, 3}, 4}, []int32{1, 2, 2, 3}},
		{"从int32的切片中删除第一个指定元素, 有指定元素值",
			args{[]int32{1, 2, 2, 3}, 2}, []int32{1, 2, 3}},
		{"从int32的切片中删除第一个指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteInt32(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteUint32(t *testing.T) {
	type args struct {
		s []uint32
		e uint32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{"从uint32的切片中删除第一个指定元素, 无指定元素值",
			args{[]uint32{1, 2, 2, 3}, 4}, []uint32{1, 2, 2, 3}},
		{"从uint32的切片中删除第一个指定元素, 有指定元素值",
			args{[]uint32{1, 2, 2, 3}, 2}, []uint32{1, 2, 3}},
		{"从uint32的切片中删除第一个指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteUint32(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteInt64(t *testing.T) {
	type args struct {
		s []int64
		e int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{"从int64的切片中删除第一个指定元素, 无指定元素值",
			args{[]int64{1, 2, 2, 3}, 4}, []int64{1, 2, 2, 3}},
		{"从int64的切片中删除第一个指定元素, 有指定元素值",
			args{[]int64{1, 2, 2, 3}, 2}, []int64{1, 2, 3}},
		{"从int64的切片中删除第一个指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteInt64(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteFromSliceInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteUint64(t *testing.T) {
	type args struct {
		s []uint64
		e uint64
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{"从uint64的切片中删除第一个指定元素, 无指定元素值",
			args{[]uint64{1, 2, 2, 3}, 4}, []uint64{1, 2, 2, 3}},
		{"从uint64的切片中删除第一个指定元素, 有指定元素值",
			args{[]uint64{1, 2, 2, 3}, 2}, []uint64{1, 2, 3}},
		{"从uint64的切片中删除第一个指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteUint64(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteIntAll(t *testing.T) {
	type args struct {
		s []int
		e int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"从int的切片中删除所有指定元素, 无指定元素值",
			args{[]int{1, 2, 2, 3}, 4}, []int{1, 2, 2, 3}},
		{"从int的切片中删除所有指定元素, 有指定元素值",
			args{[]int{1, 2, 2, 3}, 2}, []int{1, 3}},
		{"从int的切片中删除所有指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteIntAll(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteIntAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteUintAll(t *testing.T) {
	type args struct {
		s []uint
		e uint
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{"从uint的切片中删除所有指定元素, 无指定元素值",
			args{[]uint{1, 2, 2, 3}, 4}, []uint{1, 2, 2, 3}},
		{"从uint的切片中删除所有指定元素, 有指定元素值",
			args{[]uint{1, 2, 2, 3}, 2}, []uint{1, 3}},
		{"从uint的切片中删除所有指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteUintAll(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteUintAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteInt8All(t *testing.T) {
	type args struct {
		s []int8
		e int8
	}
	tests := []struct {
		name string
		args args
		want []int8
	}{
		{"从int8的切片中删除所有指定元素, 无指定元素值",
			args{[]int8{1, 2, 2, 3}, 4}, []int8{1, 2, 2, 3}},
		{"从int8的切片中删除所有指定元素, 有指定元素值",
			args{[]int8{1, 2, 2, 3}, 2}, []int8{1, 3}},
		{"从int8的切片中删除所有指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteInt8All(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteInt8All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteUint8All(t *testing.T) {
	type args struct {
		s []uint8
		e uint8
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{"从uint8的切片中删除所有指定元素, 无指定元素值",
			args{[]uint8{1, 2, 2, 3}, 4}, []uint8{1, 2, 2, 3}},
		{"从uint8的切片中删除所有指定元素, 有指定元素值",
			args{[]uint8{1, 2, 2, 3}, 2}, []uint8{1, 3}},
		{"从uint8的切片中删除所有指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteUint8All(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteUint8All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteInt16All(t *testing.T) {
	type args struct {
		s []int16
		e int16
	}
	tests := []struct {
		name string
		args args
		want []int16
	}{
		{"从int16的切片中删除所有指定元素, 无指定元素值",
			args{[]int16{1, 2, 2, 3}, 4}, []int16{1, 2, 2, 3}},
		{"从int16的切片中删除所有指定元素, 有指定元素值",
			args{[]int16{1, 2, 2, 3}, 2}, []int16{1, 3}},
		{"从int16的切片中删除所有指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteInt16All(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteInt16All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteUint16All(t *testing.T) {
	type args struct {
		s []uint16
		e uint16
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{"从uint16的切片中删除所有指定元素, 无指定元素值",
			args{[]uint16{1, 2, 2, 3}, 4}, []uint16{1, 2, 2, 3}},
		{"从uint16的切片中删除所有指定元素, 有指定元素值",
			args{[]uint16{1, 2, 2, 3}, 2}, []uint16{1, 3}},
		{"从uint16的切片中删除所有指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteUint16All(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteUint16All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteInt32All(t *testing.T) {
	type args struct {
		s []int32
		e int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{"从int32的切片中删除所有指定元素, 无指定元素值",
			args{[]int32{1, 2, 2, 3}, 4}, []int32{1, 2, 2, 3}},
		{"从int32的切片中删除所有指定元素, 有指定元素值",
			args{[]int32{1, 2, 2, 3}, 2}, []int32{1, 3}},
		{"从int32的切片中删除所有指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteInt32All(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteInt32All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteUint32All(t *testing.T) {
	type args struct {
		s []uint32
		e uint32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{"从uint32的切片中删除所有指定元素, 无指定元素值",
			args{[]uint32{1, 2, 2, 3}, 4}, []uint32{1, 2, 2, 3}},
		{"从uint32的切片中删除所有指定元素, 有指定元素值",
			args{[]uint32{1, 2, 2, 3}, 2}, []uint32{1, 3}},
		{"从uint32的切片中删除所有指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteUint32All(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteUint32All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteInt64All(t *testing.T) {
	type args struct {
		s []int64
		e int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{"从int64的切片中删除所有指定元素, 无指定元素值",
			args{[]int64{1, 2, 2, 3}, 4}, []int64{1, 2, 2, 3}},
		{"从int64的切片中删除所有指定元素, 有指定元素值",
			args{[]int64{1, 2, 2, 3}, 2}, []int64{1, 3}},
		{"从int64的切片中删除所有指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteInt64All(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteInt64All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteUint64All(t *testing.T) {
	type args struct {
		s []uint64
		e uint64
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{"从uint64的切片中删除所有指定元素, 无指定元素值",
			args{[]uint64{1, 2, 2, 3}, 4}, []uint64{1, 2, 2, 3}},
		{"从uint64的切片中删除所有指定元素, 有指定元素值",
			args{[]uint64{1, 2, 2, 3}, 2}, []uint64{1, 3}},
		{"从uint64的切片中删除所有指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteUint64All(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteUint64All() = %v, want %v", got, tt.want)
			}
		})
	}
}
