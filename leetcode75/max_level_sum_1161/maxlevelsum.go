package maxlevelsum1161

import (
	"leetcode75/list"
	"leetcode75/tree"
)

type TreeNode = tree.TreeNode

func maxLevelSum(root *TreeNode) int {
	q := list.NewQueue[TreeNode]()

	q.Add(root)
	maxSum := root.Val
	maxSumLevel := 1
	currLevel := 1
	for !q.IsEmpty() {
		tempQ := list.NewQueue[TreeNode]()
		sum := 0
		for !q.IsEmpty() {
			node := q.Remove()
			if node.Left != nil {
				tempQ.Add(node.Left)
			}
			if node.Right != nil {
				tempQ.Add(node.Right)
			}
			sum += node.Val
		}
		q.AddAll(tempQ)
		if maxSum < sum {
			maxSum = sum
			maxSumLevel = currLevel
		}
		maxSum = max(maxSum, sum)
		currLevel++
	}
	return maxSumLevel
}
