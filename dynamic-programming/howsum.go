package dynamicprogramming

func howsum(target int, nums []int) []int {
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		return nil
	}
	for _, num := range nums {
		remainder := target - num
		remainderResult := howsum(remainder, nums)
		if remainderResult != nil {
			return append(remainderResult, num)
		}
	}
	return nil
}
