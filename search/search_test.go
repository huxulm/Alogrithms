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
		r, _ := Linear(c.input, c.target)
		assert.Equal(t, c.expect, r)
	}
}

func BenchmarkLinear(b *testing.B) {
	for _, c := range cases {
		r, _ := Linear(c.input, c.target)
		assert.Equal(b, c.expect, r)
	}
}

func TestBinary(t *testing.T) {
	for _, c := range searchTests {
		assert.Equal(t, c.expected, Binary(c.data, c.key))
	}
}

func TestBinaryRecursive(t *testing.T) {
	for _, c := range searchTests {
		assert.Equal(t, c.expected, BinaryRecursive(c.data, 0, len(c.data)-1, c.key))
	}
}

func TestBinarySearchRightBound(t *testing.T) {
	for _, c := range searchTests2 {
		assert.Equal(t, c.expected, BinarySearchRightBound(c.data, c.key))
	}
}
