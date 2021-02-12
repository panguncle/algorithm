package isort

/**
QuickSort:
每次找一个标定点(随机下标), 将数组划分为两部分, 左边<=标点定, 右边>=标定点
然后再对标定点的两侧分别进行相同的操作
*/
func QuickSort(list []int) []int {
	if len(list) == 0 {
		return list
	}

	quickSort(list, 0, len(list)-1)
	return list
}

// quickSort: 辅助函数, 选择标定点, 并对标定点两侧进行再次的排序
func quickSort(list []int, start, end int) {
	if start >= end {
		return
	}

	p := partition(list, start, end)
	quickSort(list, start, p-1)
	quickSort(list, p+1, end)
}

/**
partition, 选择一个标定点p, 并对list进行一番调整操作, 使得
list[start: p] < list[p] < list[p+1, end]
*/
func partition(list []int, start, end int) int {
	// 选择标定点,
	// 为了标定点选择更加均匀, 选择一个随机的点

	ri := RandomInt(start, end, true)
	list[start], list[ri] = list[ri], list[start]
	v := list[start]

	i, j := start+1, end
	for {
		for i <= end && list[i] < v {
			i++
		}

		for j >= start+1 && list[j] > v {
			j--
		}

		// 这个时候如果i > j 就break
		if i >= j {
			break
		}

		// 这个时候, i所指向的是比v大的数值, j指向的是小于v的数值
		// 交换一下, 然后各自往前推进一步
		list[i], list[j] = list[j], list[i]
		i++
		j--
	}

	// 出来了, 这个时候满足i>=j
	// 应该将标定点的值 归入 <= v的部分,
	// 将标定点和j交换一下
	list[start], list[j] = list[j], list[start]
	return j
}
