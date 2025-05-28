package findtheduplicatenumber

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindDuplicate(t *testing.T) {
	testCases := []struct {
		description string
		given       []int
		want        int
	}{
		{
			description: "found duplicate",
			given:       []int{1, 3, 4, 2, 2},
			want:        2,
		},
		{
			description: "another found duplicate, but at the beginning",
			given:       []int{3, 1, 3, 4, 2},
			want:        3,
		},
		{
			description: "duplicate with all numbers the same",
			given:       []int{3, 3, 3, 3, 3},
			want:        3,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := findDuplicate(tc.given)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("findDuplicate() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
