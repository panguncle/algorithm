package main

import (
	"fmt"
	"log"
	"strings"

	. "github.com/panguncle/algorithm/2019/helper"
	"github.com/panguncle/algorithm/2019/tree"
	"github.com/panguncle/algorithm/2019/tree/bst"
)

func main() {
	testMinAbsDiffInBST()
}

func testInOrderTraversal() {
	arr := []*int{
		IntPtr(1), IntPtr(2), IntPtr(2), nil, IntPtr(4),
	}

	node := tree.BuildBinaryTreeFromArray(arr)

	ret := tree.InOrderTraversal(node)

	arr1 := []int{2, 4, 1, 2}

	fmt.Printf("ok:%t\n", IntSliceEqual(arr1, ret))
}

func testNAryTreePreOrder() {

}

func testIsSameTree() {
	arr := []*int{
		IntPtr(1), IntPtr(2), IntPtr(2), nil, IntPtr(4),
	}

	treeA := tree.BuildBinaryTreeFromArray(arr)
	treeB := tree.BuildBinaryTreeFromArray(arr)

	ret := tree.IsSameTree(treeA, treeB)

	fmt.Printf("isSame:%t\n", ret)

	arr2 := []*int{
		IntPtr(1), IntPtr(2), nil, IntPtr(3),
	}
	treeC := tree.BuildBinaryTreeFromArray(arr2)
	ret = tree.IsSameTree(treeA, treeC)

	fmt.Printf("isSame:%t\n", ret)
}

func testIsSymmetric() {
	arr := []*int{
		IntPtr(1), IntPtr(2), IntPtr(2), nil, IntPtr(4), IntPtr(4), nil,
	}
	treeA := tree.BuildBinaryTreeFromArray(arr)

	ret := tree.IsSymmetric(treeA)

	fmt.Printf("is_symmetric:%t\n", ret)
}

func testMaxDepthOfBinaryTree() {
	arr := []*int{
		IntPtr(1), IntPtr(2), IntPtr(2), nil, IntPtr(4), IntPtr(4), nil,
		nil, nil, IntPtr(1),
	}
	treeA := tree.BuildBinaryTreeFromArray(arr)

	ret := tree.MaxDepthOfBinaryTree(treeA)
	fmt.Printf("max_depth:%d\n", ret)
	ret = tree.MaxDepthOfBinaryTree_2(treeA)
	fmt.Printf("max_depth:%d\n", ret)
}

func testIsUnivalTree() {
	testcases := []struct {
		Source   []*int
		Excepted bool
	}{
		{
			Source:   []*int{},
			Excepted: true,
		},
		{
			Source:   []*int{IntPtr(1)},
			Excepted: true,
		},
		{
			Source:   []*int{IntPtr(3), IntPtr(3), IntPtr(3), nil, IntPtr(3), nil, nil},
			Excepted: true,
		},
		{
			Source:   []*int{IntPtr(3), IntPtr(3), IntPtr(4), nil, IntPtr(3), nil, nil},
			Excepted: false,
		},
	}

	for _, testcase := range testcases {
		sourceTree := tree.BuildBinaryTreeFromArray(testcase.Source)
		ret := tree.IsUnivalTree(sourceTree)
		if ret != testcase.Excepted {
			log.Fatalf("testcase: %#v, excepted:%t, got:%t\n", testcase, testcase.Excepted, ret)
		}
	}

	fmt.Printf("ok!\n")
}

func testLevelOrderBottomUp() {
	arr := []*int{
		IntPtr(1), IntPtr(2), IntPtr(2), nil, IntPtr(4), IntPtr(4), nil,
		nil, nil, IntPtr(1),
	}
	treeA := tree.BuildBinaryTreeFromArray(arr)

	ret := tree.LevelOrderBottom(treeA)

	fmt.Println(ret)
}

