package LinkedList

// 三指针交换，一个记录当前，一个记录前一个节点，一个记录后一个节点
func reverseList(head *ListNode) *ListNode {
	// 1 -> 2 -> 3 -> 4 -> 5
	if head == nil {
		return nil
	}
	var tmp *ListNode
	var prv *ListNode
	prv = nil
	for head != nil {
		tmp = head.Next
		head.Next = prv
		prv = head
		head = tmp
	}
	printLL(prv)
	return prv
}