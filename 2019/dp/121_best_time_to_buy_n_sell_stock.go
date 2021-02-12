package dp

// BestTimeToBuyAndSellStock 121
func BestTimeToBuyAndSellStock(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// dp[i]: 今天卖掉这只股票的最大的收益
	// dp[i] = prices[i] - minV, minV是之前最小的价格
	// 在最小的价格买入, 在最大的价格卖出

	var (
		maxP int
		minV = prices[0]
	)

	for i := 1; i < len(prices); i++ {
		maxP = max(maxP, prices[i]-minV)
		minV = min(minV, prices[i])

	}

	return maxP
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
