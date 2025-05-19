package removenthfromend

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil && n == 1 {
		return new(ListNode)
	}

	l, r := head, head
	for range n {
		r = r.Next
	}
	if r == nil {
		return l.Next
	}

	for r.Next != nil {
		r = r.Next
		l = l.Next
	}
	l.Next = l.Next.Next

	return head
}
