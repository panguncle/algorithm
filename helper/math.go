package helper

import "errors"

var (
	ErrListEmpty = errors.New("list must not be empty")
)

func MaxInt(list ...int) int {
	if len(list) == 0 {
		panic(ErrListEmpty)
	}

	max := list[0]

	for i := 1; i < len(list); i++ {
		if list[i] > max {
			max = list[i]
		}
	}

	return max
}

func MinInt(list ...int) int {
	if len(list) == 0 {
		panic(ErrListEmpty)
	}

	min := list[0]

	for i := 1; i < len(list); i++ {
		if list[i] < min {
			min = list[i]
		}
	}

	return min
}
