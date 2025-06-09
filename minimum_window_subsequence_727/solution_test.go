package minimumwindowsubsequence

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMinWindow(t *testing.T) {
	testCases := []struct {
		description string
		givenS1     string
		givenS2     string
		want        string
	}{
		{
			description: "substring is the full string t",
			givenS1:     "thims",
			givenS2:     "this",
			want:        "thims",
		},
		{
			description: "substring found",
			givenS1:     "asbfgedasfbdaaf",
			givenS2:     "bfd",
			want:        "bfged",
		},
		{
			description: "substring not found",
			givenS1:     "jmeqksfrsdcmsiwvaovztaqenprpvnbstl",
			givenS2:     "u",
			want:        "",
		},
		{
			description: "substring with duplicate letters, and last match is",
			givenS1:     "abcdebdde",
			givenS2:     "bde",
			want:        "bcde",
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := minWindow(tc.givenS1, tc.givenS2)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("minWindow() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
