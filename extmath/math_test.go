package extmath

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPow(t *testing.T) {
	type args struct {
		x int64
		y int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"power <=0",
			args{5, -1},
			1,
		},
		{
			"power 3",
			args{5, 3},
			125,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pow(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Pow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPow(b *testing.B) {
	x := rand.Int63n(100)
	y := rand.Int63n(6)
	for i := 0; i < b.N; i++ {
		Pow(x, y)
	}
}

func TestRound(t *testing.T) {
	type args struct {
		f float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"四舍",
			args{100.004, 2},
			100.00,
		},
		{
			"五入",
			args{100.005, 2},
			100.01,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Round(tt.args.f, tt.args.n); got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRound(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Round(10.45646841318513515, 9)
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		name   string
		number int64
		want   int64
	}{
		{
			"正数",
			100,
			100,
		},
		{
			"负数",
			-100,
			100,
		},
		{
			"负极数",
			math.MinInt64 + 1,
			math.MaxInt64,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.number); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange(t *testing.T) {
	var start, end = 1, 5
	res0 := Range(start, end)
	require.True(t, len(res0) == 5 && res0[0] == start && res0[len(res0)-1] == end)

	start, end = 5, 1
	res1 := Range(start, end)
	require.True(t, len(res1) == 5 && res1[0] == start && res1[len(res0)-1] == end)

	start, end = 3, 3
	res2 := Range(start, end)
	require.True(t, len(res2) == 1)
}
