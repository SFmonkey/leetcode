package LinkedList

// 先找到left的前一个和right的后一个，反转left->right，然后拼接
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	head = &ListNode{
		Val:  0,
		Next: head,
	}
	// 这里head也在被改变
	leftPoint, rightPoint := head, head
	cnt := 1
	for cnt <= right+1 {
		if cnt <= left-1 {
			leftPoint = leftPoint.Next
		}
		rightPoint = rightPoint.Next
		cnt++
	}
	l := leftPoint.Next
	var prev *ListNode
	var tmp *ListNode
	for l != rightPoint {
		tmp = l.Next
		l.Next = prev
		prev = l
		l = tmp
	}
	leftPoint.Next = prev
	for prev.Next != nil {
		prev = prev.Next
	}
	prev.Next = rightPoint
	printLL(head.Next)
	return head.Next
}
