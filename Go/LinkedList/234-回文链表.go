package LinkedList

// 快慢指针找到中间节点，反转后半部分，一次比对
func isPalindrome(head *ListNode) bool {
	// 找中间节点
	slow, fast := head, head
	// 添加一个额外头部节点，用于记录mid的前一个节点
	slowPrev := &ListNode{
		Val:  0,
		Next: head,
	}
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		slowPrev = slowPrev.Next
	}
	// 前半部分链表最后一个节点的next置空
	slowPrev.Next = nil
	// 反转
	var prev *ListNode
	prev = nil
	tmp := slow
	for slow != nil {
		tmp = slow.Next
		slow.Next = prev
		prev = slow
		slow = tmp
	}
	// 遍历比对, len(head) <= len(prev)
	// prev最多只会比head多一个节点(奇数个节点时)，只对比len(head)范围内的节点就ok
	for head != nil {
		if head.Val != prev.Val {
			return false
		}
		head = head.Next
		prev = prev.Next
	}
	return true
}
