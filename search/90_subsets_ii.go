package search

import "sort"

// SubsetsWithDup: leetcode.
func SubsetsWithDup(nums []int) [][]int {

	var ret [][]int
	if len(nums) == 0 {
		return ret
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var cur []int
	subsetsWithDup(&ret, cur, nums, 0)
	return ret
}

func subsetsWithDup(ret *[][]int, cur, nums []int, level int) {
	if level > len(nums) {
		return
	}

	tmp := make([]int, len(cur), len(cur))
	copy(tmp, cur)
	*ret = append(*ret, tmp)

	for i := level; i < len(nums); i++ {
		cur := append(cur, nums[i])
		subsetsWithDup(ret, cur, nums, i+1)
		cur = cur[:len(cur)-1]
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
	}
}
