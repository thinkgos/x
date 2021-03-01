package extrand

// Intx 随机[min,max)中的值
func Intx(min, max int) int {
	if min > max {
		panic("invalid argument to Int")
	}
	if min == max {
		return min
	}
	return Intn(max-min) + min
}

// Int31x 随机[min,max)中的值
func Int31x(min, max int32) int32 {
	if min > max {
		panic("invalid argument to Int31")
	}
	if min == max {
		return min
	}
	return Int31n(max-min) + min
}

// Int63x 随机[min,max)中的值
func Int63x(min, max int64) int64 {
	if min > max {
		panic("invalid argument to Int63")
	}
	if min == max {
		return min
	}
	return Int63n(max-min) + min
}

// Float64x 随机min,max中的值
func Float64x(min, max float64) float64 {
	if min > max {
		panic("invalid argument to Float64")
	}
	return min + (max-min)*Float64()
}
