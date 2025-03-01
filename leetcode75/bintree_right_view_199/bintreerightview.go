package bintreerightview199

import (
	"iter"
	"leetcode75/tree"
)

type TreeNode = tree.TreeNode

type Queue []TreeNode

func (q *Queue) IsEmpty() bool {
	return (*q).size() == 0
}

func (q *Queue) add(node *TreeNode) {
	*q = append(*q, *node)
}

func (q *Queue) remove() *TreeNode {
	node := (*q)[0]
	(*q) = (*q)[1:]
	return &node
}

func (q *Queue) getLast() *TreeNode {
	return &(*q)[len(*q)-1]
}
func (q *Queue) size() int {
	return len(*q)
}

func (q *Queue) addAll(anotherQ *Queue) {
	for node := range anotherQ.iter() {
		q.add(&node)
	}
}

func (q *Queue) iter() iter.Seq[TreeNode] {
	return func(yield func(TreeNode) bool) {
		for _, element := range *q {
			if !yield(element) {
				return
			}
		}
	}
}

func NewQueue() *Queue {
	return &Queue{}
}
func rightSideView(root *TreeNode) []int {

	rightview := []int{}
	if root == nil {
		return rightview
	}
	q := NewQueue()
	q.add(root)
	rightview = append(rightview, q.getLast().Val)
	for !q.IsEmpty() {
		tempList := NewQueue()

		for !q.IsEmpty() {
			curr := q.remove()
			if curr.Left != nil {
				tempList.add(curr.Left)
			}
			if curr.Right != nil {
				tempList.add(curr.Right)
			}
		}
		if !tempList.IsEmpty() {
			rightview = append(rightview, tempList.getLast().Val)
		}
		q.addAll(tempList)
	}
	return rightview
}
