package LinkedList

// 找到中间节点，反转后半部分，然后间隔插入前半部分
func reorderList(head *ListNode)  {
	// 头部添加一个额外节点，快慢指针找到中间节点
	head = &ListNode{
		Next: head,
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 反转后半部分
	var prev *ListNode
	mid, tmp := slow.Next, slow
	slow.Next = nil
	for mid != nil {
		tmp = mid.Next
		mid.Next = prev
		prev = mid
		mid = tmp
	}
	// 间隔插入前半部分
	slow = head.Next
	for slow != nil && prev != nil {
		tmp = prev.Next
		prev.Next = slow.Next
		slow.Next = prev
		slow = prev.Next
		prev = tmp
	}
	printLL(head.Next)
}

