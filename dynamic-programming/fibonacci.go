package dynamicprogramming

func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacci_dp(n int, memo map[int]int) int {
	if v, ok := memo[n]; ok {
		return v
	}
	if n == 1 || n == 2 {
		return 1
	}
	memo[n] = fibonacci_dp(n-1, memo) + fibonacci_dp(n-2, memo)
	return memo[n]
}

func fibonacci_bottom_up(n int) int {
	if n < 2 {
		return n
	}
	var t1, t2 = 1, 1
	for i := 2; i < n; i++ {
		t0 := t1 + t2
		t1 = t2
		t2 = t0
	}
	return t2
}
