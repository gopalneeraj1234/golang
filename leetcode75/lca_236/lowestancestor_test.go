package lca236

import (
	"leetcode75/tree"
	"math"
	"reflect"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	root := tree.CreateBinaryTree([]int{3, 5, 1, 6, 2, 0, 8, math.MaxInt32, math.MaxInt32, 7, 4})
	p, _ := root.GetNodeWithVal(5)
	q, _ := root.GetNodeWithVal(1)
	want, _ := root.GetNodeWithVal(3)
	got := lowestCommonAncestor(root, p, q)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("lowestCommonAncestor(%q,%q,%q)=%q, want %q", root, p, q, got, want)
	}
}

func TestLowestCommonAncestorSmall(t *testing.T) {
	root := tree.CreateBinaryTree([]int{1, 2})
	p, _ := root.GetNodeWithVal(2)
	q, _ := root.GetNodeWithVal(1)
	want, _ := root.GetNodeWithVal(1)
	got := lowestCommonAncestor(root, p, q)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("lowestCommonAncestor(%q,%q,%q)=%q, want %q", root, p, q, got, want)
	}
}

func TestLowestCommonAncestorSmallRight(t *testing.T) {
	root := tree.CreateBinaryTree([]int{2, math.MaxInt32, 1})
	p, _ := root.GetNodeWithVal(1)
	q, _ := root.GetNodeWithVal(2)
	want, _ := root.GetNodeWithVal(2)
	got := lowestCommonAncestor(root, p, q)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("lowestCommonAncestor(%q,%q,%q)=%q, want %q", root, p, q, got, want)
	}
}
