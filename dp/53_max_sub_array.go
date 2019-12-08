package dp

/**
MaxSubArray : 53
一唯的dp, 使用
dp[i] => 使用到index 为i 的数时, 最大的数
dp[i] = max(dp[i-1] + nums[i], nums[i])

*/
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 最少要包含一格元素
	sum, max := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		// 前置的数值 > 0 才加上来
		sum = maxInt(sum+nums[i], nums[i])
		if sum > max {
			max = sum
		}
	}

	return max
}

func maxInt(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
