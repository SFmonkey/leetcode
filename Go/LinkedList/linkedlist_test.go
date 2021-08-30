package LinkedList

import (
	"testing"
)

func TestMergeTwoLists21(t *testing.T)  {
	nums := []int{100,90}
	ll := buildLinkedList(nums)
	swapNodes(ll, 2)
}

func TestDesign(t *testing.T)  {
	// ["MaxStack","push","peekMax","popMax"]
	// [[],[5],[],[]]
	obj := ConstructorStack()
	obj.Push(5)
	obj.Push(1)
	obj.Push(5)
	printStack(obj.topNode)
	t.Log(obj.Pop())
	t.Log(obj.PeekMax())
	t.Log(obj.Top())

	//t.Log(obj.PeekMax())
	//t.Log(obj.Pop())

	//printStack(obj.topNode)
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
