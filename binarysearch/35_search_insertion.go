package binarysearch

// SearchInsert : leetcode.35
func SearchInsert(nums []int, target int) int {
	length := len(nums)
	low := 0
	high := length - 1

	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	// low > high && low = high + 1
	// low跟high最后一次重叠, low == mid == high 后往后移动一位
	return low
}
