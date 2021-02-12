package tree

// IsSymmetric: leetcode.101
func IsSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	return isSymmetric(root.Left, root.Right)
}

func isSymmetric(nodeA *TreeNode, nodeB *TreeNode) bool {
	if nodeA == nil && nodeB == nil {
		return true
	}

	if nodeA == nil || nodeB == nil {
		return false
	}

	if nodeA.Val != nodeB.Val {
		return false
	}

	a := isSymmetric(nodeA.Left, nodeB.Right)
	b := isSymmetric(nodeA.Right, nodeB.Left)

	return a && b
}
