package minnummovestomakepalindrome

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMinMovesToMakePalindrome(t *testing.T) {
	testCases := []struct {
		desciption string
		given      string
		want       int
	}{
		{
			desciption: "one letter",
			given:      "a",
			want:       0,
		},
		{
			desciption: "simple test",
			given:      "aabb",
			want:       2,
		},
		{
			desciption: "little more difficult",
			given:      "letelt",
			want:       2,
		},
		{
			desciption: "odd number of letters, with one which does not match others",
			given:      "xaabb",
			want:       4,
		},
		{
			desciption: "odd number of letters, with one which does not match others not on the end",
			given:      "aabxb",
			want:       3,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.desciption, i), func(t *testing.T) {
			t.Parallel()

			got := minMovesToMakePalindrome(tc.given)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("minMovesToMakePalindrome() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
