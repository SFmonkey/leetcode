package LinkedList

// 偶数链表单独拿出来，拼在奇数链表后面
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	l1 := head
	l2 := &ListNode{
		Next: nil,
	}
	l2Head := l2
	var tmp *ListNode
	for l1 != nil && l1.Next != nil {
		tmp = l1.Next.Next
		l2.Next = l1.Next
		l2.Next.Next = nil
		l1.Next = tmp
		l2 = l2.Next
		if l1.Next == nil {
			break
		}
		l1 = l1.Next
	}
	l1.Next = l2Head.Next
	printLL(head)
	return head
}