package LinkedList

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	cntA, cntB := 0, 0
	h := headA
	for h != nil {
		cntA++
		h = h.Next
	}
	h = headB
	for h != nil {
		cntB++
		h = h.Next
	}
	// 保证A比B长
	if cntB < cntA {
		headA, headB = headB, headA
		cntA, cntB = cntB, cntA
	}
	for i:=0; i<(cntA-cntB); i++ {
		headA = headA.Next
	}
	for headA != nil {
		if headA == headB {
			return headA
		}
		headA, headB = headA.Next, headB.Next
	}
	return nil
}
