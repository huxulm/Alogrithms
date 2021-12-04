package sorts

func partition(arr []int, low, high int) int {
	index := low - 1
	pivotElement := arr[high]
	for j := low; j <= high-1; j++ {
		if arr[j] <= pivotElement {
			index++
			arr[index], arr[j] = arr[j], arr[index]
		}
	}
	arr[index+1], arr[high] = arr[high], arr[index+1]
	return index + 1
}

func quickSortRange(arr []int, low, high int) {
	if len(arr) <= 1 {
		return
	}
	if low < high {
		pivot := partition(arr, low, high)
		quickSortRange(arr, low, pivot-1)
		quickSortRange(arr, pivot, high)
	}
}

// QuickSort Sorts the entire array
func QuickSort(arr []int) []int {
	quickSortRange(arr, 0, len(arr)-1)
	return arr
}
