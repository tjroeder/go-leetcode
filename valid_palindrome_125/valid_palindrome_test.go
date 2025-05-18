package validpalindrome

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestValidPalindrome(t *testing.T) {
	testCases := []struct {
		description string
		given       string
		want        bool
	}{
		{
			description: "is palindrome",
			given:       "aa",
			want:        true,
		},
		{
			description: "is palindrome with only special characters",
			given:       "&%#**(!)",
			want:        true,
		},
		{
			description: "is palindrome with special characters",
			given:       "Madam, in Eden, Im Adam",
			want:        true,
		},
		{
			description: "is not a palindrome",
			given:       "Madam, in Eden, Im Adam",
			want:        true,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := isValidPalindrome(tc.given)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("isValidPalindrome() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
