package bst

import "github.com/panguncle/algorithm/2019/tree"

var (
	store map[int]int
)

// FindModeInBST 501
func FindModeInBST(root *tree.TreeNode) []int {
	if root == nil {
		return nil
	}

	store = make(map[int]int)
	inorder(root)
	var (
		max int
		ret []int
	)

	for k, v := range store {
		if v == max {
			ret = append(ret, k)
			continue
		}

		if v > max {
			ret = []int{k}
			max = v
			continue
		}
	}

	return ret

}

func inorder(root *tree.TreeNode) {
	if root == nil {
		return
	}

	inorder(root.Left)
	store[root.Val] += 1
	inorder(root.Right)
}
