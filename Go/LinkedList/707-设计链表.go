package LinkedList

import "fmt"

type MyLinkedList struct {
	val		int
	next	*MyLinkedList
	prev	*MyLinkedList
}

func Constructor707() MyLinkedList {
	return MyLinkedList{
		val:  0,
		next: nil,
		prev: nil,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}
	head := this.next
	for i:=0; i<index; i++ {
		if head == nil {
			return -1
		}
		head = head.next
	}
	if head == nil {
		return -1
	}
	return head.val
}

func (this *MyLinkedList) AddAtHead(val int)  {
	oldHead := this.next
	this.next = &MyLinkedList{
		val:  val,
		next: oldHead,
		prev: this,
	}
	if oldHead != nil {
		oldHead.prev = this.next
	}
}

func (this *MyLinkedList) AddAtTail(val int)  {
	head := this
	for head.next != nil {
		head = head.next
	}
	head.next = &MyLinkedList{
		val:  val,
		next: nil,
		prev: head,
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	head := this.next
	for i:=1; i<=index; i++ {
		// index > length
		if head == nil {
			return
		}
		// index == length
		if head.next == nil && i == index {
			this.AddAtTail(val)
			return
		}
		head = head.next
	}
	// insert
	preNode := head.prev
	preNode.next = &MyLinkedList{
		val:  val,
		next: head,
		prev: preNode,
	}
	head.prev = preNode.next
}

func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index < 0 {
		return
	}
	head := this.next
	for i:=0; i<index; i++ {
		if head.next == nil {
			return
		}
		head = head.next
	}
	if head.next == nil {
		head.prev.next = nil
		return
	}
	// delete
	head.prev.next = head.next
	head.next.prev = head.prev
}

func Print707(head *MyLinkedList)  {
	for head != nil {
		fmt.Print(fmt.Sprintf("%d->", head.val))
		head = head.next
	}
	fmt.Println("nil")
}