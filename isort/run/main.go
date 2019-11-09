package main

import (
	"fmt"
	"github.com/panguncle/algorithm/isort"
)

func main() {
	test3WayQuickSort()
}

func testInsertionSort() {
	n := 100000
	arr := isort.GenerateRandomArray(n, 0, n)
	isort.TestSort("insert sort", isort.InsertionSort, arr)
}

func test1() {
	start, end := 2, 4
	fmt.Println((start + end) / 2)
	fmt.Println(start + (end-start)/2)
}

func testMergeSort() {
	n := 100000
	arr := isort.GenerateRandomArray(n, 0, n)
	isort.TestSort("merge sort", isort.MergeSort, arr)
}

func testMergeSort_2() {
	n := 10000000
	arr := isort.GenerateRandomArray(n, 0, n)
	arr2 := isort.CopyArr(arr)
	isort.TestSort("merge sort", isort.MergeSort, arr)
	isort.TestSort("merge sort 2", isort.MergeSort_2, arr2)
}

func testQuickSort() {
	n := 1000000
	arr := isort.GenerateRandomArray(n, 0, n)
	arr1 := isort.GenerateNearlyOrderedArray(n, 10)

	isort.TestSort("quick sort with random array", isort.QuickSort, arr)
	isort.TestSort("quick sort with nearly sorted array", isort.QuickSort, arr1)
}

func test3WayQuickSort() {
	n := 1000000
	arr := isort.GenerateRandomArray(n, 0, n)
	arr1 := isort.GenerateNearlyOrderedArray(n, 10)

	isort.TestSort("3 way quick sort with random array", isort.ThreeWayQuickSort, arr)
	isort.TestSort("3 way quick sort with nearly sorted array", isort.ThreeWayQuickSort, arr1)
}
