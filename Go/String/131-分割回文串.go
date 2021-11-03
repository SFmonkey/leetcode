package String

import "fmt"

// 回溯, i+1
var reStr [][]string
var pathList []string
var pathStr string

func partition(s string) [][]string {
	reStr = [][]string{}
	pathList = []string{}
	pathStr = ""
	bt([]byte(s), 0)
	//fmt.Println(pathList)
	return reStr
}

func bt(str []byte, startIndex int) {
	fmt.Println(pathStr)
	// 是回文字符串，添加
	if judge([]byte(pathStr)) {
		pathList = append(pathList, pathStr)
		fmt.Println(pathList)
		return
	}
	if startIndex == len(str)-1 {
		return
	}
	for i:=startIndex; i<len(str); i++ {
		pathStr += string(str[i])
		//fmt.Println(pathStr)
		bt(str, i+1)
		pathStr = pathStr[:len(pathStr)-1]
	}
}

// 首尾缩进判断回文
func judge(str []byte) bool {
	if len(str) == 0 {
		return false
	}
	i, j := 0, len(str)-1
	for i <= j {
		if str[i] == str[j] {
			continue
		} else {
			return false
		}
	}
	return true
}