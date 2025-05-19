package threesum

import "slices"

// func threeSum(nums []int) [][]int {
// 	result := make([][]int, 0)
// 	if len(nums) < 3 {
// 		return result
// 	}

// 	slices.Sort(nums)
// 	if len(nums) == 3 && nums[0]+nums[1]+nums[2] == 0 {
// 		return append(result, nums)
// 	}

// 	for i, num := range nums {
// 		if num > 0 {
// 			break
// 		}
// 		if i > 0 && num == nums[i-1] {
// 			continue
// 		}
// 		start := i + 1
// 		end := len(nums) - 1
// 		for start < end {
// 			sum := num + nums[start] + nums[end]
// 			if sum < 0 {
// 				start++
// 				continue
// 			}
// 			if sum > 0 {
// 				end--
// 				continue
// 			}
// 			result = append(result, []int{num, nums[start], nums[end]})
// 			for start < end && nums[start] == nums[start+1] {
// 				start++
// 			}
// 			for start < end && nums[end] == nums[end-1] {
// 				end--
// 			}
// 			start++
// 			end--
// 		}
// 	}
// 	return result
// }

func threeSum(nums []int) [][]int {
	triplets := make([][]int, 0)
	slices.Sort(nums)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, h := i+1, len(nums)-1
		for l < h {
			sum := nums[i] + nums[l] + nums[h]
			if sum < 0 {
				l++
				continue
			}
			if sum > 0 {
				h--
				continue
			}
			triplets = append(triplets, []int{nums[i], nums[l], nums[h]})
			for l < h && nums[l] == nums[l+1] {
				l++
			}
			for l < h && nums[h] == nums[h-1] {
				h--
			}
			l++
			h--
		}
	}
	return triplets
}
