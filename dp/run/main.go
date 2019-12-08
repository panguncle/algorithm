package main

import (
	"fmt"

	"github.com/panguncle/algorithm/dp"
)

func main() {
	testUniquePaths()
}

func b(a map[int]bool) {
	fmt.Printf("%+v", a)
	a[2] = true
}

func testUniquePaths() {
	ret := dp.UniquePaths(7, 3)
	fmt.Println(ret)
}
