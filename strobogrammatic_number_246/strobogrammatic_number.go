package strobogrammaticnumber

func isStrobogrammatic(num string) bool {
	stroboNum := map[string]string{
		"0": "0",
		"1": "1",
		"6": "9",
		"8": "8",
		"9": "6",
	}

	l, h := 0, len(num)-1

	for l <= h {
		lNum, ok := stroboNum[string(num[l])]
		if !ok {
			return false
		}

		hNum, ok := stroboNum[string(num[h])]
		if !ok {
			return false
		}

		if lNum != string(num[h]) && hNum != string(num[l]) {
			return false
		}
		l++
		h--
	}

	return true
}
