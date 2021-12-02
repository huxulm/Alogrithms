package sorts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_merge_sort(t *testing.T) {
	expect := []int{5, 6, 7, 11, 12, 13}
	arr := []int{12, 11, 13, 5, 6, 7}
	merge_sort(arr)
	assert.Equal(t, expect, arr)
}
