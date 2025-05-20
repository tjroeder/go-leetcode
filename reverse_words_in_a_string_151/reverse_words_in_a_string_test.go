package reversewordsinastring

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReverseWords(t *testing.T) {
	testCases := []struct {
		description string
		given       string
		want        string
	}{
		{
			description: "one word string",
			given:       " smart",
			want:        "smart",
		},
		{
			description: "multi word string",
			given:       "Reverse this string",
			want:        "string this Reverse",
		},
		{
			description: "multi spaces in the string",
			given:       "  MULTIPLE  SPACES  ",
			want:        "SPACES MULTIPLE",
		},
		{
			description: "with numbers",
			given:       "I have 0 pets but 1000s of plants",
			want:        "plants of 1000s but pets 0 have I",
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := reverseWords(tc.given)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("reverseWords() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
