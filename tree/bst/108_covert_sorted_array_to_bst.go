package bst

import "github.com/panguncle/algorithm/tree"

// ConvertSortedArrayToBST 108
func ConvertSortedArrayToBST(nums []int) *tree.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &tree.TreeNode{
		Val: nums[mid],
	}

	root.Left = ConvertSortedArrayToBST(nums[0:mid])

	if mid+1 < len(nums) {
		root.Right = ConvertSortedArrayToBST(nums[mid+1:])
	}

	return root
}
