package longestrepeatingcharacterreplacement

func characterReplacement(s string, k int) int {
	freqS := make(map[byte]int)
	left, longestLen, mostFreqChar := 0, 0, 0

	for right := 0; right < len(s); right++ {
		// increment the current character added to the window
		freqS[s[right]]++

		// update the most frequent character if the added character is the max
		mostFreqChar = max(mostFreqChar, freqS[s[right]])

		// if the (current window length) - most freq character > replacement(k)
		// essentially we don't have enough replacements to make a larger length
		// so we remove the left bound from the freq map and then move it to the right
		if right-left+1-mostFreqChar > k {
			freqS[s[left]]--
			left++
		}
		// update the longest length if the current window if better
		longestLen = max(longestLen, right-left+1)
	}

	return longestLen
}
