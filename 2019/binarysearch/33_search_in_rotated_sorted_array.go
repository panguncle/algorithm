package binarysearch

/*
SearchInRotatedSortedArray 33
二分查找变种:
if nums[mid] == targer {
	return
}
if nums[low] <= nums[mid] {
	// low 到 mid是有序的
}

if nums[mid] <= nums[high] {
	// mid -> high 有序
}
*/
func SearchInRotatedSortedArray(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	var (
		low, high, mid = 0, len(nums) - 1, 0
	)

	for low <= high {
		mid = low + (high-low)/2
		if nums[mid] == target {
			return mid
		}

		// [low, mid] 有序
		if nums[low] < nums[mid] {
			if nums[low] <= target && target <= nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			// [mid + 1, high] 有序
			// mid + 1可能越界
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}
