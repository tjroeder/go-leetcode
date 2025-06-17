package nthtribonaccinumber

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTribonacci(t *testing.T) {
	testCases := []struct {
		description string
		given       int
		want        int
	}{
		{
			description: "base test n=0",
			given:       0,
			want:        0,
		},
		{
			description: "base test n=1",
			given:       1,
			want:        1,
		},
		{
			description: "base test n=2",
			given:       2,
			want:        1,
		},
		{
			description: "n=4",
			given:       4,
			want:        4,
		},
		{
			description: "n=25",
			given:       25,
			want:        1389537,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := tribonacci(tc.given)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("tribonacci() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
