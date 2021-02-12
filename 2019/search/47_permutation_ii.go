package search

import "sort"

// PermuteUnique: leetcode.47
func PermuteUnique(nums []int) [][]int {

	var ret [][]int
	if len(nums) == 0 {
		return ret
	}

	var (
		cur     []int
		visited = make([]bool, len(nums))
	)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	permuteUnique(&ret, cur, nums, visited)
	return ret
}

func permuteUnique(ret *[][]int, cur, nums []int, visited []bool) {
	if len(cur) == len(nums) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*ret = append(*ret, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !visited[i] {
			visited[i] = true
			cur = append(cur, nums[i])
			permuteUnique(ret, cur, nums, visited)
			visited[i] = false
			cur = cur[:len(cur)-1]
			for i+1 < len(nums) && nums[i] == nums[i+1] {
				i++
			}
		}
	}
}
