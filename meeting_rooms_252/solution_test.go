package meetingrooms

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCanAttendMeetings(t *testing.T) {
	testCases := []struct {
		description    string
		givenIntervals [][]int
		want           bool
	}{
		{
			description:    "intervals overlap",
			givenIntervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			want:           false,
		},
		{
			description:    "intervals do not overlap, even with meeting starting and ending at same time",
			givenIntervals: [][]int{{7, 10}, {2, 7}},
			want:           true,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := canAttendMeetings(tc.givenIntervals)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("canAttendMeetings() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
