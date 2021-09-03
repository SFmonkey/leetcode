package ArrayList

// 单调递减栈
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := []int{}
	for i, v := range temperatures {
		// 当前元素大于栈顶元素，破坏了单调栈
		for len(stack) != 0 && v > temperatures[stack[len(stack)-1]] {
			// pop
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 因为是等几天，所以相减
			ans[top] = i-top
		}
		stack = append(stack, i)
	}
	return ans
}
