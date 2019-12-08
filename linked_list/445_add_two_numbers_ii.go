package linkedlist

// AddTwoNumbersII 445
// Todo
func AddTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	var (
		stack1, stack2 []int
	)

	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}

	var (
		val1, val2, val, toNext int
		node                    *ListNode
	)

	for len(stack1) > 0 || len(stack2) > 0 {
		val1, val2 = 0, 0
		if len(stack1) > 0 {
			val1 = stack1[len(stack1)-1]
			stack1 = stack1[0 : len(stack1)-1]
		}

		if len(stack2) > 0 {
			val2 = stack2[len(stack2)-1]
			stack2 = stack2[0 : len(stack2)-1]
		}

		val = val1 + val2 + toNext
		if val >= 10 {
			val -= 10
			toNext = 1
		} else {
			toNext = 0
		}
		node = &ListNode{
			Val:  val,
			Next: node,
		}

	}

	if toNext == 1 {
		node = &ListNode{
			Val:  1,
			Next: node,
		}
	}

	return node
}
