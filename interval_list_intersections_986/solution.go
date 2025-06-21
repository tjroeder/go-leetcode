package intervallistintersections

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	res := [][]int{}
	i, j := 0, 0
	for i < len(firstList) && j < len(secondList) {
		maxStarting := max(firstList[i][0], secondList[j][0])
		minEnding := min(firstList[i][1], secondList[j][1])

		if maxStarting <= minEnding {
			res = append(res, []int{maxStarting, minEnding})
		}

		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}
	return res
}
