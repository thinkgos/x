package extmath

import (
	"testing"
)

func TestGcd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1 2",
			args{1, 2},
			1,
		},
		{
			"3 0",
			args{3, 0},
			3,
		},
		{
			"20 35",
			args{20, 35},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gcd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGcdx(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1 2",
			args{1, 2},
			1,
		},
		{
			"3 0",
			args{3, 0},
			3,
		},
		{
			"20 35",
			args{20, 35},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gcdx(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Gcdx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetGcd(t *testing.T) {
	type args struct {
		n []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"empty",
			args{[]int{}},
			1,
		},
		{
			"5",
			args{[]int{5}},
			5,
		},
		{
			"45,35",
			args{[]int{35, 45}},
			5,
		},
		{
			"4,4,6,2,8,10",
			args{[]int{4, 4, 6, 2, 8, 10}},
			2,
		},
		{
			"empty",
			args{[]int{35, 105, 45, 80, 55}},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GcdSlice(tt.args.n); got != tt.want {
				t.Errorf("GcdSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLcm(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"3,5",
			args{3, 5},
			15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lcm(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Lcm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGcd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Gcd(1000, 2005)
	}
}

func BenchmarkGcdx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Gcdx(1000, 2005)
	}
}

func BenchmarkGcdSlice(b *testing.B) {
	v := []int{35, 105, 45, 80, 55}
	for i := 0; i < b.N; i++ {
		GcdSlice(v)
	}
}
