package search

//Combination: leetcode:77:https://leetcode.com/problems/combinations/
func Combination(n, k int) [][]int {

	// 定义两个集合, 一个用于存储当前的集合, 一个用于存储, 结果
	var ret [][]int
	if k == 0 {
		return ret
	}

	var cur []int

	combination(&ret, cur, 1, n, k)
	return ret
}

func combination(ret *[][]int, cur []int, start, n, k int) {
	if len(cur) > k {
		return
	}

	if len(cur) == k {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*ret = append(*ret, tmp)
		return
	}

	for i := start; i <= n; i++ {
		cur = append(cur, i)
		combination(ret, cur, i+1, n, k-1) // DFS
		cur = cur[:len(cur)-1]             // backtracking
	}

}
