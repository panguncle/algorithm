package bst

import "github.com/panguncle/algorithm/2019/tree"

var (
	store map[int]int
)

// MostFrequentSubtreeSum 508
func MostFrequentSubtreeSum(root *tree.TreeNode) []int {
	if root == nil {
		return nil
	}

	store = make(map[int]int)
	sumSubTree(root)

	var max int
	var ret []int
	for k, v := range store {
		if v > max {
			ret = []int{k}
			max = v
			continue
		}

		if v == max {
			ret = append(ret, k)
			continue
		}
	}

	return ret
}

func sumSubTree(node *tree.TreeNode) int {
	if node == nil {
		return 0
	}

	leftVal := sumSubTree(node.Left)
	rightVal := sumSubTree(node.Right)
	val := node.Val + leftVal + rightVal

	store[val] += 1

	return val

}
