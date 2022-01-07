package gcd

func ExtenedRecursive(a, b int64) (int64, int64, int64) {
	if b > 0 {
		d, x, y := ExtenedRecursive(b, a%b)
		y -= (a / b) * x
		return d, x, y
	}
	return a, 1, 0
}
