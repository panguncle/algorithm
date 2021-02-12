package tree

func LevelOrderBottom(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	type treeNodeWithLevel struct {
		*TreeNode
		Level int
	}

	queue := make([]*treeNodeWithLevel, 0)
	queue = append(queue, &treeNodeWithLevel{
		TreeNode: root,
		Level:    0,
	})

	ret := make([][]int, 0)
	for len(queue) > 0 {
		node := queue[0]
		if len(ret) < node.Level+1 {
			tmp := make([][]int, node.Level+1)
			copy(tmp, ret)
			ret = tmp
		}
		ret[node.Level] = append(ret[node.Level], node.Val)

		if node.Left != nil {
			queue = append(queue, &treeNodeWithLevel{
				TreeNode: node.Left,
				Level:    node.Level + 1,
			})
		}

		if node.Right != nil {
			queue = append(queue, &treeNodeWithLevel{
				TreeNode: node.Right,
				Level:    node.Level + 1,
			})
		}

		queue = queue[1:]
	}

	return reverse(ret)
}

func reverse(list [][]int) [][]int {
	ret := make([][]int, len(list), cap(list))
	for i := len(list) - 1; i >= 0; i-- {
		ret[len(list)-1-i] = list[i]
	}

	return ret
}
