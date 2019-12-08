package binarysearch

// SearchInRotatedArrayII : 81
// Todo
func SearchInRotatedArrayII(nums []int, target int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}

	return partitonNFind(nums, 0, length-1, target)
}

func partitonNFind(nums []int, start, end, target int) bool {
	if start > end {
		return false
	}

	pv := nums[start]
	if pv == target {
		return true
	}

	pfidx := partiton(nums, start, end)
	if pv > target {
		return partitonNFind(nums, start, pfidx-1, target)
	}

	return partitonNFind(nums, pfidx+1, end, target)
}

func partiton(nums []int, start, end int) int {

	v := nums[start]

	i, j := start+1, end
	for {
		for i <= end && nums[i] < v {
			i++
		}

		for j >= start+1 && nums[j] > v {
			j--
		}

		// 这个时候如果i > j 就break
		if i >= j {
			break
		}

		// 这个时候, i所指向的是比v大的数值, j指向的是小于v的数值
		// 交换一下, 然后各自往前推进一步
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	// 出来了, 这个时候满足i>=j
	// 应该将标定点的值 归入 <= v的部分,
	// 将标定点和j交换一下
	nums[start], nums[j] = nums[j], nums[start]
	return j
}
