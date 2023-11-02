package sorts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	assert.Equal(t, selection_sort([]int{64, 25, 12, 22, 11}), []int{11, 12, 22, 25, 64})
}

func TestBubbleSort(t *testing.T) {
	input := []int{3, 4, 8, 9, 5, 1}
	expect := []int{1, 3, 4, 5, 8, 9}
	assert.Equal(t, expect, BubbleSort(input))
}

func TestInsertSort(t *testing.T) {
	expect := []int{5, 6, 11, 12, 13}
	input := []int{12, 11, 13, 5, 6}
	assert.Equal(t, expect, InsertSort(input))
}

func TestMergeSort(t *testing.T) {
	expect := []int{5, 6, 7, 11, 12, 13}
	input := []int{12, 11, 13, 5, 6, 7}
	assert.Equal(t, expect, MergeSort(input))
}

func BenchmarkQuickSort(t *testing.B) {
	input := []int{10, 80, 30, 90, 40, 50, 70}
	expect := []int{10, 30, 40, 50, 70, 80, 90}
	QuickSort(input)
	assert.Equal(t, expect, input)
}

func TestHeapSort(t *testing.T) {
	input := []int{10, 80, 30, 90, 40, 50, 70}
	expect := []int{10, 30, 40, 50, 70, 80, 90}
	assert.Equal(t, expect, HeapSort(input))
}

func BenchmarkHeapSort(t *testing.B) {
	input := []int{10, 80, 30, 90, 40, 50, 70}
	expect := []int{10, 30, 40, 50, 70, 80, 90}
	assert.Equal(t, expect, HeapSort(input))
}
