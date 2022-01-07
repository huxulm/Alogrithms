package prime

func Factorize(n int64) map[int64]int64 {
	res := make(map[int64]int64)

	for i := int64(2); i*i <= n; i++ {
		for {
			if n%i != 0 {
				break
			}
			res[i]++
			n /= i
		}
	}

	if n > 1 {
		res[n]++
	}
	return res
}
