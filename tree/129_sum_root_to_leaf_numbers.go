package tree

import (
	"strconv"
	"strings"
)

// SumRootToLeafNumbers: 129
func SumRootToLeafNumbers(root *TreeNode) int {

	if root == nil {
		return 0
	}

	var (
		store []int
		cur   string
		sum   int
	)

	sumRootToLeafNumbers(root, &store, cur)
	for _, v := range store {
		sum += v
	}
	return sum
}

func sumRootToLeafNumbers(node *TreeNode, store *[]int, cur string) {
	if node == nil {
		return
	}

	cur += strconv.Itoa(node.Val)

	if node.Left == nil && node.Right == nil {
		curNum, _ := strconv.Atoi(strings.TrimLeft(cur, "0"))
		*store = append(*store, curNum)
		return
	}

	if node.Left != nil {
		sumRootToLeafNumbers(node.Left, store, cur)
	}

	if node.Right != nil {
		sumRootToLeafNumbers(node.Right, store, cur)
	}
}
