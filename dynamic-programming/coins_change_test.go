package dynamicprogramming

import "testing"

func TestCoinsChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 9
	expect := 3

	result := coinChange(coins, amount)
	if result != expect {
		t.Errorf("Expect %d but got: %d", expect, result)
	}
}
