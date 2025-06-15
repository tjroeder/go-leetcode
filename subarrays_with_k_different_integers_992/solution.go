package subarrayswithkdifferentintegers

// with counting distinct ints rather than len(freq)
func subarraysWithAtMostKDistinct(nums []int, k int) int {
	freq := make(map[int]int)
	l, res := 0, 0

	for r, num := range nums {
		freq[num]++
		if freq[num] == 1 {
			k--
		}

		for k < 0 {
			freq[nums[l]]--
			if freq[nums[l]] == 0 {
				k++
			}
			l++
		}
		res += r - l + 1
	}
	return res
}

// without counting the distinct integers, instead relys on
// length of freq map, and deleting keys if 0
// func subarraysWithAtMostKDistinct(nums []int, k int) int {
// 	freq := make(map[int]int)
// 	l, res := 0, 0

// 	for r, num := range nums {
// 		freq[num]++

// 		for len(freq) > k {
// 			freq[nums[l]]--
// 			if freq[nums[l]] == 0 {
// 				delete(freq, nums[l])
// 			}
// 			l++
// 		}
// 		res += r - l + 1
// 	}
// 	return res
// }

// to find the kDistinct would take O(n^2) if we don't calculate
// the atMost(k) - atMost(k-1) subarrays.
// E.g. for array [1, 2, 3, 1, 2], k = 3
// <= k integers     = [1], [2], [3], [1], [2], [1, 2], [2, 3], [3, 1], [1, 2], [1, 2, 3], [2, 3, 1], [3, 1, 2], [1, 2, 3, 1], [2, 3, 1, 2], [1, 2, 3, 1, 2]
// <= k - 1 integers = [1], [2], [3], [1], [2], [1, 2], [2, 3], [3, 1], [1, 2]
// so subtracting atMost(k) - atMost(k-1) = [1, 2, 3], [2, 3, 1], [3, 1, 2], [1, 2, 3, 1], [2, 3, 1, 2], [1, 2, 3, 1, 2] = 6
func subarraysWithKDistinct(nums []int, k int) int {
	return subarraysWithAtMostKDistinct(nums, k) - subarraysWithAtMostKDistinct(nums, k-1)
}
