package findtheduplicatenumber

// findDuplicate with a status variable
// func findDuplicate(nums []int) int {
// 	met := false
// 	slow, fast := 0, 0

// 	for {
// 		if !met {
// 			slow = nums[slow]
// 			fast = nums[nums[fast]]
// 		} else {
// 			slow = nums[slow]
// 			fast = nums[fast]
// 		}

// 		if slow == fast && !met {
// 			met = true
// 			slow = 0
// 		}

// 		if slow == fast && met {
// 			break
// 		}

// 	}
// 	return fast
// }

// findDuplicate without a status variable
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}
	}

	slow = 0
	for {
		slow = nums[slow]
		fast = nums[fast]

		if slow == fast {
			break
		}
	}
	return fast
}
