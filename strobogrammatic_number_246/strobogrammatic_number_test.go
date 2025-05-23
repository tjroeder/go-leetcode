package strobogrammaticnumber

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsStrobogrammatic(t *testing.T) {
	testCases := []struct {
		description string
		given       string
		want        bool
	}{
		{
			description: "one number",
			given:       "0",
			want:        true,
		},
		{
			description: "two numbers",
			given:       "69",
			want:        true,
		},
		{
			description: "two numbers that are the same",
			given:       "88",
			want:        true,
		},
		{
			description: "not strobogrammatic",
			given:       "962",
			want:        false,
		},
		{
			description: "not strobogrammatic, one number",
			given:       "2",
			want:        false,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := isStrobogrammatic(tc.given)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("isStrobogrammatic() mismatch (-want +got):\n%s", diff)
			}
		})
	}

}
