package reversewordsinastring

import "strings"

// with array
func reverseWords(sentence string) string {
	words := strings.Fields(sentence)

	left, right := 0, len(words)-1
	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}

	return strings.Join(words, " ")
}
