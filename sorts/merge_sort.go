package sorts

// Merges two subarrays of arr[].
// First subarray is arr[l..m]
// Second subarray is arr[m+1..r]
func merge(a, b []int) []int {
	var r = make([]int, len(a)+len(b))

	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}
	}

	for i < len(a) {
		r[i+j] = a[i+j]
		i++
	}

	for j < len(b) {
		r[i+j] = b[i+j]
		j++
	}

	return r
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr //returns recursively
	}
	var m = len(arr) / 2
	a := MergeSort(arr[:m])
	b := MergeSort(arr[m:])
	return merge(a, b)
}
