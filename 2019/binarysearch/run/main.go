package main

import (
	"fmt"
	"log"

	"github.com/panguncle/algorithm/2019/binarysearch"
	. "github.com/panguncle/algorithm/2019/helper"
)

func main() {
	// testSearchInRotateArrayII()
	testSearch2DMatrix()
}

func testSearchInsertion() {
	arr := []int{1, 3, 5, 6}
	target := 7
	ret := binarysearch.SearchInsert(arr, target)
	fmt.Println(ret)
}

func testSqrtX() {
	testcases := []struct {
		Input    int
		Excepted int
	}{
		{
			Input:    0,
			Excepted: 0,
		},
		{
			Input:    1,
			Excepted: 1,
		},
		{
			Input:    8,
			Excepted: 2,
		},
	}

	for _, testcase := range testcases {
		ret := binarysearch.SqrtX(testcase.Input)
		if ret != testcase.Excepted {
			log.Fatalf("testcase: %+v, excepted:%d, but got:%d\n", testcase, testcase.Excepted, testcase.Input)
		}
	}

	log.Println("OK")
}

func testKthSmallestElement() {
	testcases := []struct {
		Matrix   [][]int
		K        int
		Excepted int
	}{
		{
			Matrix: [][]int{
				{1, 5, 9}, {10, 11, 13}, {12, 13, 15},
			},
			K:        8,
			Excepted: 13,
		},
	}

	for _, testcase := range testcases {
		ret := binarysearch.KthSmallestElement(testcase.Matrix, testcase.K)

		if ret != testcase.Excepted {
			log.Fatalf("testcase:%+v, excepted:%d, got:%d\n", testcase, testcase.Excepted, ret)
		}
	}

	fmt.Println("ok")
}

func testFindFirstNLastPosElemInSortedArray() {
	testcases := []struct {
		Input    []int
		Target   int
		Excepted []int
	}{
		{
			Input:    []int{1, 2, 3, 3, 4, 5, 6},
			Target:   3,
			Excepted: []int{2, 3},
		},
		{
			Input:    []int{1, 2, 3, 4, 5},
			Target:   3,
			Excepted: []int{2, 2},
		},
		{
			Input:    []int{1, 2, 3, 4, 5},
			Target:   6,
			Excepted: []int{-1, -1},
		},
	}

	for _, testcase := range testcases {
		ret := binarysearch.FindFirstNLastPosElemInSortedArray(testcase.Input, testcase.Target)
		if !IntSliceEqual(ret, testcase.Excepted) {
			log.Fatalf("testcase:%+v, excepted:%+v, got:%+v\n", testcase, testcase.Excepted, ret)
		}
	}

	log.Println("ok~")
}

func testBinSearch() {
	testcases := []struct {
		Input    []int
		Target   int
		Excepted int
	}{
		{
			Input:    []int{},
			Target:   5,
			Excepted: -1,
		},
		{
			Input:    []int{1, 2, 3, 6},
			Target:   4,
			Excepted: -1,
		},
		{
			Input:    []int{1},
			Target:   1,
			Excepted: 0,
		},
		{
			Input:    []int{1, 3, 5, 7, 9},
			Target:   5,
			Excepted: 2,
		},
	}

	for _, testcase := range testcases {
		ret := binarysearch.Search(testcase.Input, testcase.Target)
		if ret != testcase.Excepted {
			log.Fatalf("testcase:%+v, excepted:%+v, got:%+v\n", testcase, testcase.Excepted, ret)
		}
	}

	log.Println("ok~~")
}

func testSearchInRotateArrayII() {
	testcases := []struct {
		Input    []int
		Target   int
		Excepted bool
	}{
		{
			Input:    []int{1, 3, 4, 7, 9},
			Target:   5,
			Excepted: false,
		},
		{
			Input:    []int{1, 3, 2, 3, 6},
			Target:   2,
			Excepted: true,
		},
		{
			Input:    []int{},
			Target:   1,
			Excepted: false,
		},
		{
			Input:    []int{2, 5, 6, 0, 0, 1, 2},
			Target:   0,
			Excepted: true,
		},
	}

	for _, testcase := range testcases {
		ret := binarysearch.SearchInRotatedArrayII(testcase.Input, testcase.Target)
		if ret != testcase.Excepted {
			log.Fatalf("testcase:%+v, excepted:%+v, got:%+v\n", testcase, testcase.Excepted, ret)
		}
	}

	log.Println("ok~~")
}

func testSearch2DMatrix() {
	testcases := []struct {
		Input    [][]int
		Target   int
		Excepted bool
	}{
		{
			Input:    [][]int{},
			Target:   5,
			Excepted: false,
		},
		{
			Input: [][]int{
				{1, 2, 3},
			},
			Target:   4,
			Excepted: false,
		},
		{
			Input: [][]int{
				{1, 3, 5},
				{7, 12, 15},
				{20, 30, 50},
			},
			Target:   12,
			Excepted: true,
		},
	}

	for _, testcase := range testcases {
		ret := binarysearch.Search2DMatrix(testcase.Input, testcase.Target)
		if ret != testcase.Excepted {
			log.Fatalf("testcase:%+v, excepted:%t, got:%t\n", testcase.Input, testcase.Excepted, ret)
		}
	}

	log.Println("ok~~")
}
