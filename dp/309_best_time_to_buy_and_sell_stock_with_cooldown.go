package dp

import (
	"math"
)

/*
 BestTimeToBuyAndSellStockWithCoolDown 309
 hold[i] = max(hold[i-1], rest[i-1] - prices[i]) 昨天持有/今天买入
 sold[i] = hold[i - 1] + prices[i] // 昨天持有, 今天卖出
 rest[i] max(rest[i-1], sold[i-1]) // 今天无作为, 或者昨天卖出
 init :rest[-] = sold[0] = 0, hold[0] == -oo
 ans: max(rest[i], sold[i])
 time: O(n)
 Space: O(n) -> O(1)

 // Todo
*/
func BestTimeToBuyAndSellStockWithCoolDown(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var sold, rest int
	hold := math.MinInt64

	for _, price := range prices {
		prevHold := sold
		sold = hold + price
		hold = max(hold, rest-price)
		rest = max(rest, prevHold)
	}
	return max(rest, sold)
}
