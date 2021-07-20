package ArrayList

import "testing"

func TestTwoAdd(t *testing.T)  {
	// 2 -> 4 -> 3
	// 5 -> 6 -> 4
	b := ListNode{
		Val:  9,
		Next: nil,
	}
	a := ListNode{
		Val: 9,
		Next: &b,
	}

	f := ListNode{
		Val:  9,
		Next: nil,
	}
	e := ListNode{
		Val:  9,
		Next: &f,
	}
	d := ListNode{
		Val: 9,
		Next: &e,
	}
	res := addTwoNumbers(&a, &d)
	t.Log(res)
}