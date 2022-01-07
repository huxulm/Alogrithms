package prime

// NaiveApproach checks if an integer is prime or not. Returns a bool.
func NaiveApproach(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// PairApproach checks primality of an integer and returns a bool. More efficient than the naive approach as number of iterations are less.
func PairApproach(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
