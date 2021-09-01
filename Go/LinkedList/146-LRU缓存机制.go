package LinkedList

import "fmt"

// 最后一个case超时
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
	hashMap		map[int]int		// map用于快速查找key是否存在
}

func ConstructorLRU(capacity int) LRUCache {
	lru := new(LRUCache)
	lru.cap = capacity
	lru.hashMap = make(map[int]int, capacity)
	return *lru
}

// 每当节点被使用时，移动到链表尾部
func (this *LRUCache) Get(key int) int {
	if _, ok := this.hashMap[key]; !ok {
		return -1
	}
	val := -1
	if this.head == nil || this.tail == nil {
		return val
	}
	// 从尾部开始遍历检索...理解错了，从头从尾相同
	// 如果是tail节点，直接返回，无需链表操作
	if this.tail.key == key {
		val = this.tail.val
		if this.tail.prev == nil {
			return val
		}
	} else {
		// 检索，找到节点后将其放置到链表尾部
		node := this.tail
		for node != nil {
			if node.key == key {
				val = node.val
				// 头节点单独执行
				if node.prev == nil {
					// 更新头节点
					this.head = node.next
					this.head.prev = nil
					// 放到链表尾部
					this.tail.next = node
					node.prev = this.tail
					node.next = nil
					this.tail = this.tail.next
					break
				}
				node.prev.next = node.next
				node.next.prev = node.prev
				this.tail.next = node
				node.prev = this.tail
				node.next = nil
				this.tail = this.tail.next
				break
			}
			node = node.prev
		}
	}
	return val
}

func (this *LRUCache) Put(key int, value int)  {
	// 检索key是否已经存在
	if _, ok := this.hashMap[key]; ok {
		node := this.head
		for node != nil {
			if node.key == key {
				node.val = value
				// 放到尾部，同Get操作
				this.Get(key)
				return
			}
			node = node.next
		}
	}
	if this.len >= this.cap {
		// 删除head
		if this.head.next != nil {
			this.head = this.head.next
			this.head.prev = nil
		} else {
			this.head = nil
			this.tail = nil
		}
		delete(this.hashMap, key)
		this.len--
	}
	// 在链表尾部插入新节点
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
	this.hashMap[key] = value
}

func printLru(node *lruNode)  {
	for node != nil {
		fmt.Print(fmt.Sprintf("%d,%d->", node.key, node.val))
		node = node.next
	}
	fmt.Println("nil")
}

// 官方题解
/*
type LRUCache struct {
    size int
    capacity int
    cache map[int]*DLinkedNode
    head, tail *DLinkedNode
}

type DLinkedNode struct {
    key, value int
    prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
    return &DLinkedNode{
        key: key,
        value: value,
    }
}

func Constructor(capacity int) LRUCache {
    l := LRUCache{
        cache: map[int]*DLinkedNode{},
        head: initDLinkedNode(0, 0),
        tail: initDLinkedNode(0, 0),
        capacity: capacity,
    }
    l.head.next = l.tail
    l.tail.prev = l.head
    return l
}

func (this *LRUCache) Get(key int) int {
    if _, ok := this.cache[key]; !ok {
        return -1
    }
    node := this.cache[key]
    this.moveToHead(node)
    return node.value
}


func (this *LRUCache) Put(key int, value int)  {
    if _, ok := this.cache[key]; !ok {
        node := initDLinkedNode(key, value)
        this.cache[key] = node
        this.addToHead(node)
        this.size++
        if this.size > this.capacity {
            removed := this.removeTail()
            delete(this.cache, removed.key)
            this.size--
        }
    } else {
        node := this.cache[key]
        node.value = value
        this.moveToHead(node)
    }
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
    node.prev = this.head
    node.next = this.head.next
    this.head.next.prev = node
    this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
    this.removeNode(node)
    this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
    node := this.tail.prev
    this.removeNode(node)
    return node
}
 */