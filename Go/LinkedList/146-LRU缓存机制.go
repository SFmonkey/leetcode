package LinkedList

import "fmt"

// 双向链表
type lruNode struct {
	key			int
	val			int
	prev		*lruNode
	next		*lruNode
}

// 越靠后的节点越近被使用
type LRUCache struct {
	head		*lruNode		// 最久未使用，头节点
	tail		*lruNode		// 尾部节点，快速新增
	cap			int				// capacity
	len			int				// length
}

func ConstructorLRU(capacity int) LRUCache {
	lru := new(LRUCache)
	lru.cap = capacity
	return *lru
}

// 每当节点被使用时，移动到链表尾部
func (this *LRUCache) Get(key int) int {
	val := -1
	if this.head == nil {
		return val
	}
	if this.tail.key == key {
		val = this.tail.val
		if this.tail.prev == nil {
			return val
		}
		//tmp := this.head.next
		//this.tail.next = this.head
		//this.head.next = nil
		//this.head.prev = this.tail
		//tmp.prev = nil
		//this.head = tmp
		//this.tail = this.tail.next
	} else {
		node := this.head
		for node != nil {
			if node.key == key {
				if node.next == nil {
					break
				}
				val = node.val
				node.prev.next = node.next
				node.next.prev = node.prev
				this.tail.next = node
				node.next = nil
				node.prev = this.tail
				this.tail = this.tail.next
				break
			}
			node = node.next
		}
	}
	return val
}

func (this *LRUCache) Put(key int, value int)  {
	if this.len >= this.cap {
		// 删除head
		if this.head.next != nil {
			this.head = this.head.next
			this.head.prev = nil
		} else {
			this.head = nil
			this.tail = nil
		}
		this.len--
	}
	// 直接在链表尾部插入新节点
	tmpNode := &lruNode{
		key:  key,
		val:  value,
		prev: this.tail,
		next: nil,
	}
	if this.tail == nil {
		this.tail = tmpNode
		this.head = tmpNode
	} else {
		this.tail.next = tmpNode
		this.tail = this.tail.next
	}
	this.len++
}

func printLru(node *lruNode)  {
	for node != nil {
		fmt.Print(fmt.Sprintf("%d,%d->", node.key, node.val))
		node = node.next
	}
	fmt.Println("nil")
}
