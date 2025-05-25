package countPairs

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCountPairs(t *testing.T) {
	testCases := []struct {
		description string
		givenNums   []int
		givenTarget int
		want        int
	}{
		{
			description: "basic test",
			givenNums:   []int{-1, 1, 2, 3, 1},
			givenTarget: 2,
			want:        3,
		},
		{
			description: "another test",
			givenNums:   []int{5, 1, 3, 2},
			givenTarget: 7,
			want:        4,
		},
		{
			description: "no pairs",
			givenNums:   []int{-1, 9, 17, 11, 6},
			givenTarget: 3,
			want:        0,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := countPairs(tc.givenNums, tc.givenTarget)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("countPairs() mismatch (-want +got):\n%s", diff)
			}

		})
	}
}
