package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark_grid_traveler(b *testing.B) {
	assert.Equal(b, grid_traveler(1, 0), 0)
	assert.Equal(b, grid_traveler(1, 1), 1)
	assert.Equal(b, grid_traveler(3, 2), 3)
	assert.Equal(b, grid_traveler(15, 16), 77558760)
}

func Benchmark_grid_traveler_dp(b *testing.B) {
	assert.Equal(b, grid_traveler_dp(1, 0, make(map[string]int)), 0)
	assert.Equal(b, grid_traveler_dp(1, 1, make(map[string]int)), 1)
	assert.Equal(b, grid_traveler_dp(3, 2, make(map[string]int)), 3)
	assert.Equal(b, grid_traveler_dp(15, 16, make(map[string]int)), 77558760)
	assert.Equal(b, grid_traveler_dp(150, 160, make(map[string]int)), 5066999816771741248)
}
