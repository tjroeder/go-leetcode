package nextpalidrome

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNextPalindrome(t *testing.T) {
	testCases := []struct {
		description string
		given       string
		want        string
	}{
		{
			description: "basic test",
			given:       "1221",
			want:        "2112",
		},
		{
			description: "no next palindrome",
			given:       "32123",
			want:        "",
		},
		{
			description: "single digit, no next palindrome",
			given:       "5",
			want:        "",
		},
		{
			description: "larger example",
			given:       "45544554",
			want:        "54455445",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := nextPalindrome(tc.given)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("nextPalindrome() mismatch (-want +got):\n%s", diff)
			}

		})
	}
}
