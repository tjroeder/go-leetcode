package slidingwindowmaximum

// super naive approach
// func maxSlidingWindow(nums []int, k int) []int {
// 	maxValues := []int{}

// 	for i := 0; i < len(nums)-k+1; i++ {
// 		maxValues = append(maxValues, slices.Max(nums[i:k+i]))
// 	}

// 	return maxValues
// }

// sliding window without deque
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 {
		return []int{}
	}

	maxValues := []int{}
	window := []int{}

	for i, num := range nums {
		// Remove indices from front if out of window
		for len(window) > 0 && window[0] <= i-k {
			window = window[1:]
		}

		// Remove indices from back if current element is greater
		for len(window) > 0 && nums[window[len(window)-1]] <= num {
			window = window[:len(window)-1]
		}
		window = append(window, i)

		// Add to maxValues if window is full
		if i >= k-1 {
			maxValues = append(maxValues, nums[window[0]])
		}
	}
	return maxValues
}

// approach with sliding window deque
// func maxSlidingWindow(nums []int, k int) []int {
// 	n := len(nums)
// 	if n == 1 {
// 		return nums
// 	}

// 	maxValues := make([]int, n-k+1)
// 	window := NewDeque()
// 	// iterate over the first k elements
// 	for i := 0; i < k; i++ {
// 		cleanUp(i, window, nums)
// 		window.PushBack(i)
// 	}

// 	// add the first max value from the initial k elements
// 	maxValues[0] = nums[window.Front()]
// 	for i := k; i < n; i++ {
// 		cleanUp(i, window, nums)
// 		// check if the first element in window is even within our current k window
// 		if window.Len() > 0 && window.Front() <= i-k {
// 			window.PopFront()
// 		}
// 		window.PushBack(i)
// 		maxValues[i-k+1] = nums[window.Front()]
// 	}
// 	return maxValues
// }

// // function to clean up the window
// func cleanUp(i int, window *Deque, nums []int) {
// 	// removes all values in window which are less than the nums[i]
// 	for window.Len() > 0 && nums[i] >= nums[window.Back()] {
// 		window.PopBack()
// 	}
// }
