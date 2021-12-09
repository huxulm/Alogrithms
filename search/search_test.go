package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input  []int
	target int
	expect int
}{
	{
		input:  []int{10, 20, 80, 30, 60, 50, 110, 100, 130, 170},
		target: 110,
		expect: 6,
	},
}

func TestLinear(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.expect, Linear(c.input, c.target))
	}
}

func BenchmarkLinear(b *testing.B) {
	for _, c := range cases {
		assert.Equal(b, c.expect, Linear(c.input, c.target))
	}
}
