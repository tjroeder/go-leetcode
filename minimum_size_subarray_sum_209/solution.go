package minimumsizesubarraysum

func minSubArrayLen(target int, nums []int) int {
	// res(ult) must be greater than length of nums
	res := len(nums) + 1
	left, right, sum := 0, 0, 0

	// iterate through nums
	for right < len(nums) {
		// keep a running sum, which we update each iteration with the new right bound
		// pointer in nums
		sum += nums[right]

		// if we are greater than our target lets save our sub array length and see if we can make it better
		// by shrinking the window, when moving our left bound
		if sum >= target {
			// save our minimum window length
			res = min(res, right-left+1)

			// now continually shrink the window with the left bound
			for sum >= target {
				// save as we go to get the smallest window
				res = min(res, right-left+1)
				// remove the left bound element value from the running sum then move the left bound
				sum -= nums[left]
				left++
			}
		}
		// move the right bound to expand our window
		right++
	}

	// if at the end we didn't update res than we found no matching subarray for our target, return 0
	if res > len(nums) {
		return 0
	}
	return res
}
