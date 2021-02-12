package tree

import "sort"

// FindFrequentTreeSum: 508
func FindFrequentTreeSum(root *TreeNode) []int {
	// Sum All Sub Tree, store in map key:sum value:count

	if root == nil {
		return nil
	}

	store := make(map[int]int)

	sumSubTree(root, store)

	//fmt.Printf("%+v\n", store)
	ret := findMaxValueFromMap(store)
	sort.Ints(ret)
	return ret
}

func sumSubTree(node *TreeNode, store map[int]int) int {
	if node == nil {
		return 0
	}

	valLeft := sumSubTree(node.Left, store)
	valRight := sumSubTree(node.Right, store)
	sum := valLeft + valRight + node.Val
	store[sum] += 1
	return sum
}

func findMaxValueFromMap(store map[int]int) []int {

	revStore := make(map[int][]int)

	var maxCount int
	for key, value := range store {
		if value < maxCount {
			continue
		}
		if _, exists := revStore[value]; !exists {
			revStore[value] = make([]int, 0)
		}
		revStore[value] = append(revStore[value], key)
		maxCount = value
	}

	return revStore[maxCount]
}

// FindFrequentTreeSum: 508
// Todo
func FindFrequentTreeSum_2(root *TreeNode) []int {
	// Sum All Sub Tree, store in map key:sum value:count

	if root == nil {
		return nil
	}

	var (
		mostCount int
		ret       []int
		store     = make(map[int]int)
	)

	sumSubTree_2(root, store, &mostCount, &ret)

	return ret
}

func sumSubTree_2(node *TreeNode, store map[int]int, mostCount *int, ret *[]int) int {
	if node == nil {
		return 0
	}

	valLeft := sumSubTree_2(node.Left, store, mostCount, ret)
	valRight := sumSubTree_2(node.Right, store, mostCount, ret)
	sum := valLeft + valRight + node.Val
	store[sum] += 1
	if store[sum] == *mostCount {
		*ret = append(*ret, sum)
	}
	if store[sum] > *mostCount {
		*mostCount = store[sum]
		tmp := make([]int, 0)
		tmp = append(tmp, sum)
		*ret = tmp
	}
	return sum
}
