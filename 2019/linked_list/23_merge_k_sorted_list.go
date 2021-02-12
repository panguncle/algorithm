package linkedlist

// MergeKSortedList: leetcode.23
func MergeKSortedList(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	ret := make([]*ListNode, len(lists))
	ret[0] = lists[0]
	for i := 1; i < len(lists); i++ {
		ret[i] = mergeKSortedList(ret[i-1], lists[i])
	}

	return ret[len(lists)-1]
}

func mergeKSortedList(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}

	if head2 == nil {
		return head1
	}

	dummyHead := new(ListNode)
	head := dummyHead

	node1, node2 := head1, head2

	for node1 != nil && node2 != nil {
		v1, v2 := node1.Val, node2.Val
		cur := new(ListNode)
		if v1 <= v2 {
			cur.Val = v1
			node1 = node1.Next
		} else {
			cur.Val = v2
			node2 = node2.Next
		}
		head.Next = cur
		head = cur
	}

	if node1 != nil {
		head.Next = node1
	}

	if node2 != nil {
		head.Next = node2
	}

	return dummyHead.Next
}

// MergeKSortedListII leetcode.23
func MergeKSortedListII(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	n := len(lists)
	for n > 1 {
		k := (n + 1) / 2 // ????
		for i := 0; i < n/2; i++ {
			lists[i] = mergeKSortedList(lists[i], lists[i+k])
		}
		n = k
	}
	return lists[0]
}
