package search

import (
	"fmt"
	"sort"
)

// CombinationSum2: leetcode:40.
func CombinationSum2(candidates []int, target int) [][]int {
	var ret [][]int
	if target == 0 {
		return ret
	}

	var cur []int
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	fmt.Println(candidates)
	combinationSum2(&ret, cur, 0, candidates, target)
	return ret
}

func combinationSum2(ret *[][]int, cur []int, level int, candidates []int, target int) {
	if target < 0 {
		return
	}

	if target == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*ret = append(*ret, tmp)
		return
	}

	for i := level; i < len(candidates); i++ {
		cur = append(cur, candidates[i])
		combinationSum2(ret, cur, i+1, candidates, target-candidates[i])
		cur = cur[:len(cur)-1]
		for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
			i++
		}
	}
}
