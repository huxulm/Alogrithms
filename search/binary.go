package search

func Binary(arr []int, query int) int {
	var (
		left  = 0
		right = len(arr) - 1
	)

	for left <= right {
		mid := left + (right-left+1)/2
		m := arr[mid]
		if m == query {
			return mid
		} else if m < query {
			left = mid - 1
		} else {
			right = mid + 1
		}
	}

	return -1
}
