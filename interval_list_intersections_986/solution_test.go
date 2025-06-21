package intervallistintersections

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIntervalIntersection(t *testing.T) {
	testCases := []struct {
		description     string
		givenFirstList  [][]int
		givenSecondList [][]int
		want            [][]int
	}{
		{
			description:     "intervals overlap",
			givenFirstList:  [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
			givenSecondList: [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			want:            [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		},
		{
			description:     "intervals do not overlap, no second list",
			givenFirstList:  [][]int{{1, 3}, {5, 9}},
			givenSecondList: [][]int{},
			want:            [][]int{},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := intervalIntersection(tc.givenFirstList, tc.givenSecondList)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("intervalIntersection() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
