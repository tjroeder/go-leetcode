package containsduplicateii

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	testCases := []struct {
		description string
		givenNums   []int
		givenK      int
		want        bool
	}{
		{
			description: "first and last indice values match, and abs(i-j) <= k",
			givenNums:   []int{1, 2, 3, 1},
			givenK:      3,
			want:        true,
		},
		{
			description: "multiple matching values",
			givenNums:   []int{1, 0, 1, 1},
			givenK:      1,
			want:        true,
		},
		{
			description: "longer first not optimal [2, 3, 2, 2]",
			givenNums:   []int{1, 2, 3, 1, 2, 3},
			givenK:      2,
			want:        false,
		},
		{
			description: "no duplicate in a single element nums array",
			givenNums:   []int{900},
			givenK:      900,
			want:        false,
		},
		{
			description: "two elements, which are the same",
			givenNums:   []int{99, 99},
			givenK:      2,
			want:        true,
		},
		{
			description: "k = 0, false",
			givenNums:   []int{1, 2, 1},
			givenK:      0,
			want:        false,
		},
		{
			description: "all increasing element values",
			givenNums:   []int{0, 1, 2, 3, 2, 5},
			givenK:      3,
			want:        true,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := containsNearbyDuplicate(tc.givenNums, tc.givenK)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("containsNearbyDuplicate() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
