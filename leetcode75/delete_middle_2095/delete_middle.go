package deletemiddle2095

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
	slowPtr := head
	fastPtr := head

	fastPtr = fastPtr.Next

	var prevSlowPtr *ListNode

	for fastPtr != nil {
		prevSlowPtr = slowPtr
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next
		if fastPtr != nil {
			fastPtr = fastPtr.Next
		} else {
			break
		}
	}
	if prevSlowPtr == nil {
		return nil
	} else {
		prevSlowPtr.Next = slowPtr.Next
	}
	return head
}
