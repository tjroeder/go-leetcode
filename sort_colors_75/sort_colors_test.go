package sortcolors

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSortColors(t *testing.T) {
	testCases := []struct {
		description string
		given       []int
		want        []int
	}{
		{
			description: "only 1 color",
			given:       []int{0},
			want:        []int{0},
		},
		{
			description: "only 1 of each color",
			given:       []int{2, 0, 1},
			want:        []int{0, 1, 2},
		},
		{
			description: "more complex sort",
			given:       []int{1, 0, 2, 1, 2, 2},
			want:        []int{0, 1, 1, 2, 2, 2},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			got := sortColors(tc.given)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("sortColors() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
