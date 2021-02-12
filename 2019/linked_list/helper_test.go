package linkedlist

import (
	"fmt"
	"strconv"
	"testing"
)

func TestBuildLinkedList(t *testing.T) {
	nodes := []int{1, 2, 3, 4}
	node := BuildLinkedList(nodes)
	value := ""
	for node != nil {
		value += strconv.Itoa(node.Val) + "-->"
		node = node.Next
	}
	value += "nil"

	fmt.Println(value)
}
