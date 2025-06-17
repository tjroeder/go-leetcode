package coinchange

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCoinChange(t *testing.T) {
	testCases := []struct {
		description string
		givenCoins  []int
		givenAmount int
		want        int
	}{
		{
			description: "base test 11 = 5 + 5 + 1",
			givenCoins:  []int{1, 2, 5},
			givenAmount: 11,
			want:        3,
		},
		// {
		// 	description: "cannot be made up of any combination of coins",
		// 	givenCoins:  []int{2},
		// 	givenAmount: 3,
		// 	want:        -1,
		// },
		// {
		// 	description: "zero amount",
		// 	givenCoins:  []int{1},
		// 	givenAmount: 0,
		// 	want:        0,
		// },
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := coinChange(tc.givenCoins, tc.givenAmount)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("coinChange() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
