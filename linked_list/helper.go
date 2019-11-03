package linkedlist

import (
	"fmt"
	"strconv"
)

func BuildLinkedList(value []int) *ListNode {
	var head, prev *ListNode
	if len(value) == 0 {
		return head
	}

	for _, v := range value {
		cur := &ListNode{
			Val: v,
		}
		if head == nil {
			head = cur
			prev = cur
			continue
		}
		prev.Next = cur
		prev = cur

	}

	return head
}

func BuildCycledLinkedList(value []int) *ListNode {
	var head, prev *ListNode
	if len(value) == 0 {
		return head
	}

	for _, v := range value {
		cur := &ListNode{
			Val: v,
		}
		if head == nil {
			head = cur
			prev = cur
			continue
		}
		prev.Next = cur
		prev = cur
	}
	prev.Next = head
	return head
}

func BuildCycleLinkedListII(value []int, idx int) *ListNode {
	if idx > len(value) {
		panic("idx must less than value len")
	}

	head := BuildLinkedList(value)
	var node, pos *ListNode
	node = head
	for i := 0; i < len(value); i++ {
		if i == idx {
			pos = node
		}
		if i < len(value)-1 {
			node = node.Next
		}
	}
	node.Next = pos
	return head
}

func PrintfLinkedListNotCycle(head *ListNode) {
	if LinkedListHasCycle(head) {
		panic("linked list cycle")
	}

	node := head
	var ret string
	for node != nil {
		ret += strconv.Itoa(node.Val) + "-->"
		node = node.Next
	}

	ret += "nil"
	fmt.Println(ret)
}

func BuildIntersectionLinkedList(pre1 []int, pre2 []int, common []int) (*ListNode, *ListNode) {
	head1 := BuildLinkedList(pre1)
	head2 := BuildLinkedList(pre2)
	commHead := BuildLinkedList(common)
	node1 := head1
	for node1.Next != nil {
		node1 = node1.Next
	}
	node1.Next = commHead

	node2 := head2
	for node2.Next != nil {
		node2 = node2.Next
	}
	node2.Next = commHead
	return head1, head2
}
