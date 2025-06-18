package insertinterval

import "slices"

// non-optimized approach which just appends the new interval and merges
// o(nlogn) due to the sort
func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	if len(intervals) == 0 {
		return append(res, newInterval)
	}

	// rather than try and insert the new interval dynamically, just append, then
	// do the same thing as merge interval by sorting than merge the intervals so there
	// are no overlaps
	intervals = append(intervals, newInterval)
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= res[len(res)-1][1] {
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}
