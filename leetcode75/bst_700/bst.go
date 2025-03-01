package bst700

import "leetcode75/tree"

type TreeNode = tree.TreeNode

func searchBST(root *TreeNode, val int) *TreeNode {
	if root.Val == val {
		return root
	} else if root.Val < val && root.Right != nil {
		return searchBST(root.Right, val)
	} else if root.Left != nil {
		return searchBST(root.Left, val)
	}
	return nil
}
