package ArrayList

import (
	"fmt"
	"strconv"
	"strings"
)

// 回溯法搜索，不能重复使用(i+1)
var ipRes []string
var ip []string
var ipTmp string
func restoreIpAddresses(s string) []string {
	// 25525511135
	// ["255.255.11.135","255.255.111.35"]
	ipRes = []string{}
	nums := strings.Split(s, "")
	backTrackingIP(nums, 0)
	return ipRes
}

func backTrackingIP(nums []string, startIndex int) {
	// 终止条件：1. 小于等于3位数，且小于等于255，且第一位不是0；或只有一位且为0
	num := convert(ip)
	if len(ip) <= 3 && num <= 255 && ip[0] != "0" {
		ipRes = append(ipRes, fmt.Sprintf("%d", num))
		return
	}
	if len(ip) == 1 && num == 0 {
		ipRes = append(ipRes, fmt.Sprintf("%d", num))
		return
	}
	for i:=startIndex; i<len(nums); i++ {
		ip = append(ip, nums[i])
		backTrackingIP(nums, i+1)
		ip = ip[:len(ip)-1]
	}
}

func convert(nums []string) (num int) {
	num, _ = strconv.Atoi(strings.Join(nums, ""))
	return
}