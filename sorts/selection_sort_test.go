package sorts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_selection_sort(t *testing.T) {
	assert.Equal(t, selection_sort([]int{64, 25, 12, 22, 11}), []int{11, 12, 22, 25, 64})
}
