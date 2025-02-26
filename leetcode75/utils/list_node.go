package utils

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
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
