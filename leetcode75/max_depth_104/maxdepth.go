package maxdepth104

import "leetcode75/tree"

type TreeNode = tree.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}
