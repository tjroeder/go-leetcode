package twosum

func twosum(nums []int, target int) []int {
	length := len(nums)
	if length < 2 {
		return []int{}
	}

	if length == 2 {
		return []int{0, 1}
	}

	numIndex := make(map[int]int)
	for i, num := range nums {
		search := target - num
		idx, ok := numIndex[search]
		if ok {
			return []int{idx, i}
		}
		numIndex[num] = i
	}
	return []int{}
}
