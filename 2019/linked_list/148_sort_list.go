package linkedlist

// SortList: leetcode.148
func SortList(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	count := CountLinkedList(head)
	mergeSortLinkedList(head, count)

	return head
}

func mergeSortLinkedList(head *ListNode, length int) *ListNode {
	if head == nil || length <= 0 {
		return head
	}

	midNode := head
	mid := length / 2
	count := 1
	for count < mid {
		midNode = midNode.Next
		count += 1
	}
	rNode := mergeSortLinkedList(midNode.Next, length-mid)
	midNode.Next = nil
	lNode := mergeSortLinkedList(head, mid)

	return mergeLinkedList(lNode, rNode)
}

//
//
//func FindNodeOfLinkedList(head *ListNode, mid int) *ListNode {
//	cur := head
//	count := 1
//	for count < mid {
//		cur = cur.Next
//		count +=1
//	}
//	PrintfLinkedListNotCycle(head)
//	fmt.Printf("mid:%+v\n", cur)
//
//	return cur
//}

func mergeLinkedList(lNode, rNode *ListNode) *ListNode {
	dummyHead := new(ListNode)
	prev := dummyHead

	for lNode != nil && rNode != nil {
		if lNode.Val < rNode.Val {
			prev.Next = lNode
			lNode = lNode.Next
		} else {
			prev.Next = rNode
			rNode = rNode.Next
		}

		prev = prev.Next
	}

	if lNode == nil {
		prev.Next = rNode
	} else {
		prev.Next = lNode
	}

	return dummyHead.Next
}

func CountLinkedList(head *ListNode) int {
	cur := head
	var count int

	for cur != nil {
		count++
		cur = cur.Next
	}

	return count
}
