package linkedlist

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// LinkedListHasCycle 141
func LinkedListHasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	// 快慢指针
	fast, slow := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}

	return false
}
