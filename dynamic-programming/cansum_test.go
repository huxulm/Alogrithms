package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	target int
	nums   []int
	expect bool
}{
	{5, []int{2, 2, 3}, true},
	{7, []int{2, 2, 4}, false},
	{10, []int{7, 14, 6}, false},
	{300, []int{7, 14}, false},
}

func BenchmarkCansum(b *testing.B) {
	for _, c := range cases {
		assert.Equal(b, cansum(c.target, c.nums), c.expect)
	}
}
func BenchmarkCansum_dp(b *testing.B) {
	for _, c := range cases {
		assert.Equal(b, cansum_dp(c.target, c.nums, make(map[int]bool)), c.expect)
	}
}
