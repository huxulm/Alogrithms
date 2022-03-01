package kmp

import (
	"testing"
)

func TestKmpSearch(t *testing.T) {
	tests := []struct {
		s, pat string
		expect int
	}{
		{pat: "ababc", s: "aacababcca", expect: 3},
		{pat: "dabc", s: "aacddabcca", expect: 4},
	}

	for _, test := range tests {
		KMP := Build(test.pat)
		if result := KMP.Search(test.s); result != test.expect {
			t.Errorf("Input s: %s pat: %s expect %d got %d", test.s, test.pat, test.expect, result)
		}
	}
}
func TestSearch(t *testing.T) {
	tests := []struct {
		s, pat string
		expect int
	}{
		{pat: "ababc", s: "aacababcca", expect: 3},
		{pat: "dabc", s: "aacddabcca", expect: 4},
	}

	for _, test := range tests {
		if result := Search(test.s, test.pat); result != test.expect {
			t.Errorf("Input s: %s pat: %s expect %d got %d", test.s, test.pat, test.expect, result)
		}
	}
}
