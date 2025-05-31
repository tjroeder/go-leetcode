package splitcircularlinkedlist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// with a lagging middle and last node, more memory but fast
func splitCircularLinkedList(head *ListNode) []*ListNode {
	var slow, fast, middleNode, lastNode *ListNode = head, head, nil, nil

	for {
		middleNode = slow
		lastNode = fast
		slow = slow.Next
		fast = fast.Next
		if fast == head {
			break
		}
		lastNode = fast
		fast = fast.Next
		if fast == head {
			break
		}
	}

	firstHalf, secondHalf := head, slow

	middleNode.Next = firstHalf
	lastNode.Next = secondHalf
	return []*ListNode{firstHalf, secondHalf}
}

// simpler but slower, less memory too
// func splitCircularLinkedList(head *ListNode) []*ListNode {
// 	slow, fast := head, head
// 	for fast.Next != head && fast.Next.Next != head {
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 	}

// 	secondHalf := slow.Next

// 	if fast.Next == head {
// 		fast.Next = secondHalf
// 	} else { // if fast.Next.Next == head
// 		fast.Next.Next = secondHalf
// 	}

// 	slow.Next = head
// 	return []*ListNode{head, secondHalf}
// }

// one more try with different loop conditions, still slower than the first, but faster than the second
// due to starting fast = head.Next
// func splitCircularLinkedList(head *ListNode) []*ListNode {
// 	slow, fast := head, head.Next
// 	for fast.Next != head {
// 		slow = slow.Next
// 		if fast.Next.Next != head {
// 			fast = fast.Next.Next
// 		} else {
// 			fast = fast.Next
// 		}
// 	}
// 	secondHalf := slow.Next
// 	slow.Next = head
// 	fast.Next = secondHalf
// 	return []*ListNode{head, secondHalf}
// }
