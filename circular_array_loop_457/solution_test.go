package circulararrayloop

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCircularArrayLoop(t *testing.T) {
	testCases := []struct {
		description string
		given       []int
		want        bool
	}{
		{
			description: "is circular",
			given:       []int{2, -1, 1, 2, 2},
			want:        true,
		},
		{
			description: "is not circular",
			given:       []int{-1, -2, -3, -4, -5, 6},
			want:        false,
		},
		{
			description: "multiple loops which are circular",
			given:       []int{1, -1, 5, 1, 4},
			want:        true,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := circularArrayLoop(tc.given)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("circularArrayLoop() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
