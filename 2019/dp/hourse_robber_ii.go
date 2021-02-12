package dp

// HourseRobberII 213
func HourseRobberII(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	return max(toRobII(nums, 0, len(nums)-1), toRobII(nums, 1, len(nums)))
}

func toRobII(nums []int, start, end int) int {

	prev1, prev2 := nums[start], nums[start+1]
	maxV := max(prev1, prev2)
	for i := start + 2; i < end; i++ {
		cur := max(prev1+nums[i], prev2)
		maxV = max(cur, maxV)
		prev1 = prev2
		prev2 = cur
	}
	return maxV
}
