package bst

import "github.com/panguncle/algorithm/2019/tree"

// InsertIntoBST 701
func InsertIntoBST(root *tree.TreeNode, val int) *tree.TreeNode {
	if root == nil {
		return &tree.TreeNode{
			Val: val,
		}
	}

	if root.Val > val {
		root.Left = InsertIntoBST(root.Left, val)
	} else {
		root.Right = InsertIntoBST(root.Right, val)
	}

	return root

}
