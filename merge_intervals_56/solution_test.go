package mergeintervals

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMerge(t *testing.T) {
	testCases := []struct {
		description string
		given       [][]int
		want        [][]int
	}{
		{
			description: "only one interval, no merge",
			given:       [][]int{{14, 20}},
			want:        [][]int{{14, 20}},
		},
		{
			description: "single overlap merge",
			given:       [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:        [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			description: "single overlap merge with only two intervals",
			given:       [][]int{{1, 4}, {2, 3}},
			want:        [][]int{{1, 4}},
		},
		{
			description: "merge all intervals into one interval",
			given:       [][]int{{1, 5}, {4, 6}, {3, 7}, {6, 8}},
			want:        [][]int{{1, 8}},
		},
		{
			description: "no overlap, no merge",
			given:       [][]int{{1, 4}, {6, 7}, {9, 10}},
			want:        [][]int{{1, 4}, {6, 7}, {9, 10}},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := merge(tc.given)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("merge() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
