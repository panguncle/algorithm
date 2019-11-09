package isort

func MergeSort(list []int) []int {

	if len(list) == 0 {
		return list
	}

	mergeSort(list, 0, len(list)-1)
	return list
}

func mergeSort(list []int, start, end int) {
	if start >= end {
		return
	}
	//mid := (start + end) / 2
	mid := start + (end-start)/2
	mergeSort(list, start, mid)
	mergeSort(list, mid+1, end)

	if list[mid] > list[mid+1] {
		merge(list, start, mid, end)
	}
}

func merge(list []int, start, mid, end int) {

	aux := make([]int, end-start+1)

	for i := start; i <= end; i++ {
		aux[i-start] = list[i]
	}

	j, k := start, mid+1
	for i := start; i <= end; i++ {
		if j > mid {
			list[i] = aux[k-start]
			k++
		} else if k > end {
			list[i] = aux[j-start]
			j++
		} else if aux[j-start] <= aux[k-start] {
			list[i] = aux[j-start]
			j++
		} else {
			list[i] = aux[k-start]
			k++
		}
	}
}

func MergeSort_2(list []int) []int {
	if len(list) == 0 {
		return list
	}

	mergeSort_2(list, 0, len(list)-1)
	return list
}

func mergeSort_2(list []int, start, end int) {
	if start >= end {
		return
	}

	mid := start + (end-start)/2
	mergeSort_2(list, start, mid)
	mergeSort_2(list, mid+1, end)

	if list[mid] > list[mid+1] {
		merge_2(list, start, mid, end)
	}
}

func merge_2(list []int, start, mid, end int) {

	// 将需要归并的部分进行归并
	tmp := make([]int, end-start+1)
	copy(tmp, list[start:end+1])

	// 使用两个指针记录两个需要归并的数组的起始位置
	j, k := start, mid+1
	for i := start; i <= end; i++ {
		if j > mid {
			copy(list[i:end+1], tmp[k-start:])
			break
		} else if k > end {
			copy(list[i:end+1], tmp[j-start:mid+1-start])
			break
		} else if tmp[j-start] > tmp[k-start] {
			list[i] = tmp[k-start]
			k++
		} else {
			list[i] = tmp[j-start]
			j++
		}
	}
}
