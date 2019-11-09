package isort

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateRandomArray(n, min, max int) []int {
	if n <= 0 {
		panic("n 必须大于0")
	}
	if min >= max {
		panic("max必须大于min")
	}
	rand.Seed(time.Now().Unix())

	// rand.Intn(n) 产生 [0~n]的数值
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		res = append(res, rand.Intn(max-min+1)+min)
	}
	return res
}

func GenerateNearlyOrderedArray(n, swapTimes int) []int {
	if n <= 0 {
		panic("n 必须大于0")
	}

	// rand.Intn(n) 产生 [0~n]的数值
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		res = append(res, i)
	}

	rand.Seed(time.Now().Unix())
	for i := 0; i < swapTimes; i++ {
		x := rand.Intn(n)
		y := rand.Intn(n)
		res[x], res[y] = res[y], res[x]
	}

	return res
}

func TestSort(sortName string, f func([]int) []int, arr []int) {
	fmt.Printf("start to test func: %s \n", sortName)
	fmt.Printf("sample num: %d\n", len(arr))
	startTime := time.Now()
	res := f(arr)
	endTime := time.Now()
	fmt.Printf("spent: %f s\n", endTime.Sub(startTime).Seconds())
	if !IsSorted(res) {
		fmt.Println("res: wrong")
		return
	}
	fmt.Println("res: true")
}

func IsSorted(arr []int) bool {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func CopyArr(arr []int) []int {
	l := len(arr)
	c := cap(arr)

	data := make([]int, l, c)
	copy(data, arr)
	return data
}

// randomInt, [start, end]/ [start, end)
func RandomInt(start, end int, includeEnd bool) int {
	if start >= end {
		panic("start must lt end")
	}

	v := end - start
	if includeEnd {
		v += 1
	}
	return rand.Intn(v) + start
}
