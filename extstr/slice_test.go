package strext

import (
	"reflect"
	"testing"
)

func TestAppend(t *testing.T) {
	type args struct {
		strs []string
		str  string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Append a string that does not exist in slice",
			args{[]string{"a"}, "b"}, []string{"a", "b"}},
		{"Append a string that does exist in slice",
			args{[]string{"a"}, "a"}, []string{"a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Append(tt.args.strs, tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Append() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		s []string
		e string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"从string的切片中删除第一个指定元素, 无指定元素值",
			args{[]string{"a", "b", "b", "c"}, "e"}, []string{"a", "b", "b", "c"}},
		{"从string的切片中删除第一个指定元素, 有指定元素值",
			args{[]string{"a", "b", "b", "c"}, "b"}, []string{"a", "b", "c"}},
		{"从string的切片中删除第一个指定元素, 切片是个nil",
			args{nil, "e"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Delete(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteAll(t *testing.T) {
	type args struct {
		s []string
		e string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"从string的切片中删除所有指定元素, 无指定元素值",
			args{[]string{"a", "b", "b", "c"}, "e"}, []string{"a", "b", "b", "c"}},
		{"从string的切片中删除所有指定元素, 有指定元素值",
			args{[]string{"a", "b", "b", "c"}, "b"}, []string{"a", "c"}},
		{"从string的切片中删除所有指定元素, 切片是个nil",
			args{nil, "e"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteAll(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	type args struct {
		s1 []string
		s2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Compare two slices that do have same elements and order",
			args{[]string{"1", "2", "3"}, []string{"1", "2", "3"}}, true},
		{"Compare two slices that do have same elements but does not have same order",
			args{[]string{"2", "1", "3"}, []string{"1", "2", "3"}}, false},
		{"Compare two slices that have different number of elements",
			args{[]string{"1", "2"}, []string{"1", "2", "3"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareU(t *testing.T) {
	type args struct {
		s1 []string
		s2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Compare two slices that do have same elements and order",
			args{[]string{"1", "2", "3"}, []string{"1", "2", "3"}}, true},
		{"Compare two slices that do have same elements but does not have same order",
			args{[]string{"2", "1", "3"}, []string{"1", "2", "3"}}, true},
		{"Compare two slices that do have different elements but has same count",
			args{[]string{"2", "1", "4"}, []string{"1", "2", "3"}}, false},
		{"Compare two slices that have different number of elements",
			args{[]string{"1", "2"}, []string{"1", "2", "3"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareU(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CompareU() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		sl  []string
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"字符串slice含有指定字符串,大小写敏感,一样", args{[]string{"A", "b", "c"}, "A"}, true},
		{"字符串slice不含有指定字符串,大小写敏感,因为小写的", args{[]string{"A", "b", "c"}, "a"}, false},
		{"字符串slice不含有指定字符串,大小写敏感", args{[]string{"A", "b", "c"}, "d"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.sl, tt.args.str); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsFold(t *testing.T) {
	type args struct {
		sl  []string
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"字符串slice含有指定字符串,忽略大小写,一样", args{[]string{"A", "b", "c"}, "A"}, true},
		{"字符串slice不含有指定字符串,忽略大小写,因为小写的", args{[]string{"A", "b", "c"}, "a"}, true},
		{"字符串slice不含有指定字符串,忽略大小写", args{[]string{"A", "b", "c"}, "d"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsFold(tt.args.sl, tt.args.str); got != tt.want {
				t.Errorf("ContainsFold() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	type args struct {
		a []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "smaller than 1",
			args: args{a: []string{"2"}},
			want: []string{"2"},
		},
		{
			name: "unique",
			args: args{a: []string{"a", "c", "d", "a", "e", "d", "x", "f", "c"}},
			want: []string{"a", "c", "d", "e", "x", "f"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Reverse String", args{"hello"}, "olleh"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.s); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
