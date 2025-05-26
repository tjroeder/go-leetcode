package happynumber

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsHappy(t *testing.T) {
	testCases := []struct {
		description string
		given       int
		want        bool
	}{
		{
			description: "happy test",
			given:       19,
			want:        true,
		},
		{
			description: "happy test with 1",
			given:       1,
			want:        true,
		},
		{
			description: "not happy test",
			given:       2,
			want:        false,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := isHappy(tc.given)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("isHappy() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
