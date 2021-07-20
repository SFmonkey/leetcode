package ArrayList

import "testing"

func TestTwoAdd(t *testing.T)  {
	// 2 -> 4 -> 3
	// 5 -> 6 -> 4
	c := ListNode{
		Val:  3,
		Next: nil,
	}
	b := ListNode{
		Val:  4,
		Next: &c,
	}
	a := ListNode{
		Val: 2,
		Next: &b,
	}

	f := ListNode{
		Val:  4,
		Next: nil,
	}
	e := ListNode{
		Val:  6,
		Next: &f,
	}
	d := ListNode{
		Val: 5,
		Next: &e,
	}
	res := addTwoNumbers(&a, &d)
	t.Log(res)
}