package maxtwinsum2130

import (
	"leetcode75/linkedlist"
	"testing"
)

var maxTwinTests = []struct {
	name  string
	input []int
	want  int
}{
	{"Example1", []int{5, 4, 2, 1}, 6},
	{"Example2", []int{4, 2, 2, 3}, 7},
	{"Example3", []int{1, 100000}, 100001},
}

func TestPairSum(t *testing.T) {
	for _, tt := range maxTwinTests {
		t.Run(tt.name, func(t *testing.T) {
			inputList := linkedlist.CreateLinkedList(tt.input)
			got := pairSum(inputList)
			if got != tt.want {
				t.Errorf("pairSum(%q)=%d, want=%d", inputList, got, tt.want)
			}
		})
	}
}
