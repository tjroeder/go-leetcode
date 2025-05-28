package palindromelinkedlist

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		description string
		given       *ListNode
		want        bool
	}{
		{
			description: "only one node is palindrome",
			given: &ListNode{
				Val:  1,
				Next: nil,
			},
			want: true,
		},
		{
			description: "two nodes not palindrome",
			given: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
			want: false,
		},
		{
			description: "longer even node palindrome",
			given: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
			want: true,
		},
		{
			description: "longer odd node palindrome",
			given: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
			},
			want: true,
		},
		{
			description: "longer not palindrome",
			given: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
			want: false,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			got := isPalindrome(tc.given)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("isPalindrome() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
