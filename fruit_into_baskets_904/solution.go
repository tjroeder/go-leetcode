package fruitintobaskets

func totalFruit(fruits []int) int {
	if len(fruits) == 1 {
		return 1
	}

	baskets := make(map[int]int)
	l, curr, res := 0, 0, 0

	for r := 0; r < len(fruits); r++ {
		baskets[fruits[r]]++
		curr++

		for len(baskets) > 2 {
			if baskets[fruits[l]] == 1 {
				delete(baskets, fruits[l])
				curr--
				l++
				break
			}
			baskets[fruits[l]]--
			curr--
			l++
		}
		res = max(res, curr)
	}
	return res
}
