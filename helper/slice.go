package helper

func IntSliceEqual(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i, v1 := range arr1 {
		if v1 != arr2[i] {
			return false
		}
	}

	return true
}
