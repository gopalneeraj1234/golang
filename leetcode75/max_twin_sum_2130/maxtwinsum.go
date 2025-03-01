package maxtwinsum2130

import (
	"leetcode75/list"
	"math"
)

type ListNode = list.ListNode

type Stack []int

func NewStack() *Stack {
	s := make(Stack, 0)
	return &s
}

func (s *Stack) push(element int) {
	*s = append(*s, element)
}

func (s *Stack) pop() int {
	element := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return element
}

func pairSum(head *ListNode) int {
	s := NewStack()
	slowPtr := head
	fastPtr := head.Next
	for fastPtr != nil {
		s.push(slowPtr.Val)
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next
		if fastPtr != nil {
			fastPtr = fastPtr.Next
		}
	}

	maxSum := math.MinInt32
	for slowPtr != nil {
		maxSum = max(maxSum, slowPtr.Val+s.pop())
		slowPtr = slowPtr.Next
	}
	return maxSum
}
