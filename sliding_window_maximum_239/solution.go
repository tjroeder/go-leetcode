package slidingwindowmaximum

// super naive approach
// func maxSlidingWindow(nums []int, k int) []int {
// 	maxValues := []int{}

// 	for i := 0; i < len(nums)-k+1; i++ {
// 		maxValues = append(maxValues, slices.Max(nums[i:k+i]))
// 	}

// 	return maxValues
// }

// approach with sliding window
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 1 {
		return nums
	}

	maxValues := make([]int, n-k+1)
	currentWindow := NewDeque()
	// iterate over the first k elements
	for i := 0; i < k; i++ {
		cleanUp(i, currentWindow, nums)
		currentWindow.PushBack(i)
	}

	// add the first max value from the initial k elements
	maxValues[0] = nums[currentWindow.Front()]
	for i := k; i < n; i++ {
		cleanUp(i, currentWindow, nums)
		// check if the first element in currentWindow is even within our current k window
		if currentWindow.Len() > 0 && currentWindow.Front() <= i-k {
			currentWindow.PopFront()
		}
		currentWindow.PushBack(i)
		maxValues[i-k+1] = nums[currentWindow.Front()]
	}
	return maxValues
}

// function to clean up the window
func cleanUp(i int, currentWindow *Deque, nums []int) {
	// removes all values in currentWindow which are less than the nums[i]
	for currentWindow.Len() > 0 && nums[i] >= nums[currentWindow.Back()] {
		currentWindow.PopBack()
	}
}
