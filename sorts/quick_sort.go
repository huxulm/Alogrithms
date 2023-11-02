package sorts

import (
	"math/rand"
)

func QuickSort(arr []int) {
	sortRange(arr, 0, len(arr) - 1)
}

func sortRange(arr []int, lo, hi int) {
	if lo >= hi { return }
	pivot := partition(arr, lo, hi)
	sortRange(arr, lo, pivot - 1)
	sortRange(arr, pivot + 1, hi)
}

func partition(arr []int, lo, hi int) int {
	pivotIdx := lo + rand.Intn(hi - lo + 1)
	arr[pivotIdx], arr[hi] = arr[hi], arr[pivotIdx]
	i := lo
	for j := lo; j < hi; j++ {
		if arr[j] < arr[hi] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[hi] = arr[hi], arr[i]
	return i
}