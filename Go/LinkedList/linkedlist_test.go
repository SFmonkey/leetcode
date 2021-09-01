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
	lru.Get(2)
	lru.Put(2, 6)
	lru.Get(1)
	lru.Put(1, 5)
	lru.Put(1, 2)
	lru.Get(1)
	res := lru.Get(2)
	t.Log(res)
	printLru(lru.head)
	//t.Log(lru.head, lru.tail)
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
