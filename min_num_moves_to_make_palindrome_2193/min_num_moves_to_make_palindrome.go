package minnummovestomakepalindrome

// first solution need to simplify
// func minMovesToMakePalindrome(s string) int {
// 	moves := 0
// 	runes := []rune(s)

// 	for l, r := 0, len(runes)-1; l < r; {
// 		isMiddle := false
// 		if runes[l] != runes[r] {
// 			currR := r
// 			for runes[l] != runes[r] {
// 				if runes[l] != runes[currR] {
// 					currR--
// 					continue
// 				}

// 				if l == currR {
// 					moves += (len(runes) / 2) - l
// 					isMiddle = true
// 					break
// 				}

// 				if runes[l] == runes[currR] {
// 					runes[currR], runes[currR+1] = runes[currR+1], runes[currR]
// 					currR++
// 					moves++
// 				}
// 			}
// 		}
// 		l++
// 		if !isMiddle {
// 			r--
// 		}
// 	}
// 	return moves
// }

func minMovesToMakePalindrome(s string) int {
	runes := []rune(s)

	moves := 0

	for l, r := 0, len(runes)-1; l < r; l++ {
		currR := r
		for currR > l {
			if runes[l] == runes[currR] {
				for currR < r {
					runes[currR], runes[currR+1] = runes[currR+1], runes[currR]
					moves++
					currR++
				}
				r--
				break
			}
			currR--
		}
		if currR == l {
			moves += len(runes)/2 - l
		}
	}
	return moves
}
