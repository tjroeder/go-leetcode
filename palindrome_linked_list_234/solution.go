package palindromelinkedlist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// isPalindrome without extracting func
// func isPalindrome(head *ListNode) bool {
// 	slow, fast := head, head

// 	// find the middle of the linked list, pointed at slow
// 	for fast != nil && fast.Next != nil {
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 	}

// 	// create reverse of the second half of the linked list
// 	var prev, next, curr *ListNode = nil, nil, slow
// 	for curr != nil {
// 		next = curr.Next
// 		curr.Next = prev
// 		prev = curr
// 		curr = next
// 	}

// 	// compare the linked lists the original half and the reversed half
// 	orig, rev := head, prev
// 	isValid := false
// 	if orig.Val == rev.Val {
// 		isValid = true
// 	}
// 	for orig != nil && rev != nil {
// 		if orig.Val != rev.Val {
// 			isValid = false
// 		}
// 		orig = orig.Next
// 		rev = rev.Next
// 	}

// 	// revert back the second half of the linked list
// 	prev, next, curr = nil, nil, rev
// 	for curr != nil {
// 		next = curr.Next
// 		curr.Next = prev
// 		prev = curr
// 		curr = next
// 	}

// 	return isValid
// }

// isPalindrome with func extracted
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head

	// find the middle of the linked list, pointed at slow
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	reversedhalf := reverseList(slow)

	// compare the linked lists the original half and the reversedhalf
	isValid := compareTwoHalves(head, reversedhalf)

	// revert back the second half of the linked list
	reverseList(reversedhalf)

	return isValid
}

// reverseList reverses the linked list
func reverseList(pointer *ListNode) *ListNode {
	var prev, next, curr *ListNode = nil, nil, pointer
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// compareTwoHalves compares two linked lists halves
func compareTwoHalves(firstHalf, secondHalf *ListNode) bool {
	for firstHalf != nil && secondHalf != nil {
		if firstHalf.Val != secondHalf.Val {
			return false
		}
		firstHalf = firstHalf.Next
		secondHalf = secondHalf.Next
	}
	return true
}
