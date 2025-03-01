package reservelist206

import (
	"leetcode75/list"
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
			inputList := list.CreateLinkedList(tt.input)
			gotList := reverseList(inputList)
			wantList := list.CreateLinkedList(tt.want)
			if !list.IsEqual(gotList, wantList) {
				t.Errorf("reverseList(%q) = %q, want = %q", inputList, gotList, wantList)
			}
		})
	}
}
