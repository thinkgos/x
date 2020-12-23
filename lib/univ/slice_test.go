package univ

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

type StructInt8 struct {
	UID   int8
	Value string
}

type StructUint16 struct {
	UID   uint16
	Value string
}

type StructFloat64 struct {
	UID   float64
	Value string
}

type StructFloat32 struct {
	UID   float32
	Value string
}

type StructString struct {
	UID   string
	Value string
}

type StructMuch struct {
	UID *string
	Err error
}

func TestStructsIntSlice(t *testing.T) {
	type args struct {
		s         interface{}
		fieldName string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"no ptr",
			args{
				[]StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]int{1, 2},
		},
		{
			"ptr",
			args{
				[]*StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]int{1, 2},
		},
		{
			"no ptr",
			args{
				[]StructUint16{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]int{1, 2},
		},
		{
			"ptr",
			args{
				[]*StructUint16{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]int{1, 2},
		},
		{
			"no ptr",
			args{
				[]StructFloat64{{1.1, "1"}, {2, "2"}},
				"UID",
			},
			[]int{1, 2},
		},
		{
			"ptr",
			args{
				[]*StructFloat64{{1.1, "1"}, {2, "2"}},
				"UID",
			},
			[]int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructsIntSlice(tt.args.s, tt.args.fieldName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructsIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructsUintSlice(t *testing.T) {
	type args struct {
		s         interface{}
		fieldName string
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{
			"no ptr",
			args{
				[]StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint{1, 2},
		},
		{
			"ptr",
			args{
				[]*StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint{1, 2},
		},
		{
			"no ptr",
			args{
				[]StructUint16{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint{1, 2},
		},
		{
			"ptr",
			args{
				[]*StructUint16{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint{1, 2},
		},
		{
			"no ptr",
			args{
				[]StructFloat64{{1.1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint{1, 2},
		},
		{
			"ptr",
			args{
				[]*StructFloat64{{1.1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructsUintSlice(tt.args.s, tt.args.fieldName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructsUintSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructsInt64Slice(t *testing.T) {
	type args struct {
		s         interface{}
		fieldName string
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			"no ptr",
			args{
				[]StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]int64{1, 2},
		},
		{
			"ptr",
			args{
				[]*StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]int64{1, 2},
		},
		{
			"no ptr",
			args{
				[]StructUint16{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]int64{1, 2},
		},
		{
			"ptr",
			args{
				[]*StructUint16{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]int64{1, 2},
		},
		{
			"no ptr",
			args{
				[]StructFloat64{{1.1, "1"}, {2, "2"}},
				"UID",
			},
			[]int64{1, 2},
		},
		{
			"ptr",
			args{
				[]*StructFloat64{{1.1, "1"}, {2, "2"}},
				"UID",
			},
			[]int64{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructsInt64Slice(tt.args.s, tt.args.fieldName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructsInt64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructsUint64Slice(t *testing.T) {
	type args struct {
		s         interface{}
		fieldName string
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			"no ptr",
			args{
				[]StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint64{1, 2},
		},
		{
			"ptr",
			args{
				[]*StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint64{1, 2},
		},
		{
			"no ptr",
			args{
				[]StructUint16{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint64{1, 2},
		},
		{
			"ptr",
			args{
				[]*StructUint16{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint64{1, 2},
		},
		{
			"no ptr",
			args{
				[]StructFloat64{{1.1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint64{1, 2},
		},
		{
			"ptr",
			args{
				[]*StructFloat64{{1.1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint64{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructsUint64Slice(tt.args.s, tt.args.fieldName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructsUint64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructStringSlice(t *testing.T) {
	type args struct {
		s         interface{}
		fieldName string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"integer no ptr",
			args{
				[]StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]string{"1", "2"},
		},
		{
			"integer ptr",
			args{
				[]*StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]string{"1", "2"},
		},
		{
			"integer no ptr",
			args{
				[]StructUint16{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]string{"1", "2"},
		},
		{
			"integer ptr",
			args{
				[]*StructUint16{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]string{"1", "2"},
		},
		{
			"string no ptr",
			args{
				[]*StructString{{"1", "1"}, {"2", "2"}},
				"UID",
			},
			[]string{"1", "2"},
		},
		{
			"string ptr",
			args{
				[]*StructString{{"1", "1"}, {"2", "2"}},
				"UID",
			},
			[]string{"1", "2"},
		},
		{
			"Float32 no ptr",
			args{
				[]StructFloat32{{1.1, "1"}, {2, "2"}},
				"UID",
			},
			[]string{"1.1", "2"},
		},
		{
			"Float32 ptr",
			args{
				[]*StructFloat32{{1.1, "1"}, {2, "2"}},
				"UID",
			},
			[]string{"1.1", "2"},
		},
		{
			"Float64 no ptr",
			args{
				[]StructFloat64{{1.1, "1"}, {2, "2"}},
				"UID",
			},
			[]string{"1.1", "2"},
		},
		{
			"Float64 ptr",
			args{
				[]*StructFloat64{{1.1, "1"}, {2, "2"}},
				"UID",
			},
			[]string{"1.1", "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructStringSlice(tt.args.s, tt.args.fieldName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlice(t *testing.T) {
	require.Panics(t, func() {
		NewSlice("aa")
	})

	one, two := "1", "2"
	sli := NewSlice([]StructMuch{{&one, errors.New("1")}, {&two, errors.New("2")}})
	t.Log(sli.Name())
	require.Panics(t, func() {
		sli.StructIntSlice("Err")
	})

	require.Panics(t, func() {
		sli.StructStringSlice("Err")
	})
	require.Panics(t, func() {
		sli.StructStringSlice("NotExist")
	})

	sli1 := NewSlice([]int{1, 2, 3})
	require.Panics(t, func() {
		sli1.StructStringSlice("not a struct or pointer of struct")
	})

	sli2 := NewSlice([]*StructMuch{nil})
	require.Panics(t, func() {
		sli2.StructStringSlice("UID")
	})
}

func TestIntSlice(t *testing.T) {
	require.Panics(t, func() {
		IntSlice([]string{"1", "2"})
	})
	tests := []struct {
		name string
		s    interface{}
		want []int
	}{
		{
			"int",
			[]int{1, 2, 3, 4},
			[]int{1, 2, 3, 4},
		},
		{
			"int",
			[]uint{1, 2, 3, 4},
			[]int{1, 2, 3, 4},
		},
		{
			"int",
			[]float64{1.1, 2.1, 3, 4},
			[]int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntSlice(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUintSlice(t *testing.T) {
	require.Panics(t, func() {
		UintSlice([]string{"1", "2"})
	})
	tests := []struct {
		name string
		s    interface{}
		want []uint
	}{
		{
			"int",
			[]int{1, 2, 3, 4},
			[]uint{1, 2, 3, 4},
		},
		{
			"int",
			[]uint{1, 2, 3, 4},
			[]uint{1, 2, 3, 4},
		},
		{
			"int",
			[]float64{1.1, 2.1, 3, 4},
			[]uint{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UintSlice(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UintSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8Slice(t *testing.T) {
	require.Panics(t, func() {
		Int8Slice([]string{"1", "2"})
	})
	tests := []struct {
		name string
		s    interface{}
		want []int8
	}{
		{
			"int",
			[]int{1, 2, 3, 4},
			[]int8{1, 2, 3, 4},
		},
		{
			"int",
			[]uint{1, 2, 3, 4},
			[]int8{1, 2, 3, 4},
		},
		{
			"int",
			[]float64{1.1, 2.1, 3, 4},
			[]int8{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8Slice(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int8Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8Slice(t *testing.T) {
	require.Panics(t, func() {
		Uint8Slice([]string{"1", "2"})
	})
	tests := []struct {
		name string
		s    interface{}
		want []uint8
	}{
		{
			"int",
			[]int{1, 2, 3, 4},
			[]uint8{1, 2, 3, 4},
		},
		{
			"int",
			[]uint{1, 2, 3, 4},
			[]uint8{1, 2, 3, 4},
		},
		{
			"int",
			[]float64{1.1, 2.1, 3, 4},
			[]uint8{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint8Slice(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint8Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16Slice(t *testing.T) {
	require.Panics(t, func() {
		Int16Slice([]string{"1", "2"})
	})
	tests := []struct {
		name string
		s    interface{}
		want []int16
	}{
		{
			"int",
			[]int{1, 2, 3, 4},
			[]int16{1, 2, 3, 4},
		},
		{
			"int",
			[]uint{1, 2, 3, 4},
			[]int16{1, 2, 3, 4},
		},
		{
			"int",
			[]float64{1.1, 2.1, 3, 4},
			[]int16{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16Slice(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int16Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16Slice(t *testing.T) {
	require.Panics(t, func() {
		Uint16Slice([]string{"1", "2"})
	})
	tests := []struct {
		name string
		s    interface{}
		want []uint16
	}{
		{
			"int",
			[]int{1, 2, 3, 4},
			[]uint16{1, 2, 3, 4},
		},
		{
			"int",
			[]uint{1, 2, 3, 4},
			[]uint16{1, 2, 3, 4},
		},
		{
			"int",
			[]float64{1.1, 2.1, 3, 4},
			[]uint16{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint16Slice(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint16Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Slice(t *testing.T) {
	require.Panics(t, func() {
		Int32Slice([]string{"1", "2"})
	})
	tests := []struct {
		name string
		s    interface{}
		want []int32
	}{
		{
			"int",
			[]int{1, 2, 3, 4},
			[]int32{1, 2, 3, 4},
		},
		{
			"int",
			[]uint{1, 2, 3, 4},
			[]int32{1, 2, 3, 4},
		},
		{
			"int",
			[]float64{1.1, 2.1, 3, 4},
			[]int32{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32Slice(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32Slice(t *testing.T) {
	require.Panics(t, func() {
		Uint32Slice([]string{"1", "2"})
	})
	tests := []struct {
		name string
		s    interface{}
		want []uint32
	}{
		{
			"int",
			[]int{1, 2, 3, 4},
			[]uint32{1, 2, 3, 4},
		},
		{
			"int",
			[]uint{1, 2, 3, 4},
			[]uint32{1, 2, 3, 4},
		},
		{
			"int",
			[]float64{1.1, 2.1, 3, 4},
			[]uint32{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint32Slice(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Slice(t *testing.T) {
	require.Panics(t, func() {
		Int64Slice([]string{"1", "2"})
	})
	tests := []struct {
		name string
		s    interface{}
		want []int64
	}{
		{
			"int",
			[]int{1, 2, 3, 4},
			[]int64{1, 2, 3, 4},
		},
		{
			"int",
			[]uint{1, 2, 3, 4},
			[]int64{1, 2, 3, 4},
		},
		{
			"int",
			[]float64{1.1, 2.1, 3, 4},
			[]int64{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64Slice(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64Slice(t *testing.T) {
	require.Panics(t, func() {
		Uint64Slice([]string{"1", "2"})
	})
	tests := []struct {
		name string
		s    interface{}
		want []uint64
	}{
		{
			"int",
			[]int{1, 2, 3, 4},
			[]uint64{1, 2, 3, 4},
		},
		{
			"int",
			[]uint{1, 2, 3, 4},
			[]uint64{1, 2, 3, 4},
		},
		{
			"int",
			[]float64{1.1, 2.1, 3, 4},
			[]uint64{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64Slice(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}
