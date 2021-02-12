package twopoints

import (
	. "github.com/panguncle/algorithm/2019/helper"
)

// ContainerWithMostWater : 11
func ContainerWithMostWater(height []int) int {
	length := len(height)
	if length == 0 {
		return 0
	}

	l, r := 0, length-1
	var max int
	for l < r && l <= length-1 && r >= 0 {

		x := r - l
		y := MinInt(height[r], height[l])
		area := x * y
		if area > max {
			max = area
		}

		if height[l] >= height[r] {
			r--
		} else {
			l++
		}
	}

	return max
}
