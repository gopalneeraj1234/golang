package countgoodnodes1448

import (
	"leetcode75/tree"
	"math"
)

type TreeNode = tree.TreeNode

func goodNodes(root *TreeNode) int {
	return dfs(root, math.MinInt32)
}

func dfs(root *TreeNode, maxVal int) int {
	if root == nil {
		return 0
	}
	currCount := 0
	if root.Val >= maxVal {
		currCount = 1
	}
	currMax := max(root.Val, maxVal)
	return currCount + dfs(root.Left, currMax) + dfs(root.Right, currMax)
}
