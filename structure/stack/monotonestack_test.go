package stack

import (
	"reflect"
	"testing"
)

func TestMonotoneStack(t *testing.T) {
	tests := []struct {
		input, expect []int
	}{
		{
			input:  []int{81, 45, 63, 11, 8, 23, 1, 7},
			expect: []int{7, 23, 63, 81},
		},
		{
			input:  []int{0, 0, 0, 1, 1, 1, 3, 3, 2, 1},
			expect: []int{7, 23, 63, 81},
		},
	}

	for _, test := range tests {
		sta := new(MonotoneStack)
		for i := range test.input {
			sta.Push(test.input[i])
		}
		if result := sta.Data(); !reflect.DeepEqual(test.expect, result) {
			t.Errorf("Input %v expect %v but got %v", test.input, test.input, result)
		}
	}
}
