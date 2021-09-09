package String

// 纵向扫描
func longestCommonPrefix(strs []string) string {
	if strs[0] == "" {
		return ""
	}
	i := 0
	ans := ""
	for {
		if i >= len(strs[0]) {
			return ans
		}
		s := strs[0][i]
		for j:=1; j<len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != s {
				return ans
			}
		}
		ans += string(s)
		i++
	}
}
