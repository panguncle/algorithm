package linkedlist

/**
InsertionSortedList: leetcode.147
假定头结点已经有序,
因此从头结点的下一位开始遍历,
如果当前节点的value < 之前节点的value,
断开当前节点, 使用上一个节点连接到当前节点 的下一个节点
则从头结点开始, 遍历, 找到第一个大于当前value的节点的位置, 插入到这个节点,
然后重置 head指针, cur 指针 和 prev指针
*/
func InsertionSortedList(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	cur, prev := head.Next, head
	for cur != nil {
		if cur.Val < prev.Val {
			prev.Next = cur.Next
			dummyHead := &ListNode{
				Val:  -1,
				Next: head,
			}
			cur1 := head
			prev1 := dummyHead
			for cur1 != nil {
				if cur1.Val > cur.Val {
					prev1.Next = cur
					cur.Next = cur1
					break
				}
				prev1 = cur1
				cur1 = cur1.Next
			}
			head = dummyHead.Next
			cur = prev.Next
			continue
		}
		prev = cur
		cur = cur.Next
	}

	return head
}

/**
@reference https://www.kancloud.cn/kancloud/data-structure-and-algorithm-notes/73018
引入dummyHead, dummyHead->next = nil
每次从dummyHead->next开始遍历, 依次和上一轮处理得节点的值进行比较, 知道找不到小于上一轮节点值的节点为止
(dummyHead)
 			O->O->O->O 	 O->O->O->O
			[排好序列的]   [还没有排好序的]
cur即是未排好序的第一个元素
每次从排好序的开始遍历(即是从dummyHead开始), 找到第一个比cur大的元素, 然后将其插入这个元素之前
*/
func InsertionSortedList_2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummyHead := new(ListNode)
	prev, cur := dummyHead, head

	for cur != nil {
		prev = dummyHead
		for prev.Next != nil && prev.Next.Val <= cur.Val {
			prev = prev.Next
		}

		tmp := cur.Next
		cur.Next = prev.Next
		prev.Next = cur
		cur = tmp
	}

	return dummyHead.Next
}
