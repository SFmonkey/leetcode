package String

func reverseStr(s string, k int) string {
	ss := []byte(s)
	length := len(s)
	// 每次移动2k个，然后反转前k个
	for i:=0; i<length; i+= 2*k {
		if i+k <= length {
			reverse(ss[i:i+k])
		} else {
			reverse(ss[i:])
		}
	}
	return string(ss)
}

func reverse(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
