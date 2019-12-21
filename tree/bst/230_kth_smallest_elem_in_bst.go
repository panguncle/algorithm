package bst

import "github.com/panguncle/algorithm/tree"

// KthSmallestElemInBST 230
func KthSmallestElemInBST(root *tree.TreeNode, k int) int {

	ret := make([]int, 0)
	inorderBST(root, &ret, k)

	return ret[k-1]
}

func inorderBST(root *tree.TreeNode, ret *[]int, k int) {
	if root == nil {
		return
	}
	if len(*ret) >= k {
		return
	}
	inorderBST(root.Left, ret, k)
	*ret = append(*ret, root.Val)
	inorderBST(root.Right, ret,k)
}

/*
 Reference 
 https://www.cnblogs.com/grandyang/p/4620012.html
*/
func KthSmallestElemInBSTII(root *tree.TreeNode, k int) int {

	var stack []*tree.TreeNode

    for k > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        pop := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        k--
        if k == 0 {
            return pop.Val
        }
        root = pop.Right
    }
    return -1
}