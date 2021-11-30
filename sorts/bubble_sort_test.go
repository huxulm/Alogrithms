package sorts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bubble_sort(t *testing.T) {
	input := []int{3, 4, 8, 9, 5, 1}
	expect := []int{1, 3, 4, 5, 8, 9}
	assert.Equal(t, expect, bubble_sort(input))
}
