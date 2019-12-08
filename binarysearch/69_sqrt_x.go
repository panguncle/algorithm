package binarysearch

// SqrtX : leetcode. 69
func SqrtX(x int) int {
	if x == 0 {
		return x
	}

	// 等价于 有 sorted list [1, x],
	// 求 最小的 i >= target
	// 即是就lower_bound
	low := 1
	high := x

	for low <= high {
		mid := low + (high-low)/2
		pow := mid * mid
		if pow == x {
			return mid
		}
		if pow > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	// 这里要找的是第一个 i * i < target的数值
	// 此处找到的是第一个 i * i > target的数值
	// 因此 - 1
	return low - 1

}
