package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/panguncle/algorithm/2019/dp"
)

func main() {
	// testUniquePaths()
	fmt.Println(len(strconv.Itoa(time.UnixNano())))
}

func b(a map[int]bool) {
	fmt.Printf("%+v", a)
	a[2] = true
}

func testUniquePaths() {
	ret := dp.UniquePaths(7, 3)
	fmt.Println(ret)
}
