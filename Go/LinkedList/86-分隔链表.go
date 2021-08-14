package LinkedList

// 大于x的节点按顺讯单独连接起来，拼接到最后一个比x小的节点的最后
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	// 头部额外添加一个节点
	head = &ListNode{
		Val:  0,
		Next: head,
	}
	l1 := head
	l2 := &ListNode{
		Val:  0,
		Next: nil,
	}
	l2Head := l2
	tmp := l1
	for l1.Next != nil {
		// 比x大的节点拼接到l2
		if l1.Next.Val >= x {
			tmp = l1.Next
			l1.Next = l1.Next.Next
			l2.Next = tmp
			l2 = l2.Next
			l2.Next = nil
		} else {
			l1 = l1.Next
		}
	}
	// l2拼接到原链表尾部
	l1.Next = l2Head.Next
	printLL(head.Next)
	return head.Next
}
