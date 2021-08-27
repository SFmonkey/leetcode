package LinkedList

// 快慢指针找到倒数第k个节点，记录pre，然后交换
func swapNodes(head *ListNode, k int) *ListNode {
	i, j := 1, 0
	head = &ListNode{
		Next: head,
	}
	// 找到k和倒数k节点的前一个节点，用于交换
	kth, slow, fast := head, head, head
	for fast != nil {
		if i < k {
			kth = kth.Next
			i++
		}
		if j > k {
			slow = slow.Next
		}
		fast = fast.Next
		j++
	}
	// 需要特殊处理: 只有k值大于总长度一半，正数k和倒数k节点顺序反了，互换一下
	if k > j/2 {
		kth, slow = slow, kth
	}
	// 需要特殊处理: 需要交换的两个节点相邻
	if kth.Next == slow {
		kn := kth.Next.Next
		slow.Next = kn.Next
		kth.Next = kn
		kn.Next = slow
	} else {
		kthNext := kth.Next.Next
		kth.Next.Next = slow.Next.Next
		slowTmp := slow.Next
		slow.Next = kth.Next
		kth.Next = slowTmp
		slowTmp.Next = kthNext
	}
	printLL(head.Next)
	return head.Next
}