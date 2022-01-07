// gcd(a,b) = gcd(b,a mod b)
// a = kb + r，r = a mod b
package gcd

// Recursive finds and returns the greatest common divisor of a given integer.
func Recursive(x, y int) int {
	if y == 0 {
		return x
	}
	return Recursive(y, x%y)
}
