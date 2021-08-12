package LinkedList

// 因为是升序，双指针指向头尾进行删除
func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 额外添加一个头节点
	head = &ListNode{
		Val:  0,
		Next: head,
	}
	left := head
	right := head.Next
	for right.Next != nil {
		// 和下一个相同，右指针移动一位
		if right.Val == right.Next.Val {
			right = right.Next
			continue
		}
		// 左右指针间隔一位，说明不需要删除节点，各移动一位
		if left.Next == right {
			left = left.Next
			right = right.Next
		} else {
			// 执行删除
			right = right.Next
			left.Next = right
		}
	}
	// 如果最有几个节点有重复，单独执行删除
	if left.Next != right {
		left.Next = nil
	}
	printLL(head.Next)
	return head.Next
}

