package palindromenumber

import "strconv"

func isPalindromeNumber(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	intStr := strconv.Itoa(x)
	length := len(intStr)
	for i := 0; i <= length/2; i++ {
		backwardIdx := length - 1 - i
		if i == backwardIdx {
			break
		}

		if intStr[i] != intStr[length-1-i] {
			return false
		}
	}
	return true
}
