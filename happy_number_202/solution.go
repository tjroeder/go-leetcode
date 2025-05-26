package happynumber

import (
	"math"
	"strconv"
)

// without hashset
func isHappy(n int) bool {
	slow, fast := n, calcSquareOfDigits(n)
	if fast == 1 {
		return true
	}

	for slow != fast {
		if fast == 1 {
			return true
		}
		slow, fast = calcSquareOfDigits(slow), calcSquareOfDigits(calcSquareOfDigits(fast))
	}
	return false
}

func calcSquareOfDigits(n int) int {
	numStr := strconv.Itoa(n)
	sum := 0

	for _, n := range numStr {
		digit, _ := strconv.Atoi(string(n))
		sum += int(math.Pow(float64(digit), 2))
	}
	return sum
}
