package linkedlist

// AddTwoNumbers 2
func AddTwoNumbers(l1, l2 *ListNode) *ListNode {
	if l1 == l2 && l1 == nil {
		return nil
	}

	dummyHead := &ListNode{
		Val: -1,
	}

	var (
		val1, val2, val, toNext int
		prev                    = dummyHead
	)
	for l1 != nil || l2 != nil {
		val1, val2 = 0, 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		val = val1 + val2 + toNext
		if val >= 10 {
			val -= 10
			toNext = 1
		} else {
			toNext = 0
		}
		prev.Next = &ListNode{
			Val: val,
		}
		prev = prev.Next

	}

	if toNext == 1 {
		prev.Next = &ListNode{
			Val: 1,
		}
	}

	return dummyHead.Next
}
