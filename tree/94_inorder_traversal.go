package tree

// InOrderTraversal: leetcode.94
func InOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	inOrderTraversal(&ret, root)
	return ret
}

func inOrderTraversal(ret *[]int, root *TreeNode) {
	if root == nil {
		return
	}

	inOrderTraversal(ret, root.Left)
	*ret = append(*ret, root.Val)
	inOrderTraversal(ret, root.Right)
}
