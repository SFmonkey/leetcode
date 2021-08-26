package LinkedList

// 先算平均长度，不够的平摊到前面几个部分
func splitListToParts(head *ListNode, k int) []*ListNode {
	var partLenNew int
	var part *ListNode
	ans := []*ListNode{}
	// 先从头跑到尾，获取总长度
	l := head
	cnt := 0
	for head != nil {
		head = head.Next
		cnt++
	}
	// 每个部分的平均长度
	partLen := cnt/k
	// 需要额外补充的个数
	extraLen := cnt - partLen*k
	head = l
	for head != nil {
		// 将多余出来的个数平摊到前面的各个部分
		partLenNew = partLen
		part = head
		if extraLen > 0 {
			partLenNew++
		}
		for i:=0; i<partLenNew-1; i++ {
			part = part.Next
		}
		tmp := part.Next
		part.Next = nil
		ans = append(ans, head)
		head = tmp
		if extraLen > 0 {
			extraLen--
		}
	}
	// 不够的后面补空
	if k > cnt {
		for i:=0; i<k-cnt; i++{
			ans = append(ans, nil)
		}
	}
	for i:=0; i<len(ans); i++ {
		printLL(ans[i])
	}
	return ans
}
