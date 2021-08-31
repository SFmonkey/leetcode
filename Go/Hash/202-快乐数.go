package Hash

import (
	"strconv"
)

func isHappy(n int) bool {
	return happyIsOk(intToList(n))
}

func happyIsOk(nums []int) bool {
	// 终止条件：只有一位数且大于1小于等于3
	if len(nums) == 1 && nums[0] > 1 && nums[0] <= 3 {
		return false
	}
	if len(nums) == 1 && nums[0] == 1 {
		return  true
	}
	val := 0
	for i:=0; i<len(nums); i++ {
		val += nums[i]*nums[i]
	}
	return happyIsOk(intToList(val))
}

func intToList(n int) []int {
	tmp := []int{}
	k := 1
	for i:=0; i<len(strconv.Itoa(n)); i++{
		if (n/k)%10 == 0 {
			k *= 10
			continue
		}
		tmp = append(tmp, (n/k)%10)
		k *= 10
	}
	return tmp
}