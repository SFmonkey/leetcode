package ArrayList

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//var ans *ListNode	// todo: * 和 & 在声明变量时的区别
	ans := ListNode{
		Val:  0,
		Next: nil,
	}
	var head ListNode
	head = ans
	f := 0
	for l1 != nil && l2 != nil {
		tmp := &ListNode{}
		v := l1.Val + l2.Val + f
		if v >= 10 {
			f = 1
			v = v - 10
		} else {
			f = 0
		}
		tmp.Val = v
		ans.Next = tmp
		ans = *ans.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	head = *head.Next
	return &head
}

func pprint(l *ListNode) {
	n := l
	for n != nil {
		fmt.Print(n.Val)
		n = n.Next
	}
}