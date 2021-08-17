package LinkedList

// 额外添加一个头部节点，遍历删除
func removeElements(head *ListNode, val int) *ListNode {
	// head固定，新定义一个l去遍历可以
	// 新定义一个l，用head去遍历的话...l是head初始的状态，没有发生改变
	head = &ListNode{
		Val:  0,
		Next: head,
	}
	l := head
	for l.Next != nil {
		if l.Next.Val == val {
			l.Next = l.Next.Next
		} else {
			l = l.Next
		}
	}
	printLL(head.Next)
	return head.Next
}
