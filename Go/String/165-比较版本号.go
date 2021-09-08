package String

import (
	"strconv"
	"strings"
)

// 切分字符串，找到左边第一个非零元素，相互比较
func compareVersion(version1 string, version2 string) int {
	v1Slice := strings.Split(version1, ".")
	v2Slice := strings.Split(version2, ".")
	length := len(v1Slice)
	if len(v1Slice) > len(v2Slice) {
		length = len(v2Slice)
	}
	i := 0
	for ; i<length; i++ {
		// 分别找到左起第一个非0元素
		tmp1 := ""
		for j := 0; j < len(v1Slice[i]); j++ {
			if v1Slice[i][j] != '0' {
				tmp1 = v1Slice[i][j:]
				break
			}
		}
		tmp2 := ""
		for j := 0; j < len(v2Slice[i]); j++ {
			if v2Slice[i][j] != '0' {
				tmp2 = v2Slice[i][j:]
				break
			}
		}
		// 相等，继续下一轮
		if tmp1 == tmp2 {
			continue
		}
		// 转换类型后判断大小
		tmpInt1, _ := strconv.Atoi(tmp1)
		tmpInt2, _ := strconv.Atoi(tmp2)
		if tmpInt1 > tmpInt2 {
			return 1
		} else {
			return -1
		}
	}
	// 如果长度不一致，判断剩余部分是否全是0
	if i < len(v1Slice) {
		for ; i<len(v1Slice); i++ {
			if len(strings.Trim(v1Slice[i], "0")) != 0 {
				return 1
			}
		}
	} else if i < len(v2Slice) {
		for ; i<len(v2Slice); i++ {
			if len(strings.Trim(v2Slice[i], "0")) != 0 {
				return -1
			}
		}
	}
	return 0
}
