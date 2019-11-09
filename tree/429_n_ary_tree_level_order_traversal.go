package tree

func NAryTreeLevelOrderTraversal(root *NTreeNode) [][]int {
	var ret [][]int

	if root == nil {
		return ret
	}

	type NTreeNodeWithLevel struct {
		*NTreeNode
		Level int
	}

	queue := make([]*NTreeNodeWithLevel, 0)
	queue = append(queue, &NTreeNodeWithLevel{
		NTreeNode: root,
		Level:     1,
	})
	ret = make([][]int, 0)
	for len(queue) > 0 {
		node := queue[0]
		if len(ret) < node.Level {
			tmp := make([][]int, node.Level)
			copy(tmp, ret)
			ret = tmp
		}

		ret[node.Level-1] = append(ret[node.Level-1], node.Val)
		if len(node.Children) > 0 {
			for i := range node.Children {
				queue = append(queue, &NTreeNodeWithLevel{
					NTreeNode: node.Children[i],
					Level:     node.Level + 1,
				})
			}
		}

		queue = queue[1:]
	}
	return ret
}
