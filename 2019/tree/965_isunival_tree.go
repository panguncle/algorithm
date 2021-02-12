package tree

// IsUnivalTree: leetcode.965
func IsUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	val := root.Val

	return nodeHasSameValue(root, val)
}

func nodeHasSameValue(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}

	if root.Val != val {
		return false
	}

	return nodeHasSameValue(root.Left, val) && nodeHasSameValue(root.Right, val)
}
