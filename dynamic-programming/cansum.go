package dynamicprogramming

func cansum(target int, nums []int) bool {
	if target == 0 {
		return true
	}
	if target < 0 {
		return false
	}
	for _, num := range nums {
		remainder := target - num
		remainderResult := cansum(remainder, nums)
		if remainderResult {
			return true
		}
	}
	return false
}
func cansum_dp(target int, nums []int, memo map[int]bool) bool {
	if v, ok := memo[target]; ok {
		return v
	}
	if target == 0 {
		return true
	}
	if target < 0 {
		return false
	}
	for _, num := range nums {
		remainder := target - num
		remainderResult := cansum_dp(remainder, nums, memo)
		if remainderResult {
			memo[target] = remainderResult
			return true
		}
	}
	memo[target] = false
	return false
}
