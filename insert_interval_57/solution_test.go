package insertinterval

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInsert(t *testing.T) {
	testCases := []struct {
		description      string
		givenIntervals   [][]int
		givenNewInterval []int
		want             [][]int
	}{
		{
			description:      "no existing interval, only newInterval",
			givenIntervals:   [][]int{},
			givenNewInterval: []int{2, 5},
			want:             [][]int{{2, 5}},
		},
		{
			description:      "only one interval is updated",
			givenIntervals:   [][]int{{1, 3}, {6, 9}},
			givenNewInterval: []int{2, 5},
			want:             [][]int{{1, 5}, {6, 9}},
		},
		{
			description:      "multiple intervals are updated, new interval [4,8] overlaps with [3,5],[6,7],[8,10]",
			givenIntervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			givenNewInterval: []int{4, 8},
			want:             [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := insert(tc.givenIntervals, tc.givenNewInterval)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("insert() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
