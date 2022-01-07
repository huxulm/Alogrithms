package lcm

import (
	"math"

	"github.com/huxulm/alogrithms/math/gcd"
)

// Lcm returns the lcm of two numbers using the fact that lcm(a,b) * gcd(a,b) = | a * b |
func Lcm(x, y int64) int64 {
	return int64(math.Abs(float64(x*y))) / int64(gcd.Iterative(x, y))
}
