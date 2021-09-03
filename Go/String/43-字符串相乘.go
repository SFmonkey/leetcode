package String

import (
	"fmt"
	"strconv"
)

// 先按位相乘，再补0错位相加
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	addSlice := []string{}
	add := 0
	// 保证上面的数位数更多
	if len(num2) > len(num1) {
		num1, num2 = num2, num1
	}
	for i:=len(num2)-1; i>=0; i-- {
		v1, _ := strconv.Atoi(string(num2[i]))
		tmpRes := ""
		for j:=len(num1)-1; j>=0; j-- {
			v2, _ := strconv.Atoi(string(num1[j]))
			tmp := v1*v2 + add
			// 需要进位
			if tmp >= 10 {
				tmpRes += strconv.Itoa(tmp%10)
				add = tmp / 10 % 10
			} else {
				tmpRes += strconv.Itoa(tmp)
				add = 0
			}
		}
		if add != 0 {
			tmpRes += strconv.Itoa(add)
		}
		add = 0
		// 补0，错位
		for k:=0; k<len(num2)-i-1; k++ {
			tmpRes = "0" + tmpRes
		}
		addSlice = append(addSlice, tmpRes)
	}
	fmt.Println(addSlice)
	// 按位相加
	totalRes := ""
	totalInt := 0
	add = 0
	for i:=0; i<len(addSlice[len(addSlice)-1]); i++ {
		totalInt = 0
		for j:=0; j<len(addSlice); j++ {
			// 没有数字的直接补0, 加法不用计算
			if i < len(addSlice[j]) {
				v, _ := strconv.Atoi(string(addSlice[j][i]))
				totalInt += v
			}
		}
		totalInt += add
		if totalInt >= 10 {
			tmpStr := strconv.Itoa(totalInt)
			totalRes = string(tmpStr[len(tmpStr)-1]) + totalRes
			add, _ = strconv.Atoi(tmpStr[:len(tmpStr)-1])
		} else {
			totalRes = strconv.Itoa(totalInt) + totalRes
			add = 0
		}
	}
	if add != 0 {
		totalRes = "1" + totalRes
	}
	fmt.Println(totalRes)
	return totalRes
}