package bintreerightview199

import (
	"leetcode75/tree"
	"math"
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	root := tree.CreateBinaryTree([]int{1, 2, 3, math.MaxInt32, 5, math.MaxInt32, 4})
	got := rightSideView(root)
	want := []int{1, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("rightSideView(%q)=%d, want %d", root, got, want)
	}
}

func TestRightSideView2(t *testing.T) {
	root := tree.CreateBinaryTree([]int{1, 2, 3, 4, math.MaxInt32, math.MaxInt32, math.MaxInt32, 5})
	got := rightSideView(root)
	want := []int{1, 3, 4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("rightSideView(%q)=%d, want %d", root, got, want)
	}
}
