package palindrome

import (
	"regexp"
	"strings"
)

func cleanString(s string) string {
	clean_text := strings.ToLower(s)
	clean_text = strings.Join(strings.Fields(clean_text), "")
	regex, _ := regexp.Compile(`[^\p{L}\p{N}]+`) // Regular expression for alphanumeric only characters
	return regex.ReplaceAllString(clean_text, "")
}

func IsPalindrome(s string) bool {
	s = cleanString(s)
	arr := []rune(s)
	n := len(arr)

	for i := 0; i < n/2; i++ {
		if arr[i] != arr[n-1-i] {
			return false
		}
	}

	return true
}
