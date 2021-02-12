package bst

import (
	"math"

	"github.com/panguncle/algorithm/2019/tree"
)

/*
 MinSmallDifferenceInBST 530
 性质: BST中序遍历, 升序数组
*/

var (
	prev *int
	min  int
)

func getMinimumDifference(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	min = math.MaxInt64
	inOrder(root)

	return min
}

func inOrder(node *tree.TreeNode) {
	if node == nil {
		return
	}
	inOrder(node.Left)
	if prev != nil {
		min = minV(min, abs(node.Val-*prev))
	}
	prev = &node.Val
	inOrder(node.Right)
}

func abs(i int) int {
	if i > 0 {
		return i
	}

	return -i
}

func minV(i, j int) int {
	if i < j {
		return i
	}

	return j
}
