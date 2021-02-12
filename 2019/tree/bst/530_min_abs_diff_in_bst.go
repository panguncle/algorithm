package bst

import (
	"math"

	"github.com/panguncle/algorithm/2019/helper"
	"github.com/panguncle/algorithm/2019/tree"
)

// MinAbsDiffInBST: leetcode.530
// LABEL: Wrong
func MinAbsDiffInBST(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	minDiff := math.MaxInt64
	findMinABbsDiff(root, &minDiff)

	return minDiff
}

func findMinABbsDiff(node *tree.TreeNode, minDiff *int) {

	if node == nil {
		return
	}

	if node.Left != nil {
		*minDiff = helper.MinInt(node.Val-node.Left.Val, *minDiff)
	}

	if node.Right != nil {
		*minDiff = helper.MinInt(node.Right.Val-node.Val, *minDiff)
	}

	findMinABbsDiff(node.Left, minDiff)
	findMinABbsDiff(node.Right, minDiff)

}

// MinAbsDiffInBST_2: leetcode.530
// BST进行中序遍历, 那么就能得到一个有序的数组, 相差最小的就在相邻两个数中间
func MinAbsDiffInBST_2(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	var ret int
	bstInOrder(root, &ret, nil)

	return ret
}

func bstInOrder(node *tree.TreeNode, ret, pre *int) {
	if node == nil {
		return
	}

	bstInOrder(node.Left, ret, pre)
	if pre != nil {
		*ret = helper.MinInt(*ret, node.Val-*pre)
	}
	pre = &node.Val
	*pre = node.Val

	bstInOrder(node.Right, ret, pre)
}
