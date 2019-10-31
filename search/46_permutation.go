package search

// Permutation: leetcode.46
func Permutation(nums []int) [][]int {
	var ret [][]int
	if len(nums) == 0 {
		return ret
	}

	var (
		cur     []int
		visited = make([]bool, len(nums))
	)
	permutation(&ret, cur, nums, visited)
	return ret
}

// 组合, 无需使用level来记录从哪一位开始, 但是需要记录哪些已经用过了
func permutation(ret *[][]int, cur []int, nums []int, visited []bool) {
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
			permutation(ret, cur, nums, visited)
			visited[i] = false
			cur = cur[0 : len(cur)-1]
		}
	}

}
