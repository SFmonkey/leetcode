package String

// stack
func isValid(s string) bool {
	stack := []byte{}
	for i:=0; i<len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			switch s[i] {
			case ')':
				if stack[len(stack)-1] != '(' {
					return false
				}
			case ']':
				if stack[len(stack)-1] != '[' {
					return false
				}
			case '}':
				if stack[len(stack)-1] != '{' {
					return false
				}
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
