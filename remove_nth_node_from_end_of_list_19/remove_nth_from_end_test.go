package removenthfromend

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRemoveNthFromEnd(t *testing.T) {
	testCases := []struct {
		description string
		listNode    *ListNode
		element     int
		want        *ListNode
	}{
		{
			description: "head of 1",
			listNode: &ListNode{
				Val:  1,
				Next: nil,
			},
			element: 1,
			want:    &ListNode{},
		},
		{
			description: "remove head of the node",
			listNode: &ListNode{
				Val: 50,
				Next: &ListNode{
					Val: 40,
					Next: &ListNode{
						Val: 30,
						Next: &ListNode{
							Val:  20,
							Next: nil,
						},
					},
				},
			},
			element: 4,
			want: &ListNode{
				Val: 40,
				Next: &ListNode{
					Val: 30,
					Next: &ListNode{
						Val:  20,
						Next: nil,
					},
				},
			},
		},
		{
			description: "remove middle element",
			listNode: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val: 6,
									Next: &ListNode{
										Val:  7,
										Next: nil,
									},
								},
							},
						},
					},
				},
			},
			element: 5,
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 6,
								Next: &ListNode{
									Val:  7,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s: %d", tc.description, i), func(t *testing.T) {
			got := removeNthFromEnd(tc.listNode, tc.element)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("removeNthFromEnd() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
