package slidingwindowmaximum

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMaxSlidingWindow(t *testing.T) {
	testCases := []struct {
		description string
		givenNums   []int
		givenK      int
		want        []int
	}{
		{
			description: "single sliding window max",
			givenNums:   []int{1},
			givenK:      1,
			want:        []int{1},
		},
		{
			description: "multiple max sliding window values",
			givenNums:   []int{1, 3, -1, -3, 5, 3, 6, 7},
			givenK:      3,
			want:        []int{3, 3, 5, 5, 6, 7},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := maxSlidingWindow(tc.givenNums, tc.givenK)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("maxSlidingWindow() mismatch (-want +got):\n%s", diff)
			}

		})
	}
}
