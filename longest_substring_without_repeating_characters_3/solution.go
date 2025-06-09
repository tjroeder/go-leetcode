package longestsubstringwithoutrepeatingcharacters

// optimized sliding window which tracks the character indexes
func lengthOfLongestSubstring(s string) int {
	// initial condition
	if len(s) == 1 {
		return 1
	}
	if len(s) == 0 {
		return 0
	}

	// character to index map
	charIdx := make(map[byte]int)
	left, right, res := 0, 0, 0

	// right bound moves through s
	for right < len(s) {
		// if the right bound does not have a match add it's index location to charIdx
		if v, ok := charIdx[s[right]]; !ok {
			charIdx[s[right]] = right
		} else {
			// if v is outside of the sliding window just update charIdx and continue
			if left > v {
				charIdx[s[right]] = right
			} else {
				// if v is inside our sliding window first lets update the longest window if possible
				res = max(res, right-left)
				// then move the left bound to the right of duplicate character to start a fresh window
				left = charIdx[s[right]] + 1
				// then set the character index into the charIdx map
				charIdx[s[right]] = right
			}
		}
		// set longest window if possible
		res = max(res, right-left)
		// move to the right
		right++
	}

	// in case we hit our longest window at the end of the string check one last time to see if we beat longest window
	return max(res, right-left)
}

// simple sliding window which doesn't store index of characters
// func lengthOfLongestSubstring(s string) int {
// 	freq := make(map[byte]int)
// 	left, right, res := 0, 0, 0

// 	// right bound moves through s
// 	for right < len(s) {
// 		// add the character to the freq map
// 		freq[s[right]]++

// 		// if the freq map has duplicates we need to move left bound till we don't
// 		for freq[s[right]] > 1 {
// 			// decrement left bound value before moving the left pointer right
// 			freq[s[left]]--
// 			left++
// 		}

// 		// calculate the window length
// 		res = max(res, right-left+1)
// 		// move the right bound right
// 		right++
// 	}

// 	return res
// }
