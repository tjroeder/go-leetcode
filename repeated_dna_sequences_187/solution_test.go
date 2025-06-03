package repeateddnasequences

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindRepeatedDNASequences(t *testing.T) {
	testCases := []struct {
		description string
		given       string
		want        []string
	}{
		{
			description: "single sequence",
			given:       "AAAAAAAAAAAAA",
			want:        []string{"AAAAAAAAAA"},
		},
		{
			description: "multiple sequence",
			given:       "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
			want:        []string{"AAAAACCCCC", "CCCCCAAAAA"},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := findRepeatedDNASequences(tc.given)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("findRepeatedDNASequences() mismatch (-want +got):\n%s", diff)
			}

		})
	}
}
