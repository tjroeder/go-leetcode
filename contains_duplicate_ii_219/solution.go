package containsduplicateii

// with index hashmap
func containsNearbyDuplicate(nums []int, k int) bool {
	freq := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := freq[nums[i]]; !ok {
			freq[nums[i]] = i
		} else {
			if i-freq[nums[i]] <= k {
				return true
			} else {
				freq[nums[i]] = i
			}
		}
	}
	return false
}

// with seen hash map
// func containsNearbyDuplicate(nums []int, k int) bool {
// 	seen := make(map[int]bool)
// 	for r := 0; r < len(nums); r++ {
// 		if seen[nums[r]] {
// 			return true
// 		}
// 		seen[nums[r]] = true

// 		if len(seen) > k {
// 			delete(seen, nums[r-k])
// 		}
// 	}
// 	return false
// }
