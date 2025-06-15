package frequencyofthemostfrequentelement

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMaxFrequency(t *testing.T) {
	testCases := []struct {
		description string
		givenNums   []int
		givenK      int
		want        int
	}{
		{
			description: "increment all other elements to match",
			givenNums:   []int{1, 2, 4},
			givenK:      5,
			want:        3,
		},
		{
			description: "multiple optimal solutions",
			givenNums:   []int{1, 4, 8, 13},
			givenK:      5,
			want:        2,
		},
		{
			description: "all elements only have 1 freq",
			givenNums:   []int{3, 9, 6},
			givenK:      2,
			want:        1,
		},
		{
			description: "longer nums",
			givenNums:   []int{2, 2, 2, 3, 5, 7, 7, 8, 10},
			givenK:      3,
			want:        4,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := maxFrequency(tc.givenNums, tc.givenK)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("maxFrequency() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
