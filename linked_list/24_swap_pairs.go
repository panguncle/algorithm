package linkedlist

// SwapPairs 24
func SwapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummyHead := &ListNode{
		Next: head,
	}
	prev, cur := dummyHead, head

	for cur != nil && cur.Next != nil {
		tmp := cur.Next
		cur.Next = cur.Next.Next
		tmp.Next = cur
		prev.Next = tmp
		prev = cur
		cur = cur.Next
	}
	return dummyHead.Next
}
