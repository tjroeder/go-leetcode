package happynumber

// isHappy with fast and slow pointers
// func isHappy(n int) bool {
// 	slow, fast := n, calcSquareOfDigits(n)
// 	if fast == 1 {
// 		return true
// 	}

// 	for slow != fast {
// 		if fast == 1 {
// 			return true
// 		}
// 		slow, fast = calcSquareOfDigits(slow), calcSquareOfDigits(calcSquareOfDigits(fast))
// 	}
// 	return false
// }

// isHappy with set
func isHappy(n int) bool {
	set := make(map[int]bool)

	for n != 1 && !set[n] {
		set[n] = true
		sum := 0

		for n > 0 {
			digit := n % 10
			sum += digit * digit
			n /= 10
		}
		n = sum
	}

	return n == 1
}

// cal square of digits with strconv
// func calcSquareOfDigits(n int) int {
// 	numStr := strconv.Itoa(n)
// 	sum := 0

// 	for _, n := range numStr {
// 		digit, _ := strconv.Atoi(string(n))
// 		sum += int(math.Pow(float64(digit), 2))
// 	}
// 	return sum
// }

// cal square of digits with mod
// func calcSquareOfDigits(n int) int {
// 	sum := 0

// 	for n > 0 {
// 		digit := n % 10
// 		sum += digit * digit
// 		n /= 10
// 	}
// 	return sum
// }
