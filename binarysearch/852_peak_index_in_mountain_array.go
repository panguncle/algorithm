package binarysearch

// PeakIndexInMountainArray : 852
func PeakIndexInMountainArray(A []int) int {
	// 找到一个数值, 这个数值的左边和右边都比他小
	length := len(A)
	if length == 0 {
		return -1
	}

	low, high := 0, length-1
	for low <= high {
		mid := low + (high-low)/2

		if A[mid] > A[mid-1] && A[mid] > A[mid+1] {
			return mid
		} else if A[mid] > A[mid-1] && A[mid] < A[mid+1] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
