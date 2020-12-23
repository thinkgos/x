package univ

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeysOfMap(t *testing.T) {
	require.Panics(t, func() { KeysOfMap("no map") })
	require.Panics(t, func() { KeysOfMap(map[int]struct{}{1: {}}) })
	require.Equal(t, []string{}, KeysOfMap(nil))
	require.Equal(t, []string{}, KeysOfMap(map[int]struct{}{}))

	want := []string{"1", "2", "3", "4"}
	tests := []struct {
		name string
		m    interface{}
		want []string
	}{
		{
			"no ptr 1",
			map[string]struct{}{
				"1": {},
				"2": {},
				"3": {},
				"4": {},
			},
			want,
		},
		{
			"ptr 1",
			&map[string]struct{}{
				"1": {},
				"2": {},
				"3": {},
				"4": {},
			},
			want,
		},
		{
			"no ptr 2",
			map[string]string{
				"1": "value",
				"2": "value",
				"3": "value",
				"4": "value",
			},
			want,
		},
		{
			"ptr 2",
			&map[string]string{
				"1": "value",
				"2": "value",
				"3": "value",
				"4": "value",
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KeysOfMap(tt.m)
			sort.Strings(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeysIntOfMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Int64Slice attaches the methods of Interface to []int64, sorting a increasing order.
type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func TestKeysIntOfMap(t *testing.T) {
	require.Panics(t, func() { KeysIntOfMap("no map") })
	require.Panics(t, func() { KeysIntOfMap(map[string]struct{}{"1": {}}) })
	require.Equal(t, []int64{}, KeysIntOfMap(nil))
	require.Equal(t, []int64{}, KeysIntOfMap(map[string]struct{}{}))

	tests := []struct {
		name string
		m    interface{}
		want []int64
	}{
		{
			"no ptr",
			map[int]string{
				1: "value",
				2: "value",
				3: "value",
				4: "value",
			},
			[]int64{1, 2, 3, 4},
		},
		{
			"ptr",
			&map[int]string{
				1: "value",
				2: "value",
				3: "value",
				4: "value",
			},
			[]int64{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KeysIntOfMap(tt.m)
			sort.Sort(Int64Slice(got))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeysIntOfMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
