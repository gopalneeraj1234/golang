package list

import (
	"iter"
	"leetcode75/tree"
)

type TreeNode = tree.TreeNode
type Queue []TreeNode

func (q *Queue) IsEmpty() bool {
	return (*q).Size() == 0
}

func (q *Queue) Add(node *TreeNode) {
	*q = append(*q, *node)
}

func (q *Queue) Remove() *TreeNode {
	node := (*q)[0]
	(*q) = (*q)[1:]
	return &node
}

func (q *Queue) GetLast() *TreeNode {
	return &(*q)[len(*q)-1]
}
func (q *Queue) Size() int {
	return len(*q)
}

func (q *Queue) AddAll(anotherQ *Queue) {
	for node := range anotherQ.Iter() {
		q.Add(&node)
	}
}

func (q *Queue) Iter() iter.Seq[TreeNode] {
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
