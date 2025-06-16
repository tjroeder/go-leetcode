package knapsack

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindMaxKnapsackProfit(t *testing.T) {
	testCases := []struct {
		description  string
		givenCap     int
		givenWeights []int
		givenValues  []int
		want         int
	}{
		{
			description:  "smallest is least valuable",
			givenCap:     30,
			givenWeights: []int{10, 20, 30},
			givenValues:  []int{22, 33, 44},
			want:         55,
		},
		{
			description:  "smallest is most valuable",
			givenCap:     5,
			givenWeights: []int{1, 2, 3, 5},
			givenValues:  []int{10, 5, 4, 8},
			want:         15,
		},
		{
			description:  "mix of most and least weight",
			givenCap:     50,
			givenWeights: []int{10, 20, 30},
			givenValues:  []int{600, 100, 120},
			want:         720,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := findMaxKnapsackProfit(tc.givenCap, tc.givenWeights, tc.givenValues)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("findMaxKnapsackProfit() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
