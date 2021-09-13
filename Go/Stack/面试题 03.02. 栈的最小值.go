package Stack

// 双栈
type MinStack struct {
	primary 	[]int		// 主栈，存数据
	minVal		[]int		// 辅栈，存每次push后的最小值
}

func Constructor() MinStack {
	stack := new(MinStack)
	stack.primary = []int{}
	stack.minVal = []int{}
	return *stack
}

func (this *MinStack) Push(x int)  {
	this.primary = append(this.primary, x)
	min := x
	if len(this.minVal) > 0 {
		min = this.minVal[len(this.minVal)-1]
	}
	if x < min || 0 == len(this.minVal) {
		this.minVal = append(this.minVal, x)
	} else {
		this.minVal = append(this.minVal, min)
	}
}

func (this *MinStack) Pop()  {
	this.primary = this.primary[:len(this.primary)-1]
	this.minVal = this.minVal[:len(this.minVal)-1]
}

func (this *MinStack) Top() int {
	return this.primary[len(this.primary)-1]
}

func (this *MinStack) GetMin() int {
	return this.minVal[len(this.minVal)-1]
}