package knapsack

// space optimized bottom up technique
func findMaxKnapsackProfit(capacity int, weights []int, values []int) int {
	n := len(weights)

	// previous (i-1)th row which will be used to fill up the current ith row
	dp := make([]int, capacity+1)

	// current ith row that will use the values of the previous (i-1)th row to fill itself.
	temp := make([]int, capacity+1)

	for i := 1; i <= n; i++ {
		for j := 1; j <= capacity; j++ {
			if weights[i-1] <= j {
				temp[j] = max(values[i-1]+dp[j-weights[i-1]], dp[j])
			} else {
				temp[j] = dp[j]
			}
		}
		copy(dp, temp)
	}

	return dp[capacity]
}

// bottom up technique
// func findMaxKnapsackProfit(capacity int, weights []int, values []int) int {
// 	// Create a table to hold intermediate values
// 	n := len(weights)
// 	dp := make([][]int, n+1)
// 	for i := range dp {
// 		dp[i] = make([]int, capacity+1)
// 	}

// 	for i := 1; i <= n; i++ {
// 		for j := 1; j <= capacity; j++ {
// 			// Check if the weight of the current item is less than the current capacity
// 			if weights[i-1] <= j {
// 				dp[i][j] = max(values[i-1]+dp[i-1][j-weights[i-1]], dp[i-1][j])
// 			} else {
// 				dp[i][j] = dp[i-1][j]
// 			}
// 		}
// 	}

// 	return dp[n][capacity]
// }

// top down memoization
// func findMaxKnapsackProfitHelper(capacity int, weights []int, values []int, n int, dp [][]int) int {
// 	// Base case
// 	if n == 0 || capacity == 0 {
// 		return 0
// 	}

// 	// If we have already solved this subproblem, fetch the result from memory
// 	if dp[n][capacity] != -1 {
// 		return dp[n][capacity]
// 	}

// 	// Otherwise, we solve it and save the result in our lookup table
// 	if weights[n-1] <= capacity {
// 		dp[n][capacity] = max(values[n-1]+findMaxKnapsackProfitHelper(capacity-weights[n-1], weights, values, n-1, dp),
// 			findMaxKnapsackProfitHelper(capacity, weights, values, n-1, dp))
// 		return dp[n][capacity]
// 	}

// 	dp[n][capacity] = findMaxKnapsackProfitHelper(capacity, weights, values, n-1, dp)
// 	return dp[n][capacity]
// }

// // top down memoization
// func findMaxKnapsackProfit(capacity int, weights []int, values []int) int {
// 	n := len(weights)
// 	// Set up the dp table to store solutions to subproblems
// 	dp := make([][]int, n+1)
// 	for i := range dp {
// 		dp[i] = make([]int, capacity+1)
// 		for j := range dp[i] {
// 			dp[i][j] = -1
// 		}
// 	}
// 	return findMaxKnapsackProfitHelper(capacity, weights, values, n, dp)
// }
