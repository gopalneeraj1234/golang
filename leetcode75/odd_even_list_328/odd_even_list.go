package oddevenlist328

import "leetcode75/linkedlist"

type ListNode = linkedlist.ListNode

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	oddPtr := head
	evenPtr := head.Next
	firstEvenPtr := evenPtr
	var prevOddPtr *ListNode
	for oddPtr != nil && evenPtr != nil {
		prevOddPtr = oddPtr
		if oddPtr.Next == nil || oddPtr.Next.Next == nil {
			break
		}
		nextOddPtr := oddPtr.Next.Next
		if evenPtr.Next == nil || evenPtr.Next.Next == nil {
			evenPtr.Next = nil
			oddPtr.Next = nextOddPtr
			oddPtr = oddPtr.Next
			break
		}
		nextEvenPtr := evenPtr.Next.Next
		oddPtr.Next = nextOddPtr
		oddPtr = oddPtr.Next

		evenPtr.Next = nextEvenPtr
		evenPtr = evenPtr.Next
	}
	if oddPtr != nil {
		prevOddPtr = oddPtr
	}
	prevOddPtr.Next = firstEvenPtr
	return head
}
