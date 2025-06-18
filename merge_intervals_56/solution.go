package mergeintervals

import "slices"

func merge(intervals [][]int) [][]int {
	// sort the intervals by start time with no tie-breaker
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	// sort the intervals by start time with tie-breaker (this is slower)
	// slices.SortFunc(intervals, func(a, b []int) int {
	// 	if a[0] != b[0] {
	// 		return a[0] - b[0]
	// 	}
	// 	return a[1] - b[1]
	// })

	// create a result 2d slice, and populate first element with the first element
	// in the sorted intervals slice
	res := [][]int{intervals[0]}
	// iterate over intervals
	for i := 1; i < len(intervals); i++ {
		// if the last value in the previous result interval slice [len(res)-1][1] is greater
		// than or equal to the first value in the current intervals slice [i][0]
		if res[len(res)-1][1] >= intervals[i][0] {
			// if the last value of the current intervals slice [i][1] is less than
			// or equal to the last value of the previous result interval slice [len(res)-1][1]
			if res[len(res)-1][1] <= intervals[i][1] {
				res[len(res)-1][1] = intervals[i][1]
			}
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}
