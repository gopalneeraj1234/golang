package longestzigzag1372

import (
	"leetcode75/tree"
)

type TreeNode = tree.TreeNode

func longestZigZag(root *TreeNode) int {
	return dfs(root, false, 0, true) - 1

}

func dfs(root *TreeNode, leftChild bool, currLen int, rootNode bool) int {
	if root == nil {
		return currLen
	}

	if rootNode {
		rightLen := dfs(root.Right, false, 1, false)
		leftLen := dfs(root.Left, true, 1, false)
		return max(rightLen, leftLen)
	}

	if leftChild {
		rightLen := dfs(root.Right, false, currLen+1, false)
		leftLen := dfs(root.Left, true, 1, false)
		return max(rightLen, leftLen)
	} else {
		rightLen := dfs(root.Right, false, 1, false)
		leftLen := dfs(root.Left, true, currLen+1, false)
		return max(rightLen, leftLen)
	}

}
