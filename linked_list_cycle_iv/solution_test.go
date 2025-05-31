package linkedlistcycleiv

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

	seven = &ListNode{
		Val:  7,
		Next: nil,
	}

	eight = &ListNode{
		Val:  8,
		Next: nil,
	}

	nine = &ListNode{
		Val:  9,
		Next: nil,
	}

	eleven = &ListNode{
		Val:  11,
		Next: nil,
	}

	twelve = &ListNode{
		Val:  12,
		Next: nil,
	}

	fourteen = &ListNode{
		Val:  14,
		Next: nil,
	}

	sixteen = &ListNode{
		Val:  16,
		Next: nil,
	}

	eighteen = &ListNode{
		Val:  18,
		Next: nil,
	}
)

func TestCountCycleLength(t *testing.T) {
	t.Run("even number of nodes test", func(t *testing.T) {
		one.Next = three
		three.Next = five
		five.Next = seven
		seven.Next = nine
		nine.Next = eleven
		eleven.Next = seven
		given := one

		got := removeCycle(given)

		one.Next = three
		three.Next = five
		five.Next = seven
		seven.Next = nine
		nine.Next = eleven
		eleven.Next = nil
		want := one

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("removeCycle() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("no cycle in the linked list", func(t *testing.T) {
		twelve.Next = fourteen
		fourteen.Next = sixteen
		sixteen.Next = eighteen
		given := twelve

		got := removeCycle(given)

		if diff := cmp.Diff(given, got); diff != "" {
			t.Errorf("removeCycle() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("odd number of nodes with test", func(t *testing.T) {
		one.Next = two
		two.Next = three
		three.Next = four
		four.Next = five
		five.Next = two
		given := one

		got := removeCycle(given)

		one.Next = two
		two.Next = three
		three.Next = four
		four.Next = five
		five.Next = nil
		want := one

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("removeCycle() mismatch (-want +got):\n%s", diff)
		}
	})
}
