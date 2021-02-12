package dp

// HourseRobber 198
func HourseRobber(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}

	// dp[i]: 到第i间房子, 所获取的最大的利益
	// dp[i] = max(dp[i - 2] + nums[i], dp[i-1])
	// 偷取i-1个房子, 不偷当前这个房子, 或者不偷, i-1, 然后偷当前这一间
	var cur, pre1, pre2 int

	for i := 0; i < len(nums); i++ {
		cur = max(nums[i]+pre2, pre1)
		pre2 = pre1
		pre1 = cur
	}

	return cur
}
