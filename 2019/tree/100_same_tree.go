package tree

// IsSameTree: leetcode.100
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	// 比较当前节点
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	// 比较左子树
	left := IsSameTree(p.Left, q.Left)
	// 比较右子树
	right := IsSameTree(p.Right, q.Right)

	// 左子树相等 && 右子树也要相等
	return left && right

}
