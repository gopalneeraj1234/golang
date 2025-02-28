package leafsimilar872

import (
	"leetcode75/tree"
	"math"
	"testing"
)

var leafSimilarTests = []struct {
	name      string
	root1nums []int
	root2nums []int
	want      bool
}{
	{"Example1", []int{3, 5, 1, 6, 2, 9, 8, math.MaxInt32, math.MaxInt32, 7, 4}, []int{3, 5, 1, 6, 7, 4, 2, math.MaxInt32, math.MaxInt32, math.MaxInt32, math.MaxInt32, math.MaxInt32, math.MaxInt32, 9, 8}, true},
	{"Example2", []int{1, 2, 3}, []int{1, 3, 2}, false},
}

func TestLeafSimilar(t *testing.T) {
	for _, tt := range leafSimilarTests {
		t.Run(tt.name, func(t *testing.T) {
			root1 := tree.CreateBinaryTree(tt.root1nums)
			root2 := tree.CreateBinaryTree(tt.root2nums)
			got := leafSimilar(root1, root2)
			if got != tt.want {
				t.Errorf("leafSimilar(%q, %q)=%t, want %t", root1, root2, got, tt.want)
			}
		})
	}
}
