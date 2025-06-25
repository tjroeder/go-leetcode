package mergesortedarray

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMergeSortedArray(t *testing.T) {
	testCases := []struct {
		description string
		givenNums1  []int
		givenM      int
		givenNums2  []int
		givenN      int
		want        []int
	}{
		{
			description: "bigger nums1 array than nums2",
			givenNums1:  []int{1, 2, 3, 0, 0, 0},
			givenM:      3,
			givenNums2:  []int{2, 5, 6},
			givenN:      3,
			want:        []int{1, 2, 2, 3, 5, 6},
		},
		{
			description: "nums2 is empty",
			givenNums1:  []int{1},
			givenM:      1,
			givenNums2:  []int{},
			givenN:      0,
			want:        []int{1},
		},
		{
			description: "nums1 is empty",
			givenNums1:  []int{0},
			givenM:      0,
			givenNums2:  []int{1},
			givenN:      1,
			want:        []int{1},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			merge(tc.givenNums1, tc.givenM, tc.givenNums2, tc.givenN)
			if diff := cmp.Diff(tc.want, tc.givenNums1); diff != "" {
				t.Errorf("merge() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
