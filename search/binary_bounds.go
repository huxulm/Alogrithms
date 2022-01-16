// 详细总结见: https://segmentfault.com/a/1190000016825704
package search

func BinarySearchRightBound(arr []int, x int) int {
	left, right := 0, len(arr)-1

	for left < right {
		mid := left + ((right - left) >> 1) + 1
		if arr[mid] > x {
			right = mid - 1
		} else {
			left = mid
		}
	}
	if arr[right] == x {
		return right
	} else {
		return -1
	}
}

func BinarySearchLeftBound(arr []int, x int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)>>1
		if arr[mid] == x {
			right = mid - 1
		} else if arr[mid] < x {
			left = mid + 1
		} else if arr[mid] > x {
			right = mid - 1
		}
	}
	return left
}
