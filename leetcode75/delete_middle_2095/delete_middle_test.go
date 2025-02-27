package delete_middle_2095

import (
	"strconv"
	"testing"

	"leetcode75/linkedlist"
)

var middleNodeTests = []struct {
	input []int
	want  []int
}{
	{[]int{1, 3, 4, 7, 1, 2, 6}, []int{1, 3, 4, 1, 2, 6}},
	{[]int{1, 2, 3, 4}, []int{1, 2, 4}},
	{[]int{2, 1}, []int{2}},
}

func TestDeleteMiddle(t *testing.T) {

	for i, tt := range middleNodeTests {
		t.Run("Test"+strconv.Itoa(i), func(t *testing.T) {
			got := deleteMiddle(linkedlist.CreateLinkedList(tt.input))
			wantList := linkedlist.CreateLinkedList(tt.want)
			t.Log(wantList)
			if !linkedlist.IsEqual(got, wantList) {
				t.Errorf("Got %q, want %q", got, wantList)
			}
		})
	}
}
