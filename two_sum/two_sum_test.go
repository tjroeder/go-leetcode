package twosum

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTwoSums(t *testing.T) {
	testCases := []struct {
		description string
		nums        []int
		target      int
		want        []int
	}{
		{
			description: "less than two items in nums",
			nums:        []int{2},
			target:      9,
			want:        []int{},
		},
		{
			description: "only two items in nums",
			nums:        []int{2, 7},
			target:      9,
			want:        []int{0, 1},
		},
		{
			description: "base test",
			nums:        []int{2, 7, 11, 15},
			target:      9,
			want:        []int{0, 1},
		},
		{
			description: "another test",
			nums:        []int{2, 7, 11, 15},
			target:      26,
			want:        []int{2, 3},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			got := twosum(tc.nums, tc.target)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("twosums() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
