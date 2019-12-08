package main

import (
	"log"

	. "github.com/panguncle/algorithm/helper"
	"github.com/panguncle/algorithm/twopoints"
)

func main() {
	testTwoSumII()
}

func testTwoSumII() {
	testcases := []struct {
		Input    []int
		Target   int
		Excepted []int
	}{
		{
			Input:    []int{1, 2, 3, 4, 5},
			Target:   2,
			Excepted: nil,
		},
		{
			Input:    []int{},
			Target:   1,
			Excepted: nil,
		},
		{
			Input:    []int{1},
			Target:   1,
			Excepted: nil,
		},
		{
			Input:    []int{2, 5, 8, 12},
			Target:   10,
			Excepted: []int{0, 2},
		},
	}

	for _, testcase := range testcases {
		ret := twopoints.TwoSumII(testcase.Input, testcase.Target)
		if IntSliceEqual(ret, testcase.Excepted) {
			log.Fatalf("testcase:%+v, excepted:%+v, got:%+v\n", testcase, testcase.Excepted, ret)
		}
	}

	log.Println("ok~~~")
}
