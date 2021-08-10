package LinkedList

// 因为是升序，判断相邻相等
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	res := head
	for head.Next != nil {
		if head.Val != head.Next.Val {
			head = head.Next
		} else {
			head.Next = head.Next.Next
		}
	}
	pprint(res)
	return res
}
