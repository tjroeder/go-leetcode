package linkedlistcycleiii

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
		three.Next = four
		four.Next = five
		five.Next = six
		six.Next = seven
		seven.Next = eight
		eight.Next = six
		given := three

		got := countCycleLength(given)

		if diff := cmp.Diff(3, got); diff != "" {
			t.Errorf("countCycleLength() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("no cycle in the linked list", func(t *testing.T) {
		twelve.Next = fourteen
		fourteen.Next = sixteen
		sixteen.Next = eighteen
		given := twelve

		got := countCycleLength(given)

		if diff := cmp.Diff(0, got); diff != "" {
			t.Errorf("countCycleLength() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("odd number of nodes test", func(t *testing.T) {
		one.Next = two
		two.Next = three
		three.Next = four
		four.Next = five
		five.Next = two

		given := one

		got := countCycleLength(given)

		if diff := cmp.Diff(4, got); diff != "" {
			t.Errorf("countCycleLength() mismatch (-want +got):\n%s", diff)
		}
	})
}

func BenchmarkCountCycleLength(b *testing.B) {
	three.Next = four
	four.Next = five
	five.Next = six
	six.Next = seven
	seven.Next = eight
	eight.Next = six
	given := three
	b.ResetTimer()
	for b.Loop() {
		got := countCycleLength(given)

		if diff := cmp.Diff(3, got); diff != "" {
			b.Errorf("countCycleLength() mismatch (-want +got):\n%s", diff)
		}
	}

	// b.ResetTimer()
	// twelve.Next = fourteen
	// fourteen.Next = sixteen
	// sixteen.Next = eighteen
	// given = twelve
	// b.ResetTimer()
	// for b.Loop() {

	// 	got := countCycleLength(given)

	// 	if diff := cmp.Diff(0, got); diff != "" {
	// 		b.Errorf("countCycleLength() mismatch (-want +got):\n%s", diff)
	// 	}
	// }

	// b.ResetTimer()
	// one.Next = two
	// two.Next = three
	// three.Next = four
	// four.Next = five
	// five.Next = two

	// given = one
	// b.ResetTimer()
	// for b.Loop() {

	// 	got := countCycleLength(given)

	// 	if diff := cmp.Diff(4, got); diff != "" {
	// 		b.Errorf("countCycleLength() mismatch (-want +got):\n%s", diff)
	// 	}
	// }
}
