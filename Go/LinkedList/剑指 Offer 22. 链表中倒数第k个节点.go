package LinkedList

// 双指针，fast比slow快k个节点
func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	fastCnt := 0
	for fast != nil {
		// fast走过k个节点后，slow再开始动
		if fastCnt >= k {
			slow = slow.Next
		}
		fast = fast.Next
		fastCnt++
	}
	printLL(slow)
	return slow
}