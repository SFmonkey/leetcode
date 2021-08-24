package LinkedList

// 反转，加一，再反转
func plusOne(head *ListNode) *ListNode {
	var pre, tmp *ListNode
	var flag int
	// 反转链表
	for head != nil {
		tmp = head.Next
		head.Next = pre
		pre = head
		head = tmp
	}
	head = pre
	flag = 1
	for head != nil {
		// 不用进位直接break
		if head.Val+flag <= 9 {
			head.Val++
			flag = 0
			break
		}
		head.Val = 0
		head = head.Next
	}
	// 需要补位
	var top *ListNode
	if head == nil && flag == 1 {
		top = &ListNode{
			Val: 1,
		}
	}
	// 再次反转
	head = pre
	var pre_ *ListNode
	for head != nil {
		tmp = head.Next
		head.Next = pre_
		pre_ = head
		head = tmp
	}
	pre = pre_
	if top != nil {
		top.Next = pre
		pre = top
	}
	return pre
}