package LinkedList

// 用map记录节点数值
func removeDuplicateNodes(head *ListNode) *ListNode {
	repeat := make(map[int]bool)
	res := head
	// 头部添加一个额外节点
	head = &ListNode{
		Next: head,
	}
	for head != nil && head.Next != nil {
		if _, ok := repeat[head.Next.Val]; !ok {
			repeat[head.Next.Val] = true
			head = head.Next
		} else {
			head.Next = head.Next.Next
		}
	}
	printLL(res)
	return res
}