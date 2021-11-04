package Stack

type MyQueue struct {
	inStack		[]int
	outStack	[]int
	queueSize	int
}


func ConstructorQ() MyQueue {
	return MyQueue{
		inStack:   []int{},
		outStack:  []int{},
		queueSize: 0,
	}
}

func (this *MyQueue) Push(x int)  {
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) Pop() int {
	var head int
	if 0 == len(this.outStack) {
		this.TransData()
	}
	head = this.outStack[0]
	this.outStack = this.outStack[1:]
	return head
}

func (this *MyQueue) Peek() int {
	if 0 == len(this.outStack) {
		this.TransData()
	}
	return this.outStack[0]
}

// 把in-stack的数据出栈到out-stack
// 相当于反转in-stack
func (this *MyQueue) TransData() {
	for i:=0; i<len(this.inStack); i++ {
		this.outStack = append(this.outStack, this.inStack[i])
	}
	// 清空in-stack
	this.inStack = []int{}
}

func (this *MyQueue) Empty() bool {
	if 0 == len(this.inStack) && 0 == len(this.outStack) {
		return true
	}
	return false
}