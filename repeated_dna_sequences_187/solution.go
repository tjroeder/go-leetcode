package repeateddnasequences

// with simple map
// func findRepeatedDNASequences(s string) []string {
// 	found := make([]string, 0)
// 	seen := make(map[string]int)

// 	for i := 0; i <= len(s)-10; i++ {
// 		subStr := s[i : i+10]
// 		if seen[subStr] == 1 {
// 			found = append(found, subStr)
// 		}
// 		seen[subStr]++
// 	}
// 	return found
// }

// with rolling hash
func findRepeatedDNASequences(s string) []string {
	// Convert DNA string to numbers
	toInt := map[rune]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}
	encodedSequence := make([]int, len(s))

	for i, c := range s {
		encodedSequence[i] = toInt[c]
	}

	k, n := 10, len(s) // k is length of substring(sequence)

	// 	If the string is shorter than 10 characters, return an empty list
	if n <= k {
		return []string{}
	}

	a, h := 5, 0 // Base-4 encoding(5 is used since next prime number), Hash value
	a_k := 1     // Stores a^L for efficient rolling hash updates
	seenHashes := make(map[int]bool)
	output := make(map[string]bool)

	// Compute initial hash for the first 10-letter substring
	// hash = (c_1​​×a^(k−1​)) + (c_2​​×a^(k−2​)) + ...... + (c_k×a^0​)
	// h_0 ​= (c_1​×4^9) + (c_2​×4^8) + ... + (c_10​×4^0)
	for i := 0; i < k; i++ {
		h = h*a + encodedSequence[i]
		a_k *= a
	}

	// Store the initial hash
	seenHashes[h] = true

	// Sliding window approach to update the hash efficiently
	for start := 1; start <= n-k; start++ {

		// Remove the leftmost character and add the new rightmost character
		h = h*a - encodedSequence[start-1]*a_k + encodedSequence[start+k-1]

		// If this hash has been seen_hashes before, add the corresponding substring to the output
		// TODO: add a check to verify that actual string matches in case of suprious hash hit
		if seenHashes[h] {
			output[s[start:start+k]] = true
		} else {
			seenHashes[h] = true
		}
	}

	// Convert set to list before returning
	result := make([]string, 0, len(output))
	for seq := range output {
		result = append(result, seq)
	}

	return result
}
