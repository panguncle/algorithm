package dp

// ClimbingStairs :70
func ClimbingStairs(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}
	// 爬一格有一种方法, 2格有2种方法
	prev1, prev2 := 1, 2

	// 从第3格开始, 一直到最后一格
	for i := 3; i <= n; i++ {
		cur := prev1 + prev2
		prev1 = prev2
		prev2 = cur
	}

	return prev2
}
