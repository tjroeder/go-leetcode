package maximumaveragesubarryi

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindMaxAverage(t *testing.T) {
	testCases := []struct {
		description string
		givenNums   []int
		givenK      int
		want        float64
	}{
		{
			description: "base example [12, -5, -6, 50]",
			givenNums:   []int{1, 12, -5, -6, 50, 3},
			givenK:      4,
			want:        12.75000,
		},
		{
			description: "one value in array",
			givenNums:   []int{5},
			givenK:      1,
			want:        5.00000,
		},
		{
			description: "target k is 1",
			givenNums:   []int{0, 4, 0, 3, 2},
			givenK:      1,
			want:        4.00000,
		},
		{
			description: "another test",
			givenNums:   []int{4, 2, 1, 3, 3},
			givenK:      2,
			want:        3.00000,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := findMaxAverage(tc.givenNums, tc.givenK)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("findMaxAverage() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
