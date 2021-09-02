package String

import (
	"strings"
)
// 回溯+剪枝
var res []string
var path string
var pathIndex string		// 记录path中的下标，字符串表示
var mapSet map[string]bool
func permutation(s string) []string {
	res = []string{}
	mapSet = make(map[string]bool)
	backTrackingStr(s)
	return res
}

func backTrackingStr(s string)  {
	if len(path) == len(s) {
		res = append(res, path)
		return
	}
	for i:=0; i<len(s); i++ {
		if strings.Contains(pathIndex, string(rune(i))) {
			continue
		}
		if _, ok := mapSet[path+string(s[i])]; !ok {
			mapSet[path] = true
		} else {
			continue
		}
		path += string(s[i])
		pathIndex += string(rune(i))
		backTrackingStr(s)
		path = path[:len(path)-1]
		pathIndex = pathIndex[:len(pathIndex)-1]
	}
}