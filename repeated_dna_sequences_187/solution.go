package repeateddnasequences

// with simple map
func findRepeatedDNASequences(s string) []string {
	found := make([]string, 0)
	seen := make(map[string]int)

	for i := 0; i <= len(s)-10; i++ {
		subStr := s[i : i+10]
		if seen[subStr] == 1 {
			found = append(found, subStr)
		}
		seen[subStr]++
	}
	return found
}
