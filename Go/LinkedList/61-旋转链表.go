package LinkedList

// 快慢指针，快指针比慢指针快k
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || 0 == k {
		return head
	}
	var slowCnt, fastCnt int
	slow := head
	fast := head
	tmp := head
	for fast.Next != nil {
		// 这里如果k大于链表长度，慢指针一直都会停在第一个节点
		if fastCnt >= slowCnt+k {
			slowCnt++
			slow = slow.Next
		}
		fastCnt++
		fast = fast.Next
	}

	// 刚好是链表长度的倍数，不用反转
	if k % (fastCnt+1) == 0 {
		printLL(head)
		return head
	}
	// k超过链表长度，慢指针向后移动
	if k > fastCnt+1 {
		for i:=0; i<fastCnt-(k%(fastCnt+1)); i++ {
			slow = slow.Next
		}
	}
	head = slow.Next
	slow.Next = nil
	fast.Next = tmp
	printLL(head)
	return head
}