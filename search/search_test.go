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
	var cases = []struct {
		input  []int
		target int
		expect int
	}{
		{
			input:  []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			target: 70,
			expect: 6,
		},
	}
	for _, c := range cases {
		assert.Equal(t, c.expect, Binary(c.input, c.target))
	}
}
