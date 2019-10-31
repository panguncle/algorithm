package main

import (
	"fmt"
	linkedlist "github.com/panguncle/algorithm/linked_list"
)

func main() {
	testRemoveDupFromSortedList()
}

func testLinkedListCycle() {
	node := linkedlist.BuildLinkedList([]int{1, 2, 3, 4})
	ret := linkedlist.LinkedListHasCycle(node)
	fmt.Println("false == ", ret)

	node = linkedlist.BuildCycledLinkedList([]int{1, 2, 3, 4})
	ret = linkedlist.LinkedListHasCycle(node)
	fmt.Println("true == ", ret)
}

func testLinkedListCycleII() {
	node := linkedlist.BuildCycleLinkedListII([]int{1, 2, 3, 4}, 2)
	ret := linkedlist.LinkedListCycleII(node)
	fmt.Printf("%+v\n", ret)
}

func testRemoveDupFromSortedList() {
	head := linkedlist.BuildLinkedList([]int{1, 2, 2, 3, 4})

	linkedlist.PrintfLinkedListNotCycle(head)
	node := linkedlist.RemoveDupFromSortedList(head)

	linkedlist.PrintfLinkedListNotCycle(node)
}
