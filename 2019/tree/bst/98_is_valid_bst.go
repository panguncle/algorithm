package bst

import (
	"math"

	"github.com/panguncle/algorithm/2019/tree"
)

// IsValidBST: 98.
func IsValidBST(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}

	// 递归的终止条件,
	// 除了考虑当前节点的左右子节点,
	// 还需要考虑到该节点的父节点以及更高的节点

	return validateBST(root, math.MinInt64, math.MaxInt64)
}

func validateBST(node *tree.TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	if node.Left != nil && (node.Val <= node.Left.Val || node.Left.Val <= min) {
		return false
	}

	if node.Right != nil && (node.Val >= node.Right.Val || node.Right.Val >= max) {
		return false
	}

	isLeftValid := validateBST(node.Left, min, node.Val)
	isRightValid := validateBST(node.Right, node.Val, max)

	return isLeftValid && isRightValid
}
