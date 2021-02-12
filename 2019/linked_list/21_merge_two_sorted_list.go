package linkedlist

// MergeTwoSortedList: leetcode.21
func MergeTwoSortedList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	dummyHead := new(ListNode)
	head := dummyHead
	for l1 != nil || l2 != nil {
		var (
			v1, v2 *int
			cur    = new(ListNode)
			win    bool
		)
		if l1 != nil {
			v1 = &l1.Val
		}

		if l2 != nil {
			v2 = &l2.Val
		}

		if v1 != nil && v2 != nil {
			if *v1 <= *v2 {
				cur.Val = *v1
				l1 = l1.Next
			} else {
				cur.Val = *v2
				l2 = l2.Next
			}
			win = true
		}

		if !win && v1 != nil {
			cur.Val = *v1
			l1 = l1.Next
			win = true
		}

		if !win && v2 != nil {
			cur.Val = *v2
			l2 = l2.Next
		}

		head.Next = cur
		head = cur

	}

	return dummyHead.Next
}

// MergeTwoSortedList_2: leetcode.21
func MergeTwoSortedList_2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	dummyHead := new(ListNode)
	head := dummyHead
	for l1 != nil && l2 != nil {
		v1 := l1.Val
		v2 := l2.Val
		cur := new(ListNode)
		if v1 <= v2 {
			cur.Val = v1
			l1 = l1.Next
		} else {
			cur.Val = v2
			l2 = l2.Next
		}

		head.Next = cur
		head = cur
	}

	if l1 != nil {
		head.Next = l1
	}

	if l2 != nil {
		head.Next = l2
	}

	return dummyHead.Next
}
