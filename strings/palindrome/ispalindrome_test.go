package palindrome

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input  string
		expect bool
	}{
		{input: "a hii :ha", expect: true},
		{input: "这句中文是回文吗文回是文中句这", expect: true},
	}

	for _, test := range tests {
		if result := IsPalindrome(test.input); test.expect != result {
			t.Errorf("Input %s expect %v but got %v", test.input, test.expect, result)
		}
	}
}
