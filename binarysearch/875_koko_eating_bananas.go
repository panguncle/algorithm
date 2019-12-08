package binarysearch

// KokoEatingBananas : leetcode.875
func KokoEatingBananas(piles []int, H int) int {
	// 由题目的意思可知, 每小时吃一个piles, 因此 H >= len(piles)
	// 使用二分搜索, 搜索的下界默认是1,
	// 上界是piles中最大的那个

	if H < len(piles) || len(piles) == 0 || H <= 0 {
		return -1
	}

	max := piles[0]
	for _, v := range piles {
		if v > max {
			max = v
		}
	}

	low, high := 1, max

	var gm = func(rate int) int {
		var count int

		for _, v := range piles {
			count += v / rate
			if v%rate != 0 {
				count++
			}
		}
		return count
	}

	for low <= high {
		mid := low + (high-low)/2
		if gm(mid) <= H {
			// 需要的时间 <= 给定的时间, 说明吃快了
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}
