package binarysearch

// FindFirstNLastPosElemInSortedArray :
// Todo
func FindFirstNLastPosElemInSortedArray(nums []int, target int) []int {
	length := len(nums)
	if length == 0 {
		return []int{-1, -1}
	}

	lower := lowerBound(nums, target)
	if lower < 0 || lower >= length || nums[lower] != target {
		return []int{-1, -1}
	}

	upper := upperBound(nums, target)

	return []int{lower, upper - 1}
}

func upperBound(nums []int, target int) int {
	length := len(nums)

	low, high := 0, length-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}

func lowerBound(nums []int, target int) int {
	length := len(nums)

	low, high := 0, length-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}
