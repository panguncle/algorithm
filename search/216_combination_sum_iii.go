package search

// CombinationSumIII: leetcode.216
func CombinationSumIII(k, n int) [][]int {
	// k number add to n
	var ret [][]int
	if k == 0 || n == 0 || n < k {
		return ret
	}

	var res []int
	combinationSumIII(&ret, res, n, k, 1)
	return ret
}

func combinationSumIII(ret *[][]int, cur []int, n, k, level int) {
	if len(cur) > k {
		return
	}

	if len(cur) == k && sumSlice(cur) == n {
		tmp := make([]int, len(cur), len(cur))
		copy(tmp, cur)
		*ret = append(*ret, tmp)
		return
	}

	for i := level; i <= 9; i++ {
		cur = append(cur, i)
		combinationSumIII(ret, cur, n, k, i+1)
		cur = cur[:len(cur)-1]
	}
}

func sumSlice(s []int) int {
	var ret int
	for i := 0; i < len(s); i++ {
		ret += s[i]
	}

	return ret
}
