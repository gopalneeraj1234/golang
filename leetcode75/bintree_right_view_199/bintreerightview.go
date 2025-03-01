package bintreerightview199

import (
	"leetcode75/list"
	"leetcode75/tree"
)

type TreeNode = tree.TreeNode

func rightSideView(root *TreeNode) []int {

	rightview := []int{}
	if root == nil {
		return rightview
	}
	q := list.NewQueue[TreeNode]()
	q.Add(root)
	rightview = append(rightview, q.GetLast().Val)
	for !q.IsEmpty() {
		tempList := list.NewQueue[TreeNode]()

		for !q.IsEmpty() {
			curr := q.Remove()
			if curr.Left != nil {
				tempList.Add(curr.Left)
			}
			if curr.Right != nil {
				tempList.Add(curr.Right)
			}
		}
		if !tempList.IsEmpty() {
			rightview = append(rightview, tempList.GetLast().Val)
		}
		q.AddAll(tempList)
	}
	return rightview
}
