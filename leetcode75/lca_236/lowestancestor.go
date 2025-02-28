package lca236

import (
	"leetcode75/tree"
	"reflect"
)

type TreeNode = tree.TreeNode

type Stack []TreeNode

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) pop() *TreeNode {
	popped := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return &popped
}

func (s *Stack) push(root *TreeNode) {
	*s = append(*s, *root)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if p == q {
		return p
	}

	pathToP := NewStack()
	getPath(root, p, pathToP)
	found := IsChild(p, q)
	if found {
		return p
	}

	for !pathToP.IsEmpty() {
		parent := pathToP.pop()
		if reflect.DeepEqual(parent, q) {
			return parent
		}
		found := IsChild(parent, q)
		if found {
			return parent
		}
	}
	return nil
}

func NewStack() *Stack {
	return &Stack{}
}

func getPath(root, node *TreeNode, path *Stack) bool {
	if root == nil {
		return false
	}
	if root == node {
		return true
	}
	path.push(root)
	found := getPath(root.Left, node, path)
	if !found {
		found = getPath(root.Right, node, path)

	}
	if !found {
		path.pop()
	}
	return found
}

func IsChild(root, q *TreeNode) bool {
	if root == q {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return false
	}
	if root.Left == nil && root.Right != nil {
		return IsChild(root.Right, q)
	}
	found := IsChild(root.Left, q)
	if !found && root.Right != nil {
		found = IsChild(root.Right, q)
	}
	return found
}
