package middleoflinkedlist

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

	six = &ListNode{
		Val:  6,
		Next: nil,
	}
)

func TestMiddleNode(t *testing.T) {
	t.Run("odd number of nodes test", func(t *testing.T) {
		one.Next = two
		two.Next = three
		three.Next = four
		four.Next = five

		given := one

		got := middleNode(given)
		if diff := cmp.Diff(three, got); diff != "" {
			t.Errorf("middleNode() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("even number of nodes test", func(t *testing.T) {
		one.Next = two
		two.Next = three
		three.Next = four
		four.Next = five
		five.Next = six

		given := one

		got := middleNode(given)
		if diff := cmp.Diff(four, got); diff != "" {
			t.Errorf("middleNode() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("one node test", func(t *testing.T) {
		one.Next = nil
		given := one

		got := middleNode(given)
		if diff := cmp.Diff(one, got); diff != "" {
			t.Errorf("middleNode() mismatch (-want +got):\n%s", diff)
		}
	})
}
