package LinkedList

// l1头部额外加一个，差位比较
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = &ListNode{
		Val:  0,
		Next: l1,
	}
	tmp := &ListNode{}
	head := l1
	for l1.Next != nil && l2 != nil {
		if l2.Val > l1.Next.Val {
			l1 = l1.Next
		} else {
			tmp = l2.Next
			l2.Next = l1.Next
			l1.Next = l2
			l2 = tmp
		}
	}
	if l2 != nil {
		// l1跑到尾部
		for l1.Next != nil {
			l1 = l1.Next
		}
		// 把l2剩余的续上去
		for l2 != nil {
			l1.Next = l2
			l1 = l1.Next
			l2 = l2.Next
		}
	}
	return head.Next
}