package sorts

func partition(arr []int, lo, hi int) int {
	index := lo - 1
	pivot := hi
	for i := lo; i < hi; i++ {
		if arr[i] <= arr[pivot] {
			index++
			arr[index], arr[i] = arr[i], arr[index]
		}
	}
	arr[index+1], arr[hi] = arr[hi], arr[index+1]
	return index + 1
}

func quickSortRange(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}
	pivot := partition(arr, lo, hi)
	quickSortRange(arr, lo, pivot-1)
	quickSortRange(arr, pivot+1, hi)
}

// QuickSort Sorts the entire array
func QuickSort(arr []int) []int {
	quickSortRange(arr, 0, len(arr)-1)
	return arr
}
