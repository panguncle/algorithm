package linkedlist

// ReverseLinkedList 206
func ReverseLinkedList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var (
		pre *ListNode
		cur = head
	)

	// 在这里, 关注的节点不需要那么多, 关注一个节点就可以
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
