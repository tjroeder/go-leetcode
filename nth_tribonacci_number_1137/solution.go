package nthtribonaccinumber

// tribonacci follows the sequence
// T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0
func tribonacci(n int) int {
	found := make(map[int]int)
	// base conditions
	found[0] = 0
	found[1] = 1
	found[2] = 1

	return tribonacciHelper(n, found)
}

// tribonacciHelper checks the map for any already calculated values, if not already
// calculated, it will recursively call tribonacciHelper to calculate the value
func tribonacciHelper(n int, found map[int]int) int {
	if val, ok := found[n]; ok {
		return val
	}

	found[n] = tribonacciHelper(n-1, found) + tribonacciHelper(n-2, found) + tribonacciHelper(n-3, found)
	return found[n]
}
