package extmath

// Pow is int64 type of math.Pow function.
func Pow(x, y int64) int64 {
	if y <= 0 {
		return 1
	}
	if y%2 == 0 {
		sqrt := Pow(x, y/2)
		return sqrt * sqrt
	}
	return Pow(x, y-1) * x
}
