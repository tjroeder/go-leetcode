package frequencyofthemostfrequentelement

import "slices"

// this solution continually finds  max window, and then moves that
// window across nums, then using the window size returns the window size
// rather than continually storing the max value
func maxFrequency(nums []int, k int) int {
	slices.Sort(nums)

	l, total := 0, 0
	for r := 0; r < len(nums); r++ {
		total += nums[r]

		if total+k < nums[r]*(r-l+1) {
			total -= nums[l]
			l++
		}
	}
	return len(nums) - l
}

// this solution tracks the max freq elements and expands
// shrinks the window as it goes to find the next max
// func maxFrequency(nums []int, k int) int {
// 	slices.Sort(nums)
// 	l, total, res := 0, 0, 0

// 	for r := 0; r < len(nums); r++ {
// 		// total is the sum of elements in the current window
// 		total += nums[r]
// 		// if the number of elements in the window * rightmost element in the
// 		// sliding window - the total sum of elements in the current window is
// 		// greater than k it means we need to shrink the window because
// 		// we have run out of k operations to fix the elements to match the
// 		// rightmost element
// 		for (r-l+1)*nums[r]-total > k {
// 			// remove the leftmost value from our running sum before moving the
// 			// left bound
// 			total -= nums[l]
// 			l++
// 		}
// 		// keep track of our highest result as we go
// 		res = max(res, r-l+1)
// 	}
// 	return res
// }

// reverse sliding works but is slow due to reseting l, after moving r
// func maxFrequency(nums []int, k int) int {
// 	if len(nums) == 1 {
// 		return 1
// 	}
// 	slices.Sort(nums)

// 	r, res, ctv := len(nums)-1, 1, nums[len(nums)-1]
// 	currK := k
// 	for l := len(nums) - 1; l >= 0; l-- {
// 		if ctv <= nums[l]+currK {
// 			res = max(res, r-l+1)
// 			currK -= ctv - nums[l]
// 		} else {
// 			r--
// 			l = r
// 			ctv = nums[r]
// 			currK = k
// 		}
// 	}
// 	return res
// }
