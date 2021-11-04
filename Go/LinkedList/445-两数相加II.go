package LinkedList

// 链表两数相加，栈存储或反转链表
// 反转链表就不写了，和前面的题有重复
// 7 -> 2 -> 4 -> 3
//      5 -> 6 -> 4
// 7 -> 8 -> 0 -> 7
func addTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := []int{}
	stack2 := []int{}
	// 两个链表节点依次入栈
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}
	// 依次出栈相加
	i, j := len(stack1)-1, len(stack2)-1
	flag := 0
	head := &ListNode{}
	res := head
	for i>=0 && j>=0 {
		tmp := stack1[i]+stack2[j]+flag
		if tmp >= 10 {
			tmp -= 10
			flag = 1
		} else {
			flag = 0
		}
		head.Next = &ListNode{
			Val:  tmp,
			Next: nil,
		}
		head = head.Next
		i--
		j--
	}
	if i >= 0 {
		for k:=i; k>=0; k-- {
			tmp := stack1[k] + flag
			if tmp >= 10 {
				tmp -= 10
				flag = 1
			} else {
				flag = 0
			}
			head.Next = &ListNode{
				Val:  tmp,
				Next: nil,
			}
			head = head.Next
		}
	}
	if j >= 0 {
		for k:=j; k>=0; k-- {
			tmp := stack2[k] + flag
			if tmp >= 10 {
				tmp -= 10
				flag = 1
			} else {
				flag = 0
			}
			head.Next = &ListNode{
				Val:  tmp,
				Next: nil,
			}
			head = head.Next
		}
	}
	if flag == 1 {
		head.Next = &ListNode{
			Val: 	1,
			Next:	nil,
		}
	}
	// 反转
	var pre, tmpNode *ListNode
	res = res.Next
	for res != nil {
		tmpNode = res.Next
		res.Next = pre
		pre = res
		res = tmpNode
	}
	return pre
}