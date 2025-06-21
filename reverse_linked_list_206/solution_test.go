package reverselinkedlist

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	one = &ListNode{
		Val:  1,
		Next: nil,
	}

	two = &ListNode{
		Val:  2,
		Next: nil,
	}

	three = &ListNode{
		Val:  3,
		Next: nil,
	}

	four = &ListNode{
		Val:  4,
		Next: nil,
	}

	five = &ListNode{
		Val:  5,
		Next: nil,
	}
)

func TestReverseList(t *testing.T) {
	t.Run("odd number of nodes to reverse", func(t *testing.T) {
		one.Next = two
		two.Next = three
		three.Next = four
		four.Next = five
		five.Next = nil

		got := reverseList(one)

		five.Next = four
		four.Next = three
		three.Next = two
		two.Next = one
		one.Next = nil

		if diff := cmp.Diff(five, got); diff != "" {
			t.Errorf("reverseList() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("two nodes to reverse", func(t *testing.T) {
		one.Next = two
		two.Next = nil

		got := reverseList(one)

		two.Next = one
		one.Next = nil

		if diff := cmp.Diff(two, got); diff != "" {
			t.Errorf("reverseList() mismatch (-want +got):\n%s", diff)
		}
	})
}
