package dynamicprogramming

import (
	"fmt"
	"testing"
)

func BenchmarkHowsum(b *testing.B) {
	fmt.Println(howsum(521, []int{3, 5, 9, 11}))
}
