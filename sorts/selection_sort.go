package sorts

// min=i, j=[i+1,..] 将每一轮最小值交换到前端
func selection_sort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
