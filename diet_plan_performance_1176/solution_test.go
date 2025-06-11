package dietplanperformance

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDietPlanPerformance(t *testing.T) {
	testCases := []struct {
		description   string
		givenCalories []int
		givenK        int
		givenLower    int
		givenUpper    int
		want          int
	}{
		{
			description:   "k = 1, considers all elements individually that match exactly the same as the upper and lower",
			givenCalories: []int{1, 2, 3, 4, 5},
			givenK:        1,
			givenLower:    3,
			givenUpper:    3,
			want:          0,
		},
		{
			description:   "array length is the same amount as k",
			givenCalories: []int{3, 2},
			givenK:        2,
			givenLower:    0,
			givenUpper:    1,
			want:          1,
		},
		{
			description:   "array length is the same amount as k",
			givenCalories: []int{3, 2},
			givenK:        2,
			givenLower:    0,
			givenUpper:    1,
			want:          1,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := dietPlanPerformance(tc.givenCalories, tc.givenK, tc.givenLower, tc.givenUpper)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("dietPlanPerformance() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
