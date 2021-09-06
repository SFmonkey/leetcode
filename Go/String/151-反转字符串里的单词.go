package String

import (
	"strings"
)

// 先整体反转字符串，再逐个单词反转
func reverseWords(s string) string {
	// 去除多余空格
	s = strings.TrimSpace(s)
	ss := []byte(" "+s)
	slow, fast := 0, 0
	// 删除单词中间的多余空格，快慢指针覆盖法，同leetcode27
	for ; fast < len(ss); fast++ {
		if fast-1 > 0 && ss[fast-1] == ss[fast] && ss[fast] == ' ' {
			continue
		}
		ss[slow] = ss[fast]
		slow++
	}
	ss = ss[:slow]
	// 整体反转
	revStr(ss)
	// 逐个单词反转
	i, j := 0, 0
	for j < len(ss) {
		if ss[j] == ' ' {
			revStr(ss[i:j])
			i = j+1
		}
		j++
	}
	return strings.TrimSpace(string(ss))
}

func revStr(s []byte)  {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}