package countgoodnodes1448

import (
	"leetcode75/tree"
	"math"
	"testing"
)

var goodNodesTest = []struct {
	name  string
	input []int
	want  int
}{
	{"Example1", []int{3, 1, 4, 3, math.MaxInt32, 1, 5}, 4},
	{"Example2", []int{3, 3, math.MaxInt32, 4, 2}, 3},
}

func TestGoodNodes(t *testing.T) {
	for _, tt := range goodNodesTest {
		t.Run(tt.name, func(t *testing.T) {
			inputTree := tree.CreateBinaryTree(tt.input)
			got := goodNodes(inputTree)
			if got != tt.want {
				t.Errorf("goodNodes(%q)=%d, want %d", inputTree, got, tt.want)
			}
		})
	}
}
