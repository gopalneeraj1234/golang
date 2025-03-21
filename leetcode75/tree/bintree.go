package tree

import (
	"math"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}

func (t *TreeNode) GetNodeWithVal(val int) (*TreeNode, bool) {
	if t == nil {
		return nil, false
	}
	if t.Val == val {
		return t, true
	}
	node, leftFound := t.Left.GetNodeWithVal(val)
	if !leftFound {
		return t.Right.GetNodeWithVal(val)
	}
	return node, leftFound
}

func CreateBinaryTree(nums []int) *TreeNode {
	return createBinaryTreeRecur(nums, 0)
}

func (t *TreeNode) String() string {
	var sb strings.Builder
	q := []TreeNode{*t}
	for len(q) > 0 {
		curr := q[0]
		if len(q) > 1 {
			q = q[1:]
		} else {
			q = nil
		}
		sb.WriteString(strconv.Itoa(curr.Val) + " ")
		if curr.Left != nil {
			q = append(q, *curr.Left)
		}
		if curr.Right != nil {
			q = append(q, *curr.Right)
		}
	}
	return sb.String()
}
func createBinaryTreeRecur(nums []int, index int) *TreeNode {
	if index >= len(nums) {
		return nil
	}
	if nums[index] == math.MaxInt32 {
		return nil
	}
	root := NewTreeNode(nums[index])
	root.Left = createBinaryTreeRecur(nums, 2*index+1)
	root.Right = createBinaryTreeRecur(nums, 2*index+2)
	return root
}

func GetSize(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return GetSize(root.Left) + GetSize(root.Right) + 1
}

func GetArray(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	treeArr := make([]int, 0, GetSize(root))

	treeToArr(root, 0, &treeArr)
	return treeArr
}

func treeToArr(root *TreeNode, rootIndex int, treeArr *[]int) {
	if len(*treeArr) <= rootIndex {
		*treeArr = (*treeArr)[:rootIndex+1]
	}
	(*treeArr)[rootIndex] = root.Val
	if root.Left != nil {
		treeToArr(root.Left, 2*rootIndex+1, treeArr)
	}
	if root.Right != nil {
		treeToArr(root.Right, 2*rootIndex+2, treeArr)
	}
}
