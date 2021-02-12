package isort

func ThreeWayQuickSort(list []int) []int {
	if len(list) == 0 {
		return list
	}

	threeWayQuickSort(list, 0, len(list)-1)
	return list
}

func threeWayQuickSort(list []int, start, end int) {
	if start >= end {
		return
	}

	lt, gt := threeWayPartition(list, start, end)

	threeWayQuickSort(list, start, lt)
	threeWayQuickSort(list, gt, end)
}

/**
threeWayPartition:
三路partition, 先取得一个标定点p, 其值为pv, 通过这个p 将list分成三部分
< pv, == pv, > pv, 得到两个标定点(lt, gt), 并对数组进行一番操作, 使得
*/
func threeWayPartition(list []int, start, end int) (int, int) {

	ri := RandomInt(start, end, true)
	list[start], list[ri] = list[ri], list[start]
	pv := list[start]
	/**
	list[start] -> list[lt] < pv
	list[i] -> list[gt-1] == pv
	list[gt] -> list[end] > pv
	*/
	lt := start // 设置为start, 最后让start 和 lt交换位置, 满足条件
	i := start + 1
	gt := end + 1 // 设置为end + 1来满足出事条件

	for i < gt {
		if list[i] > pv {
			list[i], list[gt-1] = list[gt-1], list[i]
			//i++ 换过来gt-1不一定就就 <v因此不需要移动
			gt--
		} else if list[i] < pv {
			list[i], list[lt+1] = list[lt+1], list[i]
			i++
			lt++
		} else {
			i++
		}
	}

	list[start], list[lt] = list[lt], list[start]

	return lt, gt
}
