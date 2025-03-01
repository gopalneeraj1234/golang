package bst700

import (
	"leetcode75/tree"
	"reflect"
	"testing"
)

func TestSearchBST(t *testing.T) {
	root := tree.CreateBinaryTree([]int{4, 2, 7, 1, 3})
	want, _ := root.GetNodeWithVal(2)
	got := searchBST(root, 2)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("searchBST(%q, %d)=%q, want %q", root, 2, got, want)
	}
}
