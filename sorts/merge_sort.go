package sorts

// Merges two subarrays of arr[].
// First subarray is arr[l..m]
// Second subarray is arr[m+1..r]
func merge(arr []int, l, m, r int) {
	var n1 = m - l + 1
	var n2 = r - m

	// Create temp arrays
	var L = make([]int, n1)
	var R = make([]int, n2)

	// Copy data to temp arrays L[] and R[]
	for i := 0; i < n1; i++ {
		L[i] = arr[l+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = arr[m+1+j]
	}

	// Merge the temp arrays back into arr[l..r]

	// Initial index of first subarray
	var i = 0

	// Initial index of second subarray
	var j = 0

	// Initial index of merged subarray
	var k = l

	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	// Copy the remaining elements of
	// L[], if there are any
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	// Copy the remaining elements of
	// R[], if there are any
	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}

// l is for left index and r is
// right index of the sub-array
// of arr to be sorted */
func mergeSort(arr []int, l, r int) {
	if l >= r {
		return //returns recursively
	}
	var m = l + (r-l)/2
	mergeSort(arr, l, m)
	mergeSort(arr, m+1, r)
	merge(arr, l, m, r)
}

func merge_sort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}
