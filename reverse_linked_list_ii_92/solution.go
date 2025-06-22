package reverselinkedlistii

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{Val: 0}
	dummy.Next = head
	prev := dummy

	// move through the linked list til just before left bound with prev
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	// set current node after finding the left bound
	curr := prev.Next
	// now we need to reverse between the right and left bounds
	for i := 0; i < right-left; i++ {
		nextNode := curr.Next     // move next node forward
		curr.Next = nextNode.Next // curr node moves one ahead of next, or head.Next.Next
		nextNode.Next = prev.Next // reverse next.Next to prev.Next
		prev.Next = nextNode      // prev.Next now is set to next
	}

	return dummy.Next // this is the new head (it could be the old head)
}
