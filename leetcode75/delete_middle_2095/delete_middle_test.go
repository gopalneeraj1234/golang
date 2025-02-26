package deletemiddle2095

import (
	"strconv"
	"strings"
	"testing"
)

var middleNodeTests = []struct {
	input []int
	want  []int
}{
	{[]int{1, 3, 4, 7, 1, 2, 6}, []int{1, 3, 4, 1, 2, 6}},
	{[]int{1, 2, 3, 4}, []int{1, 2, 4}},
	{[]int{2, 1}, []int{2}},
}

func createLinkedList(nums []int) *ListNode {
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

func isEqual(head1 *ListNode, head2 *ListNode) bool {
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
	} else {
		return false
	}
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

func TestDeleteMiddle(t *testing.T) {

	for i, tt := range middleNodeTests {
		t.Run("Test"+strconv.Itoa(i), func(t *testing.T) {
			got := deleteMiddle(createLinkedList(tt.input))
			wantList := createLinkedList(tt.want)
			t.Log(wantList)
			if !isEqual(got, wantList) {
				t.Errorf("Got %q, want %q", got, wantList)
			}
		})
	}
}
