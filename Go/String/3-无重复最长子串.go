package String

import (
	"strings"
)

func lengthOfLongestSubstring(s string) (int, []string) {
	if 0 == len(s) {
		return 0, []string{}
	}
	if 1 == len(s) {
		return 1, []string{}
	}
	i, j, maxLen, idx := 0, 1, 0, 0
	subStr := s
	window := s[i:j]
	res := []string{}
	jPointer := ""
	for {
		if j >= len(s){
			if 0 == maxLen || maxLen < len(window) {
				return len(window), []string{window}
			}
			break
		}
		jPointer = string(subStr[j])
		if strings.Contains(window, jPointer) {
			if maxLen < len(window) {
				maxLen = len(window)
				res = []string{window}
			}
			idx = strings.Index(subStr, jPointer)
			i = idx + 1
			j++
			window = subStr[i:j]
			subStr = subStr[i:]
			j -= idx + 1
			continue
		}
		window += jPointer
		j++
	}
	return maxLen, res
}