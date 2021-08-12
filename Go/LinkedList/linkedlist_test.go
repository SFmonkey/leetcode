package LinkedList

import "testing"

func TestMergeTwoLists21(t *testing.T)  {
	b := ListNode{
		Val:  4,
		Next: nil,
	}
	a := ListNode{
		Val: 2,
		Next: &b,
	}
	e := ListNode{
		Val:  1,
		Next: &a,
	}
	d := ListNode{
		Val: 1,
		Next: &e,
	}
	deleteDuplicatesII(&d)
}
