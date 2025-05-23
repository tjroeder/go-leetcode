package validwordabbreviation

import "strconv"

func validWordAbbreviation(word string, abbr string) bool {
	w := len(word)
	a := len(abbr)

	wIdx := 0
	aIdx := 0

	for wIdx < w && aIdx < a {
		if word[wIdx] == abbr[aIdx] {
			wIdx++
			aIdx++
			continue
		}

		if abbr[aIdx] <= '0' || abbr[aIdx] > '9' {
			return false
		}

		num := ""
		for aIdx < a && abbr[aIdx] >= '0' && abbr[aIdx] <= '9' {
			num += string(abbr[aIdx])
			aIdx++
		}

		count, _ := strconv.Atoi(num)
		wIdx += count
	}

	return w == wIdx && a == aIdx
}
