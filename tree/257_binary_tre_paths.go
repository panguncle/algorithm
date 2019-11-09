package tree

import "strconv"

// BinaryTreePaths: 257.
func BinaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	var (
		ret []string
		cur string
	)

	listBinaryTreePaths(root, &ret, cur)
	return ret
}

func listBinaryTreePaths(node *TreeNode, ret *[]string, cur string) {
	if node == nil {
		return
	}
	var tmpCur string
	if len(cur) > 0 {
		tmpCur += "->"
	}
	cur += tmpCur + strconv.Itoa(node.Val)
	if node.Left == nil && node.Right == nil {
		*ret = append(*ret, cur)
		return
	}

	if node.Left != nil {
		listBinaryTreePaths(node.Left, ret, cur)
	}

	if node.Right != nil {
		listBinaryTreePaths(node.Right, ret, cur)
	}
}
