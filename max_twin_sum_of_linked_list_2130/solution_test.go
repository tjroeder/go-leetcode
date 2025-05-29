package maxtwinsumoflinkedlist

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTwinSum(t *testing.T) {
	testCases := []struct {
		description string
		given       *ListNode
		want        int
	}{
		{
			description: "twin sum of n=4, where both sums are the same",
			given: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
			want: 6,
		},
		{
			description: "twin sum of n=4, where one sum is greater",
			given: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
			},
			want: 7,
		},
		{
			description: "only one sum possible with n=2",
			given: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  100000,
					Next: nil,
				},
			},
			want: 100001,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			got := twinSum(tc.given)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("twinSum() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
