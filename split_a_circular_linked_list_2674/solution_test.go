package splitcircularlinkedlist

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

	five = &ListNode{
		Val:  5,
		Next: nil,
	}

	six = &ListNode{
		Val:  6,
		Next: nil,
	}

	seven = &ListNode{
		Val:  7,
		Next: nil,
	}
)

func TestSplitCircularLinkedList(t *testing.T) {
	t.Run("odd number of nodes test", func(t *testing.T) {
		one.Next = five
		five.Next = seven
		seven.Next = one
		given := one

		got := splitCircularLinkedList(given)

		one.Next = five
		five.Next = one
		want1 := one
		seven.Next = seven
		want2 := seven
		want := []*ListNode{want1, want2}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("splitCircularLinkedList() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("even number of nodes test", func(t *testing.T) {
		two.Next = six
		six.Next = one
		one.Next = five
		five.Next = two
		given := two

		got := splitCircularLinkedList(given)

		two.Next = six
		six.Next = two
		want1 := two
		one.Next = five
		five.Next = one
		want2 := one
		want := []*ListNode{want1, want2}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("splitCircularLinkedList() mismatch (-want +got):\n%s", diff)
		}
	})
}
