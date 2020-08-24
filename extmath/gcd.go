package extmath

// GCD/Greatest Common Divisor

// Gcd get Greatest Common Divisor
func Gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return Gcd(y, x%y)
}

// Gcdx get Greatest Common Divisor
func Gcdx(x, y int) int {
	if y == 0 {
		return x
	}

	var tmp int

	for {
		tmp = (x % y)
		if tmp > 0 {
			x = y
			y = tmp
		} else {
			return y
		}
	}
}

// GcdSlice get Greatest Common Divisor
func GcdSlice(n []int) int {
	switch len(n) {
	case 0:
		return 1
	case 1:
		return n[0]
	default:
		g := n[0]
		for i := 1; i < len(n); i++ {
			g = Gcdx(g, n[i])
		}
		return g
	}
}

// Lcm get least common multiple
func Lcm(x, y int) int {
	return x * y / Gcdx(x, y)
}
