package ipo

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindMaximizedCapital(t *testing.T) {
	testCases := []struct {
		description   string
		givenK        int
		givenW        int
		givenProfits  []int
		givenCapitals []int
		want          int
	}{
		{
			description:   "zero initial capital",
			givenK:        2,
			givenW:        0,
			givenProfits:  []int{1, 2, 3},
			givenCapitals: []int{0, 1, 1},
			want:          4,
		},
		{
			description:   "zero initial capital select all projects",
			givenK:        3,
			givenW:        0,
			givenProfits:  []int{1, 2, 3},
			givenCapitals: []int{0, 1, 2},
			want:          6,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := findMaximizedCapital(tc.givenK, tc.givenW, tc.givenProfits, tc.givenCapitals)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("findMaximizedCapital() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
