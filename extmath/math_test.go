package extmath

import (
	"math/rand"
	"testing"
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
