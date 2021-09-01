package String

import "fmt"

func compressString(S string) string {
	if len(S) == 0 {
		return ""
	}
	cnt := 0
	res := ""
	for i:=0; i<len(S)-1; i++ {
		if S[i] == S[i+1] {
			cnt++
		} else {
			res += fmt.Sprintf("%s%d", string(S[i]), cnt+1)
			cnt = 0
		}
	}
	res += fmt.Sprintf("%s%d", string(S[len(S)-1]), cnt+1)
	if len(res) >= len(S) {
		res = S
	}
	return res
}