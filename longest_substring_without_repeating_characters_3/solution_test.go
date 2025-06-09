package longestsubstringwithoutrepeatingcharacters

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		description string
		given       string
		want        int
	}{
		{
			description: "base example (`abc`)",
			given:       "abcabcbb",
			want:        3,
		},
		{
			description: "all duplicate letters",
			given:       "bbbbb",
			want:        1,
		},
		{
			description: "subsequence(`pwke`) is not a substring(`wke`)",
			given:       "pwwkew",
			want:        3,
		},
		{
			description: "just a single space",
			given:       " ",
			want:        1,
		},
		{
			description: "two characters",
			given:       "au",
			want:        2,
		},
		{
			description: "longer with multiple duplicates",
			given:       "ggububgvfk",
			want:        6,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := lengthOfLongestSubstring(tc.given)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("lengthOfLongestSubstring() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
