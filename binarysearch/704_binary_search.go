package binarysearch

// Search : 704
func Search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	low, high := 0, length-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
