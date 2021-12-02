package sorts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertSort(t *testing.T) {
	expect := []int{5, 6, 11, 12, 13}
	input := []int{12, 11, 13, 5, 6}
	assert.Equal(t, expect, InsertSort(input))
}
