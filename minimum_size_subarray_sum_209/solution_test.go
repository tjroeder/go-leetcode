package minimumsizesubarraysum

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMinSubArrayLen(t *testing.T) {
	testCases := []struct {
		description string
		givenTarget int
		givenNums   []int
		want        int
	}{
		{
			description: "base example [4,3] has minimum length ",
			givenTarget: 7,
			givenNums:   []int{2, 3, 1, 2, 4, 3},
			want:        2,
		},
		{
			description: "one element equals target",
			givenTarget: 4,
			givenNums:   []int{1, 4, 4},
			want:        1,
		},
		{
			description: "cannot equal sum",
			givenTarget: 11,
			givenNums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			want:        0,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := minSubArrayLen(tc.givenTarget, tc.givenNums)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("minSubArrayLen() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
