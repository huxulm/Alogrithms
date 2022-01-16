package search

// Binary iterative
func Binary(arr []int, x int) int {
	var (
		left  = 0
		right = len(arr) - 1
	)

	for left <= right {
		mid := left + (right-left)>>1
		if arr[mid] == x {
			return mid
		} else if arr[mid] < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// Binary recursive
func BinaryRecursive(arr []int, lowIndex, highIndex, x int) int {
	for lowIndex <= highIndex {
		mid := lowIndex + (highIndex-lowIndex)>>1
		if arr[mid] == x {
			return mid
		} else if arr[mid] < x {
			return BinaryRecursive(arr, mid+1, highIndex, x)
		} else {
			return BinaryRecursive(arr, lowIndex, mid-1, x)
		}
	}
	return -1
}
