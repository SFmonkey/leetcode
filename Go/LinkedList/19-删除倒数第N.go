package LinkedList

import "fmt"

// 快慢指针，快指针比慢指针快n
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var fastCnt, slowCnt = 0, 0
	var fast, slow = head, head
	for fast.Next != nil {
		if fastCnt-slowCnt >= n {
			slowCnt++
			slow = slow.Next
		}
		fastCnt++
		fast = fast.Next
	}
	// 只有一个节点
	if fastCnt == 0 {
		return nil
	}
	// 删除的是第一个节点，且总共只有两个节点
	if 0 == slowCnt && n == slowCnt+1 && fastCnt == 1{
		head.Next = nil
		return head
	}
	// 删除的是第一个节点
	if 0 == slowCnt && n == fastCnt+1 {
		return head.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func printLL(l *ListNode)  {
	for l != nil {
		fmt.Print(fmt.Sprintf("%d->", l.Val))
		l = l.Next
	}
	fmt.Println("nil")
}