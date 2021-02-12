package dp

/*
MaxProductSubArray : 152
dpMax[i] = max(dpMax[i - 1] * nums[i], dpMin[i-1] * nums[i], nums[i])
dpMin[i] = min(dpMin[i-1] * nums[i], dpMin[i - 1] * nums[i], nums[i])

max, min 不一定是正负数, 只要记下就好, 最小的可正可负, 但是有负数肯定更小, 所以记下
*/
func MaxProductSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum, minSum, max := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		p1 := maxSum * nums[i]
		p2 := minSum * nums[i]
		maxSum, minSum = maxMinIntn(p1, p2, nums[i])
		if maxSum > max {
			max = maxSum
		}
	}

	return maxSum
}

func maxMinIntn(ints ...int) (int, int) {
	max := ints[0]
	min := ints[0]
	for i := 1; i < len(ints); i++ {
		if ints[i] > max {
			max = ints[i]
		}
		if ints[i] < min {
			min = ints[i]
		}
	}
	return max, min
}
