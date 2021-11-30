package sorts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_insert_sort(t *testing.T) {
	expect := []int{5, 6, 11, 12, 13}
	input := []int{12, 11, 13, 5, 6}
	assert.Equal(t, expect, insert_sort(input))
}
