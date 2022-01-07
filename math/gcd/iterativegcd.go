package gcd

func Iterative(x, y int64) int64 {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
