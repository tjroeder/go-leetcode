package threesum

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestThreeSum(t *testing.T) {
	testCases := []struct {
		description string
		given       []int
		want        [][]int
	}{
		{
			description: "single return test",
			given:       []int{0, 0, 0},
			want:        [][]int{{0, 0, 0}},
		},
		{
			description: "multi return test",
			given:       []int{-3, -2, 0, 1, 2, 3},
			want:        [][]int{{-3, 0, 3}, {-3, 1, 2}, {-2, 0, 2}},
		},
		{
			description: "all positive numbers test",
			given:       []int{1, 2, 3, 4, 5, 6},
			want:        [][]int{},
		},
		{
			description: "large test",
			given:       []int{2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4, 5, 5, -10},
			want:        [][]int{{-10, 5, 5}, {-5, 0, 5}, {-4, 2, 2}, {-3, -2, 5}, {-3, 1, 2}, {-2, 0, 2}},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := threeSum(tc.given)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("threeSum() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
