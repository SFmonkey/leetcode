package LinkedList

import (
	"testing"
)

func TestMergeTwoLists21(t *testing.T)  {
	nums := []int{1, 2, 3, 2}
	ll := buildLinkedList(nums)
	deleteDuplicatesUnsorted(ll)
}

func TestDesign(t *testing.T)  {
	obj := Constructor("leetcode")
	obj.Visit("google")
	obj.Visit("facebook")
	obj.Visit("youtube")
	url := obj.Back(1)
	t.Log(url)
	url = obj.Back(1)
	t.Log(url)
	url = obj.Forward(1)
	t.Log(url)
	obj.Visit("linkedin")
	url = obj.Forward(2)
	t.Log(url)
	url = obj.Back(2)
	t.Log(url)
	url = obj.Back(7)
	t.Log(url)
	printHistory(&obj)
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
