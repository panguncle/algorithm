package helper

// InSliceEqual
func IntSliceEqual(arr1, arr2 []int) bool {
	lenA, lenB := len(arr1), len(arr2)
	if lenA == 0 && lenB == 0 {
		return true
	}

	if lenA != lenB {
		return false
	}

	for i, v1 := range arr1 {
		if v1 != arr2[i] {
			return false
		}
	}

	return true
}
