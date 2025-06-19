package insertinterval

// O(n) with insert in place
func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}

	i := 0
	// add all the intervals before newInterval
	for i < len(intervals) && intervals[i][0] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}

	// append or merge newInterval depending on if there is an overlap
	// last element of res's ending interval <  newIntervals beginning interval
	// we append newInterval without issue
	if len(res) == 0 || res[len(res)-1][1] < newInterval[0] {
		res = append(res, newInterval)
	} else { // if last element of res's ending interval >= newIntervals beginning interval
		res[len(res)-1][1] = max(res[len(res)-1][1], newInterval[1]) // pick the largest last interval
	}

	// now we need to merge any potential intervals we could have broken during our insert, or we
	// just append the remaining intervals
	for i < len(intervals) {
		if res[len(res)-1][1] >= intervals[i][0] { // last element of res ending interval >= current intervals beginning interval
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1]) // pick the largest last interval
		} else {
			res = append(res, intervals[i:]...)
			break
		}
		i++
	}
	return res
}

// non-optimized approach which just appends the new interval and merges
// o(nlogn) due to the sort
// func insert(intervals [][]int, newInterval []int) [][]int {
// 	res := [][]int{}
// 	if len(intervals) == 0 {
// 		return append(res, newInterval)
// 	}

// 	// rather than try and insert the new interval dynamically, just append, then
// 	// do the same thing as merge interval by sorting than merge the intervals so there
// 	// are no overlaps
// 	intervals = append(intervals, newInterval)
// 	slices.SortFunc(intervals, func(a, b []int) int {
// 		return a[0] - b[0]
// 	})

// 	res = append(res, intervals[0])

// 	for i := 1; i < len(intervals); i++ {
// 		if intervals[i][0] <= res[len(res)-1][1] {
// 			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
// 		} else {
// 			res = append(res, intervals[i])
// 		}
// 	}
// 	return res
// }

// leetcode solution, with minor refactor to remove math import and float64/int conversions
// also removed the last for loop to use append with remaining elements expanded
// func insert(intervals [][]int, newInterval []int) [][]int {
// 	var result [][]int

// 	// Iterate through intervals and add non-overlapping intervals before newInterval
// 	i := 0
// 	for i < len(intervals) && intervals[i][1] < newInterval[0] {
// 		result = append(result, intervals[i])
// 		i++
// 	}

// 	// Merge overlapping intervals
// 	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
// 		newInterval[0] = min(newInterval[0], intervals[i][0])
// 		newInterval[1] = max(newInterval[1], intervals[i][1])
// 		i++
// 	}

// 	// Add merged newInterval
// 	result = append(result, newInterval)

// 	// Add non-overlapping intervals after newInterval
// 	// for i < len(intervals) {
// 	result = append(result, intervals[i:]...)

// 	return result
// }
