package LinkedList

// 差值法，现将长的链表移到和短的链表一样长的地方，然后开始遍历判断
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	a := headA
	b := headB
	for a.Next != nil {
		lenA++
		a = a.Next
	}
	for b.Next != nil {
		lenB++
		b = b.Next
	}
	// 根据大小重新复制，避免代码冗余
	if lenB > lenA {
		lenA, lenB = lenB, lenA
		headA, headB = headB, headA
	}
	for i:=0; i<lenA-lenB; i++ {
		headA = headA.Next
	}
	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}