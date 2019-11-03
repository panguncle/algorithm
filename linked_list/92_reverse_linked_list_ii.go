package linkedlist

// ReverseLinkedListII: leetcode.92
func ReverseLinkedListII(head *ListNode, m, n int) *ListNode { // Todo:Adv
	if head == nil || m > n {
		return head
	}

	dummyHead := &ListNode{
		Next: head,
	}

	cur, prev := head, dummyHead

	i := 1
	for cur.Next != nil {
		if i == m {
			prev.Next = reverseLinkedLikedList(cur, n-m)
			break
		}

		prev = cur
		cur = cur.Next
		i++
	}
	return dummyHead.Next

}

func reverseLinkedLikedList(head *ListNode, n int) *ListNode {

	dummyHead := &ListNode{
		Next: head,
	}
	prev, cur := dummyHead, head

	i := 0
	for {
		if i > n {
			break
		}

		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
		i++
	}
	head.Next = cur
	return prev
}
