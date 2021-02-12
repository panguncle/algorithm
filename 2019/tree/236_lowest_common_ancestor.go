package tree

// LowestCommonAncestor: 88
// Todo
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	// 如果当前节点就是要找的其中一个节点, 那么向父级汇报
	if root == nil || p == root || q == root {
		return root
	}

	// 找一下左子树和右子树
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)

	// 如果找完发现在左子树和右子树都存在, 那么就是这个了,
	// 自底向上汇报的
	if left != nil && right != nil {
		return root
	}

	// 不然就向上汇报在字数汇报
	if left != nil {
		return left
	}

	return right
}
