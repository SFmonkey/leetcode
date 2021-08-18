package LinkedList

import (
	"testing"
)

func TestMergeTwoLists21(t *testing.T)  {
	ll := []int{1,2}
	l := buildLinkedList(ll)
	res := isPalindrome(l)
	t.Log(res)
}

func buildLinkedList(nums []int) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	l := head
	for i:=0; i<len(nums); i++ {
		tmp := &ListNode{
			Val:  nums[i],
			Next: nil,
		}
		l.Next = tmp
		l = l.Next
	}
	return head.Next
}
