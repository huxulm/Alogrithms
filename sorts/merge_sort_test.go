package sorts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	expect := []int{5, 6, 7, 11, 12, 13}
	input := []int{12, 11, 13, 5, 6, 7}
	assert.Equal(t, expect, MergeSort(input))
}
