package nextpalidrome

func findNextPermutation(digits []byte) bool {
	// Step 1: Find the first digit that is smaller than the digit after it
	i := len(digits) - 2
	for i >= 0 && digits[i] >= digits[i+1] {
		i--
	}
	if i == -1 {
		return false
	}

	// Step 2: Find the next largest digit to swap with digits[i]
	j := len(digits) - 1
	for digits[j] <= digits[i] {
		j--
	}

	// Step 3: Swap and reverse the rest to get the smallest next permutation
	digits[i], digits[j] = digits[j], digits[i]
	reverse(digits[i+1:])
	return true
}

func reverse(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func nextPalindrome(numStr string) string {
	n := len(numStr)

	if n == 1 {
		return ""
	}

	halfLength := n / 2
	leftHalf := []byte(numStr[:halfLength])

	// Step 1: Get the next permutation for the left half
	if !findNextPermutation(leftHalf) {
		return ""
	}

	// Step 2: Form the next palindrome by mirroring the left half
	var nextPalindrome string
	if n%2 == 0 {
		nextPalindrome = string(leftHalf) + string(reverseBytes(leftHalf))
	} else {
		middleDigit := numStr[halfLength]
		nextPalindrome = string(leftHalf) + string(middleDigit) + string(reverseBytes(leftHalf))
	}

	// Step 3: Ensure the result is larger than the original number
	if nextPalindrome > numStr {
		return nextPalindrome
	}
	return ""
}

func reverseBytes(arr []byte) []byte {
	reversed := make([]byte, len(arr))
	copy(reversed, arr)
	reverse(reversed)
	return reversed
}
