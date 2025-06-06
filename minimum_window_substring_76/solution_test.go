package minimumwindowsubstring

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMinWindowEducative(t *testing.T) {
	testCases := []struct {
		description string
		givenS      string
		givenT      string
		want        string
	}{
		{
			description: "substring is the full string t",
			givenT:      "thims",
			givenS:      "this",
			want:        "thims",
		},
		{
			description: "substring in a sentence",
			givenT:      "this is a test string",
			givenS:      "ths",
			want:        "this",
		},
		{
			description: "substring in a jumble of letters",
			givenT:      "asbfgedasfbdaaf",
			givenS:      "bfd",
			want:        "bfged",
		},
		{
			description: "substring not found",
			givenT:      "Hello how are you?",
			givenS:      "ok",
			want:        "",
		},
		{
			description: "substring with duplicate letters, and last match is",
			givenT:      "abcdbebe",
			givenS:      "bbe",
			want:        "bebe",
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := minWindowEducative(tc.givenT, tc.givenS)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("minWindowEducative() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
