package linkedlistcycleiv

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head

	for {
		if fast.Next == nil || fast.Next.Next == nil {
			return head
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	for fast.Next != slow {
		fast = fast.Next
	}

	fast.Next = nil
	return head
}
