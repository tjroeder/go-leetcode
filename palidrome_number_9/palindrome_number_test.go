package palindromenumber

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsPalindromeNumber(t *testing.T) {
	testCases := []struct {
		description string
		given       int
		want        bool
	}{
		{
			description: "true: single digit",
			given:       0,
			want:        true,
		},
		{
			description: "true: simple palindrome",
			given:       121,
			want:        true,
		},
		{
			description: "false: negative integer case",
			given:       -121,
			want:        false,
		},
		{
			description: "false: simple fail case",
			given:       10,
			want:        false,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := isPalindromeNumber(tc.given)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("palindromeNumber() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
