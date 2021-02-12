package twopoints

// SquaresOfASortedArray :977
func SquaresOfASortedArray(A []int) []int {
	length := len(A)
	if length == 0 {
		return nil
	}

	ret := make([]int, length)

	l, r, cur := 0, length-1, length-1

	for cur >= 0 {
		a := A[l] * A[l]
		b := A[r] * A[r]
		if a >= b {
			ret[cur] = a
			l++
		} else {
			ret[cur] = b
			r--
		}
		cur--
	}

	return ret
}
