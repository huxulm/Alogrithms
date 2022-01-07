package prime

import (
	"reflect"
	"testing"
)

func TestFactorize(t *testing.T) {
	var tests = []struct {
		n        int64
		expected map[int64]int64
	}{
		{4, map[int64]int64{2: 2}},
		{5, map[int64]int64{5: 1}},
		{7, map[int64]int64{7: 1}},
		{10, map[int64]int64{2: 1, 5: 1}},
		{999, map[int64]int64{3: 3, 37: 1}},
		{999999999999878, map[int64]int64{2: 1, 19: 1, 26315789473681: 1}},
	}

	for _, c := range tests {
		r := Factorize(c.n)
		if !reflect.DeepEqual(c.expected, r) {
			t.Errorf("Wrong result! Expect: %v but got %v", c.expected, r)
		}
	}
}
