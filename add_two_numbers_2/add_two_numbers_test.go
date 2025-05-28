package addtwonumbers

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAddTwoNumbers(t *testing.T) {
	testCases := []struct {
		description string
		l1          *ListNode
		l2          *ListNode
		want        *ListNode
	}{
		{
			description: "single node",
			l1:          &ListNode{Val: 0, Next: nil},
			l2:          &ListNode{Val: 0, Next: nil},
			want:        &ListNode{Val: 0, Next: nil},
		},
		{
			description: "simple test",
			l1:          &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}},
			l2:          &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}},
			want:        &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: nil}}},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			t.Parallel()

			got := addTwoNumbers(tc.l1, tc.l2)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("addTwoNumbers() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
