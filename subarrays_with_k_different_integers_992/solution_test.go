package subarrayswithkdifferentintegers

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSubarraysWithKDistinct(t *testing.T) {
	testCases := []struct {
		description string
		givenNums   []int
		givenK      int
		want        int
	}{
		{
			description: "subarray matches array",
			givenNums:   []int{1},
			givenK:      1,
			want:        1,
		},
		{
			description: "one subarray with only 3 distinct integers",
			givenNums:   []int{1, 2, 3, 1, 2},
			givenK:      3,
			want:        6,
		},
		{
			description: "multiple subarrays, [1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2]",
			givenNums:   []int{1, 2, 1, 2, 3},
			givenK:      2,
			want:        7,
		},
		{
			description: "multiple subarrays, [1,2,1,3], [2,1,3], [1,3,4]",
			givenNums:   []int{1, 2, 1, 3, 4},
			givenK:      3,
			want:        3,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := subarraysWithKDistinct(tc.givenNums, tc.givenK)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("subarraysWithKDistinct() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
