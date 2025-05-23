package validwordabbreviation

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestValidWordAbbreviation(t *testing.T) {
	testCases := []struct {
		description string
		givenWord   string
		givenAbbr   string
		want        bool
	}{
		{
			description: "textbook example",
			givenWord:   "internationalization",
			givenAbbr:   "i18n",
			want:        true,
		},
		{
			description: "single letter abbreviation",
			givenWord:   "a",
			givenAbbr:   "1",
			want:        true,
		},
		{
			description: "single letter not correct count",
			givenWord:   "a",
			givenAbbr:   "2",
			want:        false,
		},
		{
			description: "not abbreviation",
			givenWord:   "mindset",
			givenAbbr:   "min3et",
			want:        false,
		},
		{
			description: "leading zero in the abbreviation",
			givenWord:   "leadership",
			givenAbbr:   "lead04ip",
			want:        false,
		},
		{
			description: "not even the same letters abbreviation",
			givenWord:   "elite",
			givenAbbr:   "l33t",
			want:        false,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := validWordAbbreviation(tc.givenWord, tc.givenAbbr)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("validWordAbbreviation() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
