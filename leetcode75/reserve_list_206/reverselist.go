package reservelist206

import "leetcode75/linkedlist"

type ListNode = linkedlist.ListNode

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	var prevPtr *ListNode
	currPtr := head
	nextPtr := head.Next

	for nextPtr != nil {
		currPtr.Next = prevPtr
		prevPtr = currPtr
		currPtr = nextPtr
		nextPtr = nextPtr.Next
	}
	currPtr.Next = prevPtr
	return currPtr
}
