package tree

import (
	"github.com/panguncle/algorithm/helper"
)

// MaxDepthOfBinaryTree: leetcode.104
func MaxDepthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type TreeNodeWithLevel struct {
		*TreeNode
		Level int
	}
	// BFS, 带上层级
	queue := make([]*TreeNodeWithLevel, 0)

	rootN := &TreeNodeWithLevel{
		TreeNode: root,
		Level:    0,
	}

	queue = append(queue, rootN)
	var level int
	for len(queue) > 0 {
		node := queue[0]
		if node.Level > level {
			level = node.Level
		}
		if node.Left != nil {
			queue = append(queue, &TreeNodeWithLevel{
				TreeNode: node.Left,
				Level:    node.Level + 1,
			})
		}

		if node.Right != nil {
			queue = append(queue, &TreeNodeWithLevel{
				TreeNode: node.Right,
				Level:    node.Level + 1,
			})
		}

		queue = queue[1:]
	}

	return level + 1
}

func MaxDepthOfBinaryTree_2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := MaxDepthOfBinaryTree_2(root.Left)
	rightDepth := MaxDepthOfBinaryTree_2(root.Right)

	return helper.MaxInt(leftDepth, rightDepth) + 1
}
