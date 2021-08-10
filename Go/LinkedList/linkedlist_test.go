package LinkedList

import "testing"

func TestMergeTwoLists21(t *testing.T)  {
	b := ListNode{
		Val:  2,
		Next: nil,
	}
	a := ListNode{
		Val: 1,
		Next: &b,
	}
	e := ListNode{
		Val:  3,
		Next: nil,
	}
	d := ListNode{
		Val: -1,
		Next: &e,
	}
	mergeTwoLists(&a, &d)
}
