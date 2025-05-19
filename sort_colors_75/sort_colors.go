package sortcolors

func sortColors(colors []int) []int {
	// red == 0, white == 1, blue == 2
	for start, current, end := 0, 0, len(colors)-1; current <= end; {
		if colors[current] == 0 {
			colors[start], colors[current] = colors[current], colors[start]
			start++
			current++
			continue
		}
		if colors[current] == 1 {
			current++
			continue
		}
		if colors[current] == 2 {
			colors[end], colors[current] = colors[current], colors[end]
			end--
		}
	}
	return colors
}
