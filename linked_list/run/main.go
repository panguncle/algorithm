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

func testRemoveDupFromSortedListII() {
	head := linkedlist.BuildLinkedList([]int{1, 2, 3, 3, 4, 4, 5})
	linkedlist.PrintfLinkedListNotCycle(head)
	node := linkedlist.RemoveDupFromSortedListII(head)
	linkedlist.PrintfLinkedListNotCycle(node)
}

func testIntersectionOfTwoLinkedList() {
	head1, head2 := linkedlist.BuildIntersectionLinkedList([]int{1, 2, 3}, []int{7, 8, 9}, []int{4, 5, 6})
	linkedlist.PrintfLinkedListNotCycle(head1)
	linkedlist.PrintfLinkedListNotCycle(head2)

	node := linkedlist.IntersectionOfTwoLinkedList(head1, head2)
	fmt.Printf("%+v\n", node)
}

func testMergeTwoSortedList() {
	head1 := linkedlist.BuildLinkedList([]int{1, 2, 4})
	head2 := linkedlist.BuildLinkedList([]int{1, 3, 4})
	linkedlist.PrintfLinkedListNotCycle(head1)
	linkedlist.PrintfLinkedListNotCycle(head2)

	node := linkedlist.MergeTwoSortedList(head1, head2)
	linkedlist.PrintfLinkedListNotCycle(node)
}

func testMergeTwoSortedList2() {
	head1 := linkedlist.BuildLinkedList([]int{1, 2, 4})
	head2 := linkedlist.BuildLinkedList([]int{1, 3, 4})
	linkedlist.PrintfLinkedListNotCycle(head1)
	linkedlist.PrintfLinkedListNotCycle(head2)

	node := linkedlist.MergeTwoSortedList_2(head1, head2)
	linkedlist.PrintfLinkedListNotCycle(node)
}

func testMergeKSortedList() {
	head1 := linkedlist.BuildLinkedList([]int{1, 2, 4})
	head2 := linkedlist.BuildLinkedList([]int{1, 3, 4})
	head3 := linkedlist.BuildLinkedList([]int{4, 5, 6})
	linkedlist.PrintfLinkedListNotCycle(head1)
	linkedlist.PrintfLinkedListNotCycle(head2)
	linkedlist.PrintfLinkedListNotCycle(head3)
	node := linkedlist.MergeKSortedList([]*linkedlist.ListNode{head1, head2, head3})
	linkedlist.PrintfLinkedListNotCycle(node)
}

func testMergeKSortedList_2() {
	head1 := linkedlist.BuildLinkedList([]int{1, 2, 4})
	head2 := linkedlist.BuildLinkedList([]int{1, 3, 4})
	head3 := linkedlist.BuildLinkedList([]int{4, 5, 6})
	linkedlist.PrintfLinkedListNotCycle(head1)
	linkedlist.PrintfLinkedListNotCycle(head2)
	linkedlist.PrintfLinkedListNotCycle(head3)
	node := linkedlist.MergeKSortedList_2([]*linkedlist.ListNode{head1, head2, head3})
	linkedlist.PrintfLinkedListNotCycle(node)
}

func testReverseLinkedListII() {
	head := linkedlist.BuildLinkedList([]int{1, 2, 3, 4, 5})
	linkedlist.PrintfLinkedListNotCycle(head)
	node := linkedlist.ReverseLinkedListII(head, 2, 4)
	linkedlist.PrintfLinkedListNotCycle(node)

}
