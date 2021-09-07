package String

import (
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	// 保证上面的数字更长
	if len(num2) > len(num1) {
		num1, num2 = num2, num1
	}
	ans := ""
	num1Slice := []byte(num1)
	num2Slice := []byte(num2)
	// 反转
	r(num1Slice)
	r(num2Slice)
	i, j, add := 0, 0, 0
	for {
		v, _ := strconv.Atoi(string(num1Slice[i]))
		if j < len(num2Slice) {
			v1, _ := strconv.Atoi(string(num2Slice[j]))
			tmp := v1+v+add
			if tmp >= 10 {
				add = 1
				ans = strconv.Itoa(tmp-10) + ans
			} else {
				add = 0
				ans = strconv.Itoa(tmp) + ans
			}
			i++
			j++
		} else {
			v += add
			if v >= 10 {
				add = 1
				ans = strconv.Itoa(v-10) + ans
			} else {
				add = 0
				ans = strconv.Itoa(v) + ans
			}
			i++
		}
		if i >= len(num1Slice) {
			break
		}
	}
	if add == 1 {
		ans = "1" + ans
	}
	return ans
}

func r(s []byte)  {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}