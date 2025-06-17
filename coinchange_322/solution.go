package coinchange

import "math"

func coinChange(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}
	counter := make([]int, amount)
	for index := range counter {
		counter[index] = math.MaxInt32
	}
	return calculateMinimumCoins(coins, amount, counter)
}

// Helper function that calculates amount left to be calculated and tells what it's value can be
func calculateMinimumCoins(coins []int, remainingAmount int, counter []int) int {
	if remainingAmount < 0 {
		return -1
	}
	if remainingAmount == 0 {
		return 0
	}
	if counter[remainingAmount-1] != math.MaxInt32 {
		return counter[remainingAmount-1]
	}
	minimum := math.MaxInt32

	// Recursive approach to keep in account every number's result
	for _, value := range coins {
		result := calculateMinimumCoins(coins, remainingAmount-value, counter)
		if result >= 0 && result < minimum {
			minimum = 1 + result
		}
	}

	if minimum != math.MaxInt32 {
		counter[remainingAmount-1] = minimum
	} else {
		counter[remainingAmount-1] = -1
	}

	return counter[remainingAmount-1]
}
