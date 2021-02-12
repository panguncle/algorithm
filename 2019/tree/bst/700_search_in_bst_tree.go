package bst

import "github.com/panguncle/algorithm/2019/tree"

// SearchBST 700
func SearchBST(root *tree.TreeNode, val int) *tree.TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if val > root.Val {
		return SearchBST(root.Right, val)
	}

	return SearchBST(root.Left, val)
}
