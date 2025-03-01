package delete_middle_2095

import "leetcode75/list"

type ListNode = list.ListNode

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
		if fastPtr == nil {
			break
		}
		fastPtr = fastPtr.Next
	}
	if prevSlowPtr == nil {
		return nil
	}
	prevSlowPtr.Next = slowPtr.Next
	return head
}
