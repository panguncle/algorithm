package linkedlist

// RemoveDupFromSortedList: leetcode.83
func RemoveDupFromSortedList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummyHead := &ListNode{
		Val:  -1,
		Next: head,
	}

	prev := dummyHead
	cur := head

	for cur != nil {
		if cur.Val == prev.Val {
			prev.Next = cur.Next
			cur = cur.Next
			continue
		}
		prev = cur
		cur = cur.Next
	}

	return dummyHead.Next
}
