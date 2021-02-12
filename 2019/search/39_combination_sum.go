package search

import (
	"sort"
)

// CombinationSum: leetcode:39.
func CombinationSum(candidates []int, target int) [][]int {

	var ret [][]int

	if target == 0 {
		return ret
	}

	var cur []int
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	combinationSum(&ret, cur, candidates, target, 0)
	return ret

}

func combinationSum(ret *[][]int, cur []int, candidates []int, target, start int) {
	if target < 0 {
		return
	}

	if target == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*ret = append(*ret, tmp)
		return
	}

	for i := start; i < len(candidates); i++ {
		cur = append(cur, candidates[i])
		combinationSum(ret, cur, candidates, target-candidates[i], i)
		cur = cur[:len(cur)-1]
	}
}
