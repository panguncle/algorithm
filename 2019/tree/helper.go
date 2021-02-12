package tree

// TreeNode: Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildBinaryTreeFromArray(nodes []*int) *TreeNode {
	// nodes 当做完全二叉树来处理,
	var root *TreeNode
	if len(nodes) == 0 {
		return root
	}

	root = &TreeNode{
		Val: *nodes[0],
	}

	store := make([]*TreeNode, len(nodes))
	store[0] = root

	for i := 1; i < len(nodes); i++ {
		if nodes[i] == nil {
			continue
		}

		parentIdx := (i - 1) / 2
		mod := i % 2
		curNode := &TreeNode{
			Val: *nodes[i],
		}
		parentNode := store[parentIdx]
		if mod == 0 {
			parentNode.Right = curNode
		} else {
			parentNode.Left = curNode
		}

		store[i] = curNode
	}

	return root
}

func PrintBinaryTree(head *TreeNode) {

}
