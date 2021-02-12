package dp

// UniquePaths : 62
func UniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 1
	}

	mem := make([][]int, 0)
	for i := 0; i < m; i++ {
		mem = append(mem, make([]int, n, n))
	}

	for i := 0; i < n; i++ {
		mem[0][i] = 1
	}

	for j := 0; j < m; j++ {
		mem[j][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			mem[i][j] = mem[i-1][j] + mem[i][j-1]
		}
	}

	return mem[m-1][n-1]
}
