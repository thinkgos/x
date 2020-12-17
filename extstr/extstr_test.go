package extstr

import (
	"reflect"
	"testing"
)

func TestJoin(t *testing.T) {
	type args struct {
		elems []int64
		sep   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"empty nil",
			args{
				nil,
				",",
			},
			"",
		},
		{
			"empty",
			args{
				[]int64{},
				",",
			},
			"",
		},
		{
			"1",
			args{
				[]int64{1},
				",",
			},
			"1",
		},
		{
			"> 1",
			args{
				[]int64{1, 10, 11, 12},
				",",
			},
			"1,10,11,12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Join(tt.args.elems, tt.args.sep); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJoinInt(t *testing.T) {
	type args struct {
		elems []int
		sep   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"empty nil",
			args{
				nil,
				",",
			},
			"",
		},
		{
			"empty",
			args{
				[]int{},
				",",
			},
			"",
		},
		{
			"1",
			args{
				[]int{1},
				",",
			},
			"1",
		},
		{
			"> 1",
			args{
				[]int{1, 10, 11, 12},
				",",
			},
			"1,10,11,12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinInt(tt.args.elems, tt.args.sep); got != tt.want {
				t.Errorf("JoinInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join([]int64{1, 2, 3, 4, 5, 6}, ",")
	}
}

func TestSplit(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			"empty",
			args{
				"",
				",",
			},
			[]int64{},
		},
		{
			"1",
			args{
				"1",
				",",
			},
			[]int64{1},
		},
		{
			"> 1",
			args{
				"1,10,11,12",
				",",
			},
			[]int64{1, 10, 11, 12},
		},
		{
			"> 1 contain space",
			args{
				"1, 10, 11 ,  12",
				",",
			},
			[]int64{1, 10, 11, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Split(tt.args.s, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("1,2,3,4,5", ",")
	}
}

func TestSplitInt(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"empty",
			args{
				"",
				",",
			},
			[]int{},
		},
		{
			"1",
			args{
				"1",
				",",
			},
			[]int{1},
		},
		{
			"> 1",
			args{
				"1,10,11,12",
				",",
			},
			[]int{1, 10, 11, 12},
		},
		{
			"> 1 contain space",
			args{
				"1, 10, 11 ,  12",
				",",
			},
			[]int{1, 10, 11, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitInt(tt.args.s, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
