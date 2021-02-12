package tree

type NTreeNode struct {
	Val      int
	Children []*NTreeNode
}

// NAryTreePreOrder
func NAryTreePreOrder(root *NTreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	nAryTreePreOrder(&ret, root)
	return ret
}

func nAryTreePreOrder(ret *[]int, root *NTreeNode) {
	if root == nil {
		return
	}

	*ret = append(*ret, root.Val)
	if len(root.Children) == 0 {
		return
	}

	for i := range root.Children {
		nAryTreePreOrder(ret, root.Children[i])
	}
}
