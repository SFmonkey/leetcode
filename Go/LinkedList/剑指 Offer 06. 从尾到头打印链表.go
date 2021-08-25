package LinkedList

// 1. 反转链表
func reversePrint1(head *ListNode) []int {
	res := []int{}
	var pre, tmp *ListNode
	for head != nil {
		tmp = head.Next
		head.Next = pre
		pre = head
		head = tmp
	}
	for pre != nil {
		res = append(res, pre.Val)
		pre = pre.Next
	}
	return res
}

// 2. 递归
var res []int
func reversePrint(head *ListNode) []int {
	res = []int{}
	print_(head)
	return res
}

func print_(node *ListNode)  {
	if node == nil {
		return
	}
	print_(node.Next)
	res = append(res, node.Val)
}



