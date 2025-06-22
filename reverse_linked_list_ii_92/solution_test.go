package reverselinkedlistii

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
	t.Run("reverse in the middle", func(t *testing.T) {
		one.Next = two
		two.Next = three
		three.Next = four
		four.Next = five
		five.Next = nil

		left, right := 2, 4
		got := reverseBetween(one, left, right)

		one.Next = four
		four.Next = three
		three.Next = two
		two.Next = five
		five.Next = nil

		if diff := cmp.Diff(one, got); diff != "" {
			t.Errorf("reverseBetween() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("one nodes to reverse does not change", func(t *testing.T) {
		one.Next = nil

		left, right := 1, 1
		got := reverseBetween(one, left, right)

		one.Next = nil

		if diff := cmp.Diff(one, got); diff != "" {
			t.Errorf("reverseBetween() mismatch (-want +got):\n%s", diff)
		}
	})
}
