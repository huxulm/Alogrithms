package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJumpSearch(t *testing.T) {
	for _, c := range searchTests {
		r, err := JumpSearch(c.data, c.key)
		assert.Equal(t, c.expected, r)
		assert.Equal(t, c.expectedError, err)
	}
}
