package dp

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var store map[*TreeNode]int

// HourseRobberIII 337
func HourseRobberIII(root *TreeNode) int {

	if root == nil {
		return 0
	}

	store = make(map[*TreeNode]int)
	return robber3Helper(root)
}

func robber3Helper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if v, exists := store[root]; exists {
		return v
	}
	var (
		val, llVal, lrVal, rlVal, rrVal int
	)

	val = root.Val
	if root.Left != nil {
		llVal = robber3Helper(root.Left.Left)
		lrVal = robber3Helper(root.Left.Right)
	}

	if root.Right != nil {
		rlVal = robber3Helper(root.Right.Left)
		rrVal = robber3Helper(root.Right.Right)
	}

	maxV := maxInt(val+llVal+lrVal+rlVal+rrVal, robber3Helper(root.Left)+robber3Helper(root.Right))
	store[root] = maxV
	return maxV
}

// func maxInt(items ...int) int {
// 	if len(items) == 0 {
// 		panic("items must not be empty")
// 	}

// 	max := items[0]
// 	for i := 1; i < len(items); i++ {
// 		if items[i] > max {
// 			max = items[i]
// 		}
// 	}

// 	return max
// }
