package oddevenlist328

import (
	"leetcode75/utils"
	"strconv"
	"testing"
)

var oddEvenTests = []struct {
	input []int
	want  []int
}{
	{[]int{1, 2, 3, 4, 5}, []int{1, 3, 5, 2, 4}},
	{[]int{2, 1, 3, 5, 6, 4, 7}, []int{2, 3, 6, 7, 1, 5, 4}},
	{[]int{}, []int{}},
	{[]int{1}, []int{1}},
	{[]int{1, 2}, []int{1, 2}},
	{[]int{1, 2, 3}, []int{1, 3, 2}},
	{[]int{1, 2, 3, 4}, []int{1, 3, 2, 4}},
}

func TestOddEvenList(t *testing.T) {

	for i, tt := range oddEvenTests {
		t.Run("Test"+strconv.Itoa(i), func(t *testing.T) {
			gotList := oddEvenList(utils.CreateLinkedList(tt.input))
			wantList := utils.CreateLinkedList(tt.want)
			if !utils.IsEqual(gotList, wantList) {
				t.Errorf("Got %q, Want %q", gotList, wantList)
			}
		})
	}
}
