package isort

/**
InsertionSort: 插入排序
前置的元素已经排好序, 每次从最后一个排好序的元素的下一位开始, 往前比较, 将元素插入相应的位置
*/
func InsertionSort(list []int) []int {
	length := len(list)
	if length == 0 {
		return list
	}

	// 从1开始, 因为0已经是有序的
	for i := 1; i < length; i++ {
		// 从最后一个排好序的元素的下一位开始, 往前比较
		j := i
		v := list[i]
		// 提前终止, 只要是前面的值不能大过当前值
		for ; j > 0 && list[j-1] > v; j-- {
			list[j] = list[j-1]
		}
		list[j] = v
	}

	return list

}
