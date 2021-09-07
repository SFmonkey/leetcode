package String

// 同189，三次反转，反转前n个元素->反转剩余元素->整体反转
func reverseLeftWords(s string, n int) string {
	ss := []byte(s)
	reverse58(ss[:n])
	reverse58(ss[n:])
	reverse58(ss)
	return string(ss)
}

func reverse58(s []byte)  {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}