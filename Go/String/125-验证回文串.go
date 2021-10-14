package String

import (
	"strings"
)

// 双指针头尾缩紧，遇到非字母跳过
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i <= j {
		if s[i] < 48 || (s[i] > 57 && s[i] < 97) || s[i] > 122 {
			i++
			continue
		}
		if s[j] < 48 || (s[j] > 57 && s[j] < 97) || s[j] > 122 {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}