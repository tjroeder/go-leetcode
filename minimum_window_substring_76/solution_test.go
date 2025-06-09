package minimumwindowsubstring

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMinWindow(t *testing.T) {
	testCases := []struct {
		description string
		givenS      string
		givenT      string
		want        string
	}{
		{
			description: "substring is the full string t",
			givenT:      "a",
			givenS:      "a",
			want:        "a",
		},
		{
			description: "substring with multiple matches",
			givenT:      "ADOBECODEBANC",
			givenS:      "ABC",
			want:        "BANC",
		},
		{
			description: "substring with duplicate no matches",
			givenT:      "a",
			givenS:      "aa",
			want:        "",
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := minWindow(tc.givenT, tc.givenS)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("minWindow() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
