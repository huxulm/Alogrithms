package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkFibonacci(b *testing.B) {
	assert.Equal(b, fibonacci(8), 21)
}

func BenchmarkFibonacci_dp(b *testing.B) {
	assert.Equal(b, fibonacci_dp(191, make(map[int]int)), 3867759273386675073)
}

func BenchmarkFibonacci_bottom_up(b *testing.B) {
	assert.Equal(b, fibonacci_bottom_up(191), 3867759273386675073)
}
