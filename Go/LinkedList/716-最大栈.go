package LinkedList

import "fmt"

type MaxStack struct {
	stackLen	int
	topNode		*stackNode
	maxNode		*stackNode
}

type stackNode struct {
	val		int
	next	*stackNode
}

func ConstructorStack() MaxStack {
	stack := new(MaxStack)
	stack.stackLen = 0
	return *stack
}

// push的时候，判断和当前最大节点的大小
func (this *MaxStack) Push(x int)  {
	newNode := &stackNode{
		val:  x,
		next: this.topNode,
	}
	this.topNode = newNode
	this.stackLen++
	// 更新最大值节点
	if this.maxNode != nil {
		if this.topNode.val >= this.maxNode.val {
			this.maxNode = this.topNode
		}
	} else {
		this.maxNode = this.topNode
	}
}

func (this *MaxStack) Pop() int {
	// 栈中只有一个元素，全部置空
	if this.stackLen == 1 {
		val := this.topNode.val
		this.topNode = nil
		this.maxNode = nil
		this.stackLen = 0
		return val
	}
	topVal := this.topNode.val
	// 如果最大值节点的值 == 栈顶节点值，遍历找到新的最大值节点
	if this.topNode.val == this.maxNode.val {
		tmp := this.topNode.next
		this.maxNode.val = tmp.val
		for tmp != nil {
			if tmp.val >= this.maxNode.val {
				this.maxNode = tmp
			}
			tmp = tmp.next
		}
	}
	this.topNode = this.topNode.next
	this.stackLen--
	return topVal
}

func (this *MaxStack) Top() int {
	return this.topNode.val
}

func (this *MaxStack) PeekMax() int {
	return this.maxNode.val
}

// 遍历移除
func (this *MaxStack) PopMax() int {
	topVal := this.maxNode.val
	tmp := this.topNode
	// 如果最大值节点的值 == 栈顶节点值，等同于pop
	if this.maxNode.val == this.topNode.val {
		return this.Pop()
	} else {
		// 遍历寻找并删除最大值节点
		for tmp.next != nil {
			if tmp.next.val == this.maxNode.val {
				tmp.next = tmp.next.next
				break
			}
			tmp = tmp.next
		}
	}
	// 寻找新的最大值节点
	this.maxNode.val = this.topNode.val
	tmp = this.topNode
	for tmp != nil {
		if tmp.val >= this.maxNode.val {
			this.maxNode = tmp
		}
		tmp = tmp.next
	}
	this.stackLen--
	return topVal
}

func printStack(node *stackNode)  {
	for node != nil {
		fmt.Print(fmt.Sprintf("%d->", node.val))
		node = node.next
	}
	fmt.Println("nil")
}
