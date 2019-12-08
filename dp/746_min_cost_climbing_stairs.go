package dp

/*
MinCosClimbingStairs
@see https://www.youtube.com/watch?v=v3WqNLmmBdk
## Recursion + Memorization
Time: O(n)
Space: O(n)
dp(n) := min cost to climb to n-th step
dp(n) = min(dp(n-1) + cost[n-1], dp(n-2) + cost[n-2])
ans = dp[n]
走到第n阶台阶, 因此0 不在数组内, 最后是dp[len(cost)]

### DP
Time: O(n)
Space: O(n) -> O(1)
dp(n): min cost before leaving n-th step
dp(n) = min(dp(n-1), dp(n-2) + cost[n])
-> 在前一步可以直接跳过他, 无需交钱;
在两步前最多走到这里, 要离开这里就需要交钱
*/
func MinCosClimbingStairs(cost []int) int {
	if len(cost) == 0 || len(cost) == 1 {
		return 0
	}
	dp := make([]int, len(cost)+1)
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[len(cost)]
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}
