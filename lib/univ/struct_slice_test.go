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
		New("aa")
	})

	one, two := "1", "2"
	sli := New([]StructMuch{{&one, errors.New("1")}, {&two, errors.New("2")}})
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

	sli1 := New([]int{1, 2, 3})
	require.Panics(t, func() {
		sli1.StructStringSlice("not a struct or pointer of struct")
	})

	sli2 := New([]*StructMuch{nil})
	require.Panics(t, func() {
		sli2.StructStringSlice("UID")
	})
}
