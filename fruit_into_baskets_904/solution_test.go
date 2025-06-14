package fruitintobaskets

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTotalFruit(t *testing.T) {
	testCases := []struct {
		description string
		givenFruits []int
		want        int
	}{
		{
			description: "able to pick all 3 trees",
			givenFruits: []int{1, 2, 1},
			want:        3,
		},
		{
			description: "is not best to pick the first tree [1, 2, 2]",
			givenFruits: []int{0, 1, 2, 2},
			want:        3,
		},
		{
			description: "longer first not optimal [2, 3, 2, 2]",
			givenFruits: []int{1, 2, 3, 2, 2},
			want:        4,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := totalFruit(tc.givenFruits)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("totalFruit() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
