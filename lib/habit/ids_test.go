package habit

import (
	"reflect"
	"testing"
)

func TestIdsGroup(t *testing.T) {
	tests := []struct {
		name string
		ids  []int64
		want string
	}{
		{
			"empty nil",
			nil,
			"",
		},
		{
			"empty",
			[]int64{},
			"",
		},
		{
			"1",
			[]int64{1},
			"1",
		},
		{
			"> 1",
			[]int64{1, 10, 11, 12},
			"1,10,11,12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IdsGroup(tt.ids); got != tt.want {
				t.Errorf("IdsGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseIdsGroup(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []int64
	}{
		{
			"empty",
			"",
			[]int64{},
		},
		{
			"1",
			"1",
			[]int64{1},
		},
		{
			"> 1",
			"1,10,11,12",
			[]int64{1, 10, 11, 12},
		},
		{
			"> 1 contain space",
			"1, 10, 11 ,  12",
			[]int64{1, 10, 11, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseIdsGroup(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseIdsGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
