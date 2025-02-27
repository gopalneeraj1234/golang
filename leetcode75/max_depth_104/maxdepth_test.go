package maxdepth104

import (
	"leetcode75/tree"
	"math"
	"testing"
)

var maxDepthTests = []struct {
	name  string
	input []int
	want  int
}{
	{"Example1", []int{3, 9, 20, math.MaxInt32, math.MaxInt32, 15, 7}, 3},
	{"Example1", []int{1, math.MaxInt32, 2}, 2},
}

func TestMaxDepth(t *testing.T) {
	for _, tt := range maxDepthTests {
		t.Run(tt.name, func(t *testing.T) {
			inputTree := tree.CreateBinaryTree(tt.input)
			got := maxDepth(inputTree)
			if got != tt.want {
				t.Errorf("maxDepth(%q) = %d, but want %d", inputTree, got, tt.want)
			}
		})
	}
}
