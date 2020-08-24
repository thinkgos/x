// Copyright [2020] [thinkgos] thinkgo@aliyun.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package extmath extend math
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
