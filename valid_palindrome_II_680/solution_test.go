package validpalindromeII

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
			description: "basic test",
			given:       "aba",
			want:        true,
		},
		{
			description: "another test",
			given:       "abca",
			want:        true,
		},
		{
			description: "one more",
			given:       "eceec",
			want:        true,
		},
		{
			description: "have to delete more than one",
			given:       "abc",
			want:        false,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := validPalindrome(tc.given)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("validPalindrome() mismatch (-want +got):\n%s", diff)
			}

		})
	}
}
