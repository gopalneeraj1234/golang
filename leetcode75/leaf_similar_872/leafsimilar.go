package leafsimilar872

import (
	"leetcode75/tree"
	"reflect"
)

type TreeNode = tree.TreeNode

type List[T comparable] struct {
	elements []T
}

func (l List[T]) isDeepEqual(leaves2List List[int]) bool {
	return reflect.DeepEqual(l.elements, leaves2List.elements)
}

func (l *List[T]) add(val *T) {
	l.elements = append(l.elements, *val)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var leaves1List List[int]
	var leaves2List List[int]

	dfs(root1, &leaves1List)
	dfs(root2, &leaves2List)

	return leaves1List.isDeepEqual(leaves2List)
}

func dfs(root *TreeNode, leavesList *List[int]) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		leavesList.add(&root.Val)
	}
	dfs(root.Left, leavesList)
	dfs(root.Right, leavesList)
}
