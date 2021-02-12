package twopoints

// TwoSumII : 167
func TwoSumII(numbers []int, target int) []int {
	length := len(numbers)
	if length < 2 {
		return nil
	}

	l, r := 0, length-1

	for l < r && r >= 0 && l <= length-1 {
		v := numbers[l] + numbers[r]
		if v == target {
			return []int{l, r}
		}

		if v > target {
			r--
		} else {
			l++
		}
	}

	return nil
}
