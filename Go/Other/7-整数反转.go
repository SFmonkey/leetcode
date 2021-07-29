package Other

import (
	"strconv"
	"strings"
)

// 转字符串，双指针头尾交换
func reverse(x int) int {
	var flag bool = false
	if x < 0 {
		flag  = true
		x *= -1
	}
	strX := strings.Split(strconv.Itoa(x), "")
	i, j := 0, len(strX)-1
	for i < j {
		strX[i], strX[j] = strX[j], strX[i]
		i++
		j--
	}
	intX, _ := strconv.Atoi(strings.Join(strX, ""))
	if flag {
		intX *= -1
	}
	if intX > 2147483647 || intX < -2147483648 {
		return 0
	}
	return intX
}