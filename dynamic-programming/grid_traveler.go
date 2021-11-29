package dynamicprogramming

import "fmt"

func grid_traveler(m, n int) int {
	if m == 1 && n == 1 {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}
	return grid_traveler(m-1, n) + grid_traveler(m, n-1)
}

func grid_traveler_dp(m, n int, memo map[string]int) int {
	key := fmt.Sprintf("%d,%d", m, n)
	if v, ok := memo[key]; ok {
		return v
	}
	if m == 1 && n == 1 {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}
	memo[key] = grid_traveler_dp(m-1, n, memo) + grid_traveler_dp(m, n-1, memo)
	return memo[key]
}
