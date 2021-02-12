package tree

func NAryTreePostOrderTraversal(root *NTreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	naryTreePostOrderTraversal(&ret, root)
	return ret
}

func naryTreePostOrderTraversal(ret *[]int, root *NTreeNode) {
	if root == nil {
		return
	}

	if root.Children != nil {
		for i := range root.Children {
			naryTreePostOrderTraversal(ret, root.Children[i])
		}
	}

	*ret = append(*ret, root.Val)
}
