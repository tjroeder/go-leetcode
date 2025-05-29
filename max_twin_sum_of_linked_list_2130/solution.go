package maxtwinsumoflinkedlist

// / Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func twinSum(head *ListNode) int {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	reversedHalf := reverseLinkedList(slow)
	maxSum := 0
	curr := head

	for reversedHalf != nil {
		if (curr.Val + reversedHalf.Val) > maxSum {
			maxSum = curr.Val + reversedHalf.Val
		}
		curr = curr.Next
		reversedHalf = reversedHalf.Next
	}
	return maxSum
}

// reverseLinkedList accepts a pointer to a ListNode
// and reverses the remaining nodes in the Linked List
func reverseLinkedList(pointer *ListNode) *ListNode {
	var prev, next, curr *ListNode = nil, nil, pointer
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
