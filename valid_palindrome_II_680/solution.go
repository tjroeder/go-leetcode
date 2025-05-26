package validpalindromeII

func validPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; {
		if s[l] != s[r] {
			if isPalindrome(s[l:r]) || isPalindrome(s[l+1:r+1]) {
				return true
			}
			return false
		}
		l++
		r--
	}
	return true
}

func isPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
