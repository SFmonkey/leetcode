package LinkedList

func swapPairs(head *ListNode) *ListNode {
	// 头部添加一个额外节点
	head = &ListNode{
		Val:  0,
		Next: head,
	}
	res := head
	tmp := head
	for head != nil && head.Next != nil && head.Next.Next != nil {
		tmp = head.Next.Next
		head.Next.Next = head.Next.Next.Next
		tmp.Next = head.Next
		head.Next = tmp
		head = head.Next.Next
	}
	return res.Next
}
