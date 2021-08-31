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
	lru := ConstructorLRU(2)
	lru.Put(2, 1)
	lru.Put(2, 2)
	printLru(lru.head)
	t.Log(lru.head, lru.tail)
	res := lru.Get(2)
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
