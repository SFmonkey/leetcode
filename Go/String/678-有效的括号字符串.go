package String

// stack
func checkValidString(s string) bool {
	// stack 记录下标
	leftStack := []int{}
	starStack := []int{}
	for i:=0; i<len(s); i++ {
		// 左括号和star直接入栈
		if s[i] == '(' {
			leftStack = append(leftStack, i)
			continue
		}
		if s[i] == '*' {
			starStack = append(starStack, i)
			continue
		}
		// 右括号先判断左括号栈
		if len(leftStack) != 0 {
			// 有左括号，pop
			leftStack = leftStack[:len(leftStack)-1]
		} else {
			// 没左括号，但有star, star stack pop
			if len(starStack) != 0 {
				starStack = starStack[:len(starStack)-1]
			} else {
				// 什么都没有，返回false
				return false
			}
		}
	}
	// 如果剩余的(数量大于*数量，直接false; *数量和少于( -> "(*)"
	if len(leftStack) > len(starStack) {
		return false
	}
	// 依次pop，如果有*(这种情况，直接false
	for len(leftStack) != 0 && len(starStack) != 0 {
		if leftStack[len(leftStack)-1] > starStack[len(starStack)-1] {
			return false
		}
		// pop
		leftStack = leftStack[:len(leftStack)-1]
		starStack = starStack[:len(starStack)-1]
	}
	return true
}
