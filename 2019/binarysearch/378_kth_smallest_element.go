package binarysearch

// KthSmallestElement : leetcode.378
// Todo
func KthSmallestElement(matrix [][]int, k int) int {
	rowCount := len(matrix)
	low := matrix[0][0]
	high := matrix[rowCount-1][rowCount-1]
	var count int
	for low <= high {
		mid := low + (high-low)/2
		// 找出有多少个元素 <= mid
		count = 0
		for _, row := range matrix {
			count += rankKthInRow(row, mid)
		}

		// mid是第k小, val(第8小) > val(第5小)
		// 调整上界
		if count >= k {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}

func rankKthInRow(row []int, k int) int {

	low := 0
	high := len(row) - 1

	for low <= high {
		mid := low + (high-low)/2
		if row[mid] >= k {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}
