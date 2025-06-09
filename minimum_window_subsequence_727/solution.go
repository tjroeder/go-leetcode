package minimumwindowsubsequence

func minWindow(s1 string, s2 string) string {
	s1Len, s2Len := len(s1), len(s2)
	minSubLen := s1Len + 1
	minSubIdx := 0
	s1Idx, s2Idx := 0, 0

	for s1Idx < s1Len {
		if s1[s1Idx] == s2[s2Idx] {
			s2Idx++
			if s2Idx == s2Len {
				start, end := s1Idx-1, s1Idx
				s2Idx = s2Idx - 2
				for s2Idx >= 0 {
					if s1[start] == s2[s2Idx] {
						if s2Idx == 0 {
							break
						}
						s2Idx--
					}
					start--
				}

				// currLen := end - start + 1
				if len(s1[start:end+1]) < minSubLen {
					minSubLen = len(s1[start:end])
					minSubIdx = start
				}
				s2Idx = 0
				s1Idx = start
			}
		}
		s1Idx++
	}
	if minSubLen > s1Len {
		return ""
	}
	return s1[minSubIdx : minSubIdx+minSubLen+1]
}

// old version
// func minWindow(s string, t string) string {
// 	sLen, tLen := len(s), len(t)
// 	minSubLen := math.Inf(1)
// 	minSubSequence := ""
// 	sIdx, tIdx := 0, 0

// 	// iterate over s
// 	for sIdx < sLen {
// 		// if character at s[sIdx] matches t[tIdx] we know we need to increment
// 		// both values at minimum, we may not be at the end of t (substring) yet
// 		if s[sIdx] == t[tIdx] {
// 			tIdx++

// 			// if we've reached the end of our substring t
// 			if tIdx == tLen {
// 				// now we need to back track to find the true minimum
// 				// start will go back till we find t[0] in s
// 				start, end := sIdx-1, sIdx
// 				tIdx--
// 				tIdx--
// 				for tIdx >= 0 {
// 					// decrement tIdx only if there is a match always decrement start
// 					if s[start] == t[tIdx] {
// 						if tIdx == 0 {
// 							break
// 						}
// 						tIdx--
// 					}
// 					start--
// 				}

// 				currLen := float64(end - start + 1)
// 				if currLen < minSubLen {
// 					minSubLen = currLen
// 					minSubSequence = s[start : end+1]
// 				}
// 				sIdx = start
// 				tIdx = 0
// 			}
// 		}
// 		// no match so keep moving through s
// 		sIdx++
// 	}
// 	return minSubSequence
// }
