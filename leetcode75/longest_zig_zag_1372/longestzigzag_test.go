package longestzigzag1372

import (
	"leetcode75/tree"
	"testing"
)

var longestZigZagTests = []struct {
	name  string
	input []int
	want  int
}{
	{"Example3", []int{1}, 0},
}

func TestLongestZigZag(t *testing.T) {
	for _, tt := range longestZigZagTests {
		t.Run(tt.name, func(t *testing.T) {
			inputTree := tree.CreateBinaryTree(tt.input)
			got := longestZigZag(inputTree)
			if got != tt.want {
				t.Errorf("longestZigZag(%q) = %d, want %d", inputTree, got, tt.want)
			}
		})
	}
}
