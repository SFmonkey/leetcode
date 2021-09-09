package String

// 纵向扫描
func longestCommonPrefix(strs []string) string {
	if strs[0] == "" {
		return ""
	}
	i := 0
	ans := ""
	s := strs[0][0]
	for {
		for j:=0; j<len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != s {
				return ans
			}
			s = strs[j][i]
		}
		ans += string(s)
		i++
	}
}
