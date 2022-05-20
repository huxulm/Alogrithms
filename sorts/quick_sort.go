package sorts

import (
	"math"
	"math/rand"
)

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
	pivot := randomPartition(arr, lo, hi)
	quickSortRange(arr, lo, pivot-1)
	quickSortRange(arr, pivot+1, hi)
}

func init() {
	rand.Seed(math.MaxInt64)
}

//《算法导论》9.2：期望为线性的选择算法
func randomPartition(arr []int, lo, hi int) int {
	pivot := rand.Int()%(hi-lo+1) + lo
	arr[pivot], arr[hi] = arr[hi], arr[pivot]
	return partition(arr, lo, hi)
}

// QuickSort Sorts the entire array
func QuickSort(arr []int) []int {
	quickSortRange(arr, 0, len(arr)-1)
	return arr
}
