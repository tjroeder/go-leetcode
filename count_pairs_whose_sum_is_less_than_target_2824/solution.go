package countPairs

func countPairs(nums []int, target int) int {
	count := 0

	for l := 0; l < len(nums)-1; l++ {
		r := len(nums) - 1
		for l < r {
			if nums[l]+nums[r] < target {
				count++
			}
			r--
		}
	}
	return count
}
