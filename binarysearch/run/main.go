package main

import (
	"fmt"
	"github.com/panguncle/algorithm/binarysearch"
)

func main() {
	testSearchInsertion()
}

func testSearchInsertion() {
	arr := []int{1,3,5,6}
	target := 7
	ret := binarysearch.SearchInsert(arr, target)
	fmt.Println(ret)
}