package reservelist206

import (
	"leetcode75/linkedlist"
	"testing"
)

var reverseListTestInputs = []struct {
	name        string
	input, want []int
}{
	{name: "Example1", input: []int{1, 2, 3, 4, 5}, want: []int{5, 4, 3, 2, 1}},
	{name: "Example2", input: []int{1, 2}, want: []int{2, 1}},
	{name: "Example3", input: []int{}, want: []int{}},
}

func TestReverseList(t *testing.T) {
	for _, tt := range reverseListTestInputs {
		t.Run(tt.name, func(t *testing.T) {
			inputList := linkedlist.CreateLinkedList(tt.input)
			gotList := reverseList(inputList)
			wantList := linkedlist.CreateLinkedList(tt.want)
			if !linkedlist.IsEqual(gotList, wantList) {
				t.Errorf("reverseList(%q) = %q, want = %q", inputList, gotList, wantList)
			}
		})
	}
}
