package LinkedList

// 使用map记录, 两次遍历
func deleteDuplicatesUnsorted(head *ListNode) *ListNode {
	record := make(map[int]int)
	head = &ListNode{
		Next: head,
	}
	p := head.Next
	for p != nil {
		record[p.Val]++
		p = p.Next
	}
	p = head
	for p != nil && p.Next != nil {
		if record[p.Next.Val] != 1 {
			p.Next = p.Next.Next
			continue
		}
		p = p.Next
	}
	printLL(head.Next)
	return head.Next
}