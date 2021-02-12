package binarysearch

// Search2DMatrix :74.
func Search2DMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	for _, row := range matrix {
		if search2DMatrixRow(row, target) {
			return true
		}
	}

	return false
}

func search2DMatrixRow(row []int, target int) bool {
	length := len(row)

	// 剪枝, 看看是否有必要在这个rows里面进行搜索
	if row[0] > target || row[length-1] < target {
		return false
	}

	low, high := 0, length-1
	for low <= high {
		mid := low + (high-low)/2
		if row[mid] == target {
			return true
		} else if row[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
