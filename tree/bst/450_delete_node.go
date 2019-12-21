package bst

import "github.com/panguncle/algorithm/tree"

// DeleteNodeInBST 450
// @see https://zxi.mytechroad.com/blog/?s=450.
func DeleteNodeInBST(root *tree.TreeNode, key int) *tree.TreeNode {
	if root == nil {
		return nil
	}

	// 如果找到这个节点, 那么就是要删除它,
	// 找到右子树中最小的节点/左子树中最大的节点
	if root.Val == key {
		if root.Left != nil && root.Right != nil {
			min := root.Right
			for min.Left != nil {
				min = min.Left
			}

			root.Val = min.Val
			root.Right = DeleteNodeInBST(root.Right, min.Val)
			return root
		} else if root.Left != nil {
			return root.Left
		} else {
			return root.Right
		}
		// 根据bst树的特点, 找寻需要删除的节点在左子树还是右子树
	} else if key > root.Val {
		root.Right = DeleteNodeInBST(root.Right, key)
	} else {
		root.Left = DeleteNodeInBST(root.Left, key)
	}
	return root
}
