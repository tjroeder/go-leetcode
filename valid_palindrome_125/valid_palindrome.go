package validpalindrome

import (
	"unicode"
)

func isValidPalindrome(s string) bool {
	runes := []rune(s)
	l, r := 0, len(runes)-1

	for l < r {
		for l < r && !unicode.IsLetter(runes[l]) && !unicode.IsDigit(runes[l]) {
			l++
		}
		for l < r && !unicode.IsLetter(runes[r]) && !unicode.IsDigit(runes[r]) {
			r--
		}

		if unicode.ToLower(runes[l]) != unicode.ToLower(runes[r]) {
			return false
		}
		l++
		r--
	}

	return true
}
