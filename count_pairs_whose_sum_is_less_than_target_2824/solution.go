package countPairs

import "sort"

// initial without sorting
// func countPairs(nums []int, target int) int {
// 	count := 0

// 	for l := 0; l < len(nums)-1; l++ {
// 		r := len(nums) - 1
// 		for l < r {
// 			if nums[l]+nums[r] < target {
// 				count++
// 			}
// 			r--
// 		}
// 	}
// 	return count
// }

// with sorting
func countPairs(nums []int, target int) int {
	sort.Ints(nums)
	count := 0
	l, r := 0, len(nums)-1

	for l < r {
		if nums[l]+nums[r] < target {
			count += r - l
			l++
		} else {
			r--
		}
	}

	return count
}
