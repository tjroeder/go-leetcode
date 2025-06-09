package minimumwindowsubstring

// my looong but working leetcode, think it is a little more readible
// though than the one below I found on leetcode.
func minWindow(s, t string) string {
	// check initial conditions
	if s == "" || t == "" || len(s) < len(t) {
		return ""
	}

	// make a freqency map of the characters we are looking for
	tMap := make(map[rune]int)
	for _, char := range t {
		tMap[char]++
	}

	sLen, count := len(s), len(t)
	// minSubLen could also be math.MaxInt or int(^uint(0)>>1) which is
	// bitwise NOT (^) unit(0) then shifted to the right one to get rid
	// of the leading zero. Then converted to int... I think it's easier to
	// just know that we shouldn't be greater than the length of s...
	minSubLen := sLen + 1
	left, right := 0, 0
	startIdx := 0

	// move left through s to find the left bound, while right will
	// do the right.
	for right < sLen {
		// if we get a hit on the map we always decrement the character value
		if _, ok := tMap[rune(s[right])]; ok {
			// we only decrement the count though if we are still >0 on the character values
			// this is because if the value is negative we don't want to double count
			// that character and make us think that we have less characters to find
			if tMap[rune(s[right])] > 0 {
				count--
			}
			tMap[rune(s[right])]--
		}
		// move pointer to the right
		right++

		// if we have no characters left to find enter a loop to see if we
		// can shrink the window
		for count == 0 {
			// if the left and right bounds of the window are less than the
			// last substring length, set the startIdx to the left bound of
			// the window, this in combination with the length will give the
			// substring in the end rather than continuously updating a substring
			// variable
			if right-left < minSubLen {
				startIdx = left
				minSubLen = len(s[left:right])
			}

			// if there is a hit on the left pointer, increase the character value before
			// we move the pointer to the right
			if _, ok := tMap[rune(s[left])]; ok {
				tMap[rune(s[left])]++
				// only increase the count if we are not negative on the character value
				// this is because if we are still 0 or negative it means we have already
				// met our desired amount for the substring
				if tMap[rune(s[left])] > 0 {
					count++
				}
			}
			left++
		}
	}

	// this only happens if we don't find a matching substring(t) in s
	if minSubLen > sLen {
		return ""
	}
	return s[startIdx : startIdx+minSubLen]
}

// little more elegant version from leetcode
// this one just uses an int slice with 128 ints for the
// the latin alphabet range, this has the benefit of not checking
// runes as well it keeps track of all characters not just the
// ones we are looking for in t, but only decreases the count if we
// are looking for that character. It also keeps track of where the
// left pointer previously was for the substring, rather than continuously
// updating a substring value holder (old note---like my minSubString return.)
// func minWindow(s string, t string) string {
// 	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
// 		return ""
// 	}

// 	mapS := make([]int, 128)
// 	count := len(t)
// 	start, end := 0, 0
// 	minLen, startIndex := int(^uint(0)>>1), 0
// 	/// UPVOTE !
// 	for _, char := range t {
// 		mapS[char]++
// 	}

// 	for end < len(s) {
// 		if mapS[s[end]] > 0 {
// 			count--
// 		}
// 		mapS[s[end]]--
// 		end++

// 		for count == 0 {
// 			if end-start < minLen {
// 				startIndex = start
// 				minLen = end - start
// 			}

// 			if mapS[s[start]] == 0 {
// 				count++
// 			}
// 			mapS[s[start]]++
// 			start++
// 		}
// 	}

// 	if minLen == int(^uint(0)>>1) {
// 		return ""
// 	}

// 	return s[startIndex : startIndex+minLen]
// }
