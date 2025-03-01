package list

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	var head *ListNode
	prev := head
	for i, num := range nums {
		if i == 0 {
			head = &ListNode{
				Val: num,
			}
			prev = head
		} else {
			prev.Next = &ListNode{
				Val: num,
			}
			prev = prev.Next
		}
	}
	return head
}

func IsEqual(head1 *ListNode, head2 *ListNode) bool {
	temp1 := head1
	temp2 := head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val != temp2.Val {
			return false
		}
		temp1 = temp1.Next
		temp2 = temp2.Next
	}
	if temp1 == nil && temp2 == nil {
		return true
	}
	return false
}

func (l ListNode) String() string {
	var sb strings.Builder
	for l.Next != nil {
		sb.WriteString(strconv.Itoa(l.Val) + " -> ")
		l = *l.Next
	}
	sb.WriteString(strconv.Itoa(l.Val))
	return sb.String()
}
