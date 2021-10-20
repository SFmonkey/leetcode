package String

func backspaceCompare(s string, t string) bool {
	i, j := len(s)-1, len(t)-1
	sSkipNum, tSkipNum := 0, 0
	for {
		// 遇到#向前抵消，抵消到双方目前都没有#，再比较是否一致
		for i >= 0 {
			if s[i] == '#' {
				sSkipNum++
			} else {
				if sSkipNum > 0 {
					sSkipNum--
				} else {
					break
				}
			}
			i--
		}
		for j >= 0 {
			if t[j] == '#' {
				tSkipNum++
			} else {
				if tSkipNum > 0 {
					tSkipNum--
				} else {
					break
				}
			}
			j--
		}
		// 只要有一方到头，就可以退出去了
		if i < 0 || j < 0 {
			break
		}
		if s[i] != t[j] {
			return false
		}
		i--
		j--
	}
	// 两方都到头，说明都是空，一致
	if i == -1 && j == -1 {
		return true
	}
	return false
}