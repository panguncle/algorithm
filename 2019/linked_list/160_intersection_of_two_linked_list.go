package linkedlist

// IntersectionOfTwoLinkedList: leetcode.160
func IntersectionOfTwoLinkedList(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	// 先取得链表A的末尾
	curA := headA
	for curA.Next != nil {
		curA = curA.Next
	}

	// A尾部接B头部
	curA.Next = headB
	// 判定是否有环
	fast, slow := curA, curA
	var hasCycle bool
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			hasCycle = true
			break
		}
	}

	if !hasCycle {
		return nil
	}

	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	return headA
}
