package LinkedList

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//var ans *ListNode	// todo: * 和 & 在声明变量时的区别
	var add, f int
	head := &ListNode{}
	ans := head
	for l1 != nil || l2 != nil {
		add = 0
		add += f
		if l1 != nil {
			add += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			add += l2.Val
			l2 = l2.Next
		}
		if add >= 10 {
			f = 1
			add -= 10
		} else {
			f = 0
		}
		tmp := ListNode{
			Val:  add,
			Next: nil,
		}
		head.Next = &tmp
		head = head.Next
	}
	if 1 == f {
		head.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return ans.Next
}

func pprint(l *ListNode) {
	n := l
	for n != nil {
		fmt.Print(n.Val)
		n = n.Next
	}
}