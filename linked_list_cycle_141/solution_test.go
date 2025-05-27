package linkedlistcycle

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	negFour = &ListNode{
		Val:  -4,
		Next: nil,
	}

	zero = &ListNode{
		Val:  0,
		Next: nil,
	}

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
)

func TestHasCycle(t *testing.T) {
	t.Run("base test", func(t *testing.T) {
		three.Next = two
		two.Next = zero
		zero.Next = negFour
		negFour.Next = two

		given := three

		got := hasCycle(given)
		if diff := cmp.Diff(true, got); diff != "" {
			t.Errorf("hasCycle() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("short linked list", func(t *testing.T) {
		one.Next = two
		two.Next = one
		given := one

		got := hasCycle(given)
		if diff := cmp.Diff(true, got); diff != "" {
			t.Errorf("hasCycle() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("one node has no cycle", func(t *testing.T) {
		one.Next = nil
		given := one

		got := hasCycle(given)
		if diff := cmp.Diff(false, got); diff != "" {
			t.Errorf("hasCycle() mismatch (-want +got):\n%s", diff)
		}
	})
}
