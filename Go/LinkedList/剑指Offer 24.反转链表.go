package LinkedList

func reverseList24(head *ListNode) *ListNode {
	var tmp, pre *ListNode
	for head != nil {
		tmp = head.Next
		head.Next = pre
		pre = head
		head = tmp
	}
	printLL(pre)
	return pre
}