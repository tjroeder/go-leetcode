package linkedlistcycleiii

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// with no nested loop
func countCycleLength(head *ListNode) int {
	slow, fast := head, head
	for {
		if fast.Next == nil || fast.Next.Next == nil {
			return 0
		}
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}

	count := 0
	for {
		fast = fast.Next
		count++
		if slow == fast {
			break
		}
	}
	return count
}

// with nested loop
// func countCycleLength(head *ListNode) int {
// 	slow, fast := head, head

// 	for fast != nil && fast.Next != nil {
// 		slow = slow.Next
// 		fast = fast.Next.Next

// 		if slow == fast {
// 			length := 1
// 			slow = slow.Next

// 			for slow != fast {
// 				length++
// 				slow = slow.Next
// 			}

// 			return length
// 		}
// 	}
// 	return 0
// }
