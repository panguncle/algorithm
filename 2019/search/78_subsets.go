package search

// Subsets: leetcode.78
func Subsets(nums []int) [][]int {

	var ret [][]int
	if len(nums) == 0 {
		return ret
	}

	var cur []int
	subsets(&ret, cur, nums, 0)
	return ret
}

func subsets(ret *[][]int, cur, nums []int, level int) {
	if level > len(nums) {
		return
	}
	tmp := make([]int, len(cur))
	copy(tmp, cur)
	*ret = append(*ret, tmp)

	for i := level; i < len(nums); i++ {
		cur = append(cur, nums[i])
		subsets(ret, cur, nums, i+1)
		cur = cur[:len(cur)-1]
	}
}
