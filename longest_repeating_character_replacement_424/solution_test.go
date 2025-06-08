package longestrepeatingcharacterreplacement

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCharacterReplacement(t *testing.T) {
	testCases := []struct {
		description string
		givenS      string
		givenK      int
		want        int
	}{
		{
			description: "replace either character with the other",
			givenS:      "ABAB",
			givenK:      2,
			want:        4,
		},
		{
			description: "replace one character to achieve the longest",
			givenS:      "AABABBA",
			givenK:      1,
			want:        4,
		},
		{
			description: "can replace all characters, so longest is givenS length",
			givenS:      "FZFZFZ",
			givenK:      6,
			want:        6,
		},
		{
			description: "no replacements needed, all characters are the same",
			givenS:      "XXXXX",
			givenK:      1,
			want:        5,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := characterReplacement(tc.givenS, tc.givenK)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("characterReplacement() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