func testFindFrequentTreeSum() {
	arr := []*int{
		IntPtr(5), IntPtr(2), IntPtr(-5),
	}
	treeA := tree.BuildBinaryTreeFromArray(arr)

	ret := tree.FindFrequentTreeSum(treeA)

	fmt.Printf("%+v\n", ret)
	arrB := []*int{
		IntPtr(5), IntPtr(2), IntPtr(-3),
	}

	treeB := tree.BuildBinaryTreeFromArray(arrB)
	ret = tree.FindFrequentTreeSum(treeB)

	fmt.Printf("%+v\n", ret)

}

func testSumRootToLeafNumbers() {
	arr := []*int{
		IntPtr(4), IntPtr(9), IntPtr(0), IntPtr(5), IntPtr(1),
	}
	treeA := tree.BuildBinaryTreeFromArray(arr)

	ret := tree.SumRootToLeafNumbers(treeA)

	fmt.Printf("%+v\n", ret)

	arrB := []*int{
		IntPtr(0), IntPtr(0), IntPtr(4), IntPtr(9), IntPtr(0), IntPtr(5), IntPtr(1),
	}
	treeB := tree.BuildBinaryTreeFromArray(arrB)

	ret = tree.SumRootToLeafNumbers(treeB)

	fmt.Printf("%+v\n", ret)
}

func testTrimPrefix() {
	a := "00001"
	fmt.Println(strings.TrimPrefix(a, "0"))
	fmt.Println(strings.TrimLeft(a, "0"))
}

func testBinaryTreePaths() {
	arr := []*int{
		IntPtr(4), IntPtr(9), IntPtr(0), IntPtr(5), IntPtr(1),
	}
	treeA := tree.BuildBinaryTreeFromArray(arr)

	ret := tree.BinaryTreePaths(treeA)

	fmt.Printf("%+v\n", ret)

	arrB := []*int{
		IntPtr(0), IntPtr(0), IntPtr(4), IntPtr(9), IntPtr(0), IntPtr(5), IntPtr(1),
	}
	treeB := tree.BuildBinaryTreeFromArray(arrB)

	ret = tree.BinaryTreePaths(treeB)

	fmt.Printf("%+v\n", ret)
}

func testIsValidBST() {
	treeA := tree.BuildBinaryTreeFromArray([]*int{
		IntPtr(2), IntPtr(3), IntPtr(1),
	})

	treeB := tree.BuildBinaryTreeFromArray([]*int{
		IntPtr(3), IntPtr(2), IntPtr(4),
		IntPtr(1), IntPtr(4),
		IntPtr(-2), IntPtr(5),
	})

	treeC := tree.BuildBinaryTreeFromArray([]*int{
		IntPtr(6), IntPtr(2), IntPtr(10),
		nil, IntPtr(4), IntPtr(7), IntPtr(20),
	})

	ret := bst.IsValidBST(treeA)
	if ret {
		log.Fatalf("excepted:%#v is not a valid bst, but got true", treeA)
	}

	ret = bst.IsValidBST(treeB)
	if ret {
		log.Fatalf("excepted:%#v is not a valid bst, but got true", treeB)
	}

	ret = bst.IsValidBST(treeC)
	if !ret {
		log.Fatalf("excepted:%#v is a valid bst, but got false", treeA)
	}

	fmt.Println("ok!")
}

func testMinAbsDiffInBST() {
	treeC := tree.BuildBinaryTreeFromArray([]*int{
		IntPtr(6), IntPtr(2), IntPtr(10),
		nil, IntPtr(4), IntPtr(7), IntPtr(20),
	})
	fmt.Println(bst.MinAbsDiffInBST_2(treeC))
	//[236,104,701,null,227,null,911]
	treeD := tree.BuildBinaryTreeFromArray([]*int{
		IntPtr(236), IntPtr(104), IntPtr(701), nil,
		IntPtr(227), nil, IntPtr(911),
	})
	fmt.Println(bst.MinAbsDiffInBST_2(treeD))
}
