package circulararrayloop

func circularArrayLoop(nums []int) bool {
	hm := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if hm[i] {
			continue
		}
		slow, fast := i, i
		isForward := nums[i] > 0

		for {
			slow = next(nums, slow)
			if isNotCycle(nums, isForward, slow) {
				break
			}

			// first movement of fast
			fast = next(nums, fast)
			if isNotCycle(nums, isForward, fast) {
				break
			}

			// second movement of fast
			fast = next(nums, fast)
			if isNotCycle(nums, isForward, fast) {
				break
			}
			hm[slow] = true
			hm[fast] = true

			if slow == fast {
				return true
			}
		}
	}

	return false
}

// isCycle checks that the direction is the same during traversal of the array,
// and that the value doesn't back to the same element
func isNotCycle(nums []int, prevDirection bool, pointer int) bool {
	currDirection := nums[pointer] >= 0

	if (prevDirection != currDirection) || (nums[pointer]%len(nums) == 0) {
		return true
	}
	return false
}

// next takes the array and index pointer, and moves the pointer to the next
// index pointer based on the index value and it's current index
func next(nums []int, pointer int) int {
	pointer = (pointer + nums[pointer]) % len(nums)
	if pointer < 0 {
		pointer += len(nums)
	}
	return pointer
}
