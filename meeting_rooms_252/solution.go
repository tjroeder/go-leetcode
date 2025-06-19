package meetingrooms

import "slices"

func canAttendMeetings(intervals [][]int) bool {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	for i, interval := range intervals {
		if i < len(intervals)-1 && interval[1] > intervals[i+1][0] {
			return false
		}
	}
	return true
}
