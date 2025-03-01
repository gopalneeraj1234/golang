package maxlevelsum1161

import (
	"leetcode75/tree"
	"math"
	"testing"
)

var maxLevelTests = []struct {
	name  string
	input []int
	want  int
}{
	{"Example1", []int{1, 7, 0, 7, -8, math.MaxInt32, math.MaxInt32}, 2},
	{"Example2", []int{989, math.MaxInt32, 10250, 98693, -89388, math.MaxInt32, math.MaxInt32, math.MaxInt32, -32127}, 2},
}

func TestMaxLevelSum(t *testing.T) {
	for _, tt := range maxLevelTests {
		t.Run(tt.name, func(t *testing.T) {
			root := tree.CreateBinaryTree(tt.input)
			got := maxLevelSum(root)
			if got != tt.want {
				t.Errorf("maxLevelSum(%q)=%d, want %d", root, got, tt.want)
			}
		})
	}
}
