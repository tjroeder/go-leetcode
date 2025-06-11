package maximumaveragesubarryi

func findMaxAverage(nums []int, k int) float64 {
	left, right, currSum := 0, 0, 0

	for right < k {
		currSum += nums[right]
		right++
	}
	maxSum := currSum

	for right < len(nums) {
		currSum += nums[right] - nums[left]
		maxSum = max(maxSum, currSum)
		right++
		left++
	}

	return float64(maxSum) / float64(k)
}
