package linkedlist

// LinkedListCycleII: leetcode.142
func LinkedListCycleII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	fast, slow := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			/**
			a + b + c + b = 2 (a + b)
			a + 2b + c = 2a + 2b
			a = c
			*/
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}

			return slow

		}
	}

	return nil
}
