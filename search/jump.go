package search

import "math"

func JumpSearch(arr []int, target int) (int, error) {

	n := len(arr)
	if n == 0 {
		return -1, ErrNotFound
	}

	step := int(math.Round(math.Sqrt(float64(n))))

	prev := 0
	curr := step

	for arr[curr-1] < target {
		prev = curr
		if prev >= n {
			return -1, ErrNotFound
		}

		curr += step

		// prevent jump over range
		if curr > n {
			curr = n
		}
	}

	for arr[prev] < target {
		prev++
		if prev == curr {
			return -1, ErrNotFound
		}
	}

	if arr[prev] == target {
		return prev, nil
	}

	return -1, ErrNotFound
}
