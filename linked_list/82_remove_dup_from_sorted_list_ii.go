package linkedlist

// RemoveDupFromSortedListII: leetcode: 82
func RemoveDupFromSortedListII(head *ListNode) *ListNode {

	var cur, prev *ListNode
	if head == nil {
		return cur
	}

	dummyHead := &ListNode{
		Val:  -1,
		Next: head,
	}
	cur = head
	prev = dummyHead

	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			// 找出到最后一个相同的value
			for cur != nil && cur.Next != nil {
				if cur.Val != cur.Next.Val {
					break
				}
				cur = cur.Next
			}
			prev.Next = cur.Next
			cur = cur.Next
			continue
		}
		prev = cur
		cur = cur.Next
	}

	return dummyHead.Next
}
