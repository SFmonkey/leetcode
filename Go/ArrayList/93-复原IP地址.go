package ArrayList

import (
	"fmt"
	"strconv"
	"strings"
)

// 回溯法搜索，不能重复使用(i+1)
var ipRes []string
var ip []string
var ipTmp []int
func restoreIpAddresses(s string) []string {
	// 25525511135
	// ["255.255.11.135","255.255.111.35"]
	ipRes = []string{}
	nums := strings.Split(s, "")
	backTrackingIP(nums, 0)
	fmt.Println(ipTmp)
	return ipRes
}

func backTrackingIP(nums []string, startIndex int) {
	// 终止条件：所有数字都使用了一遍(横向遍历)
	// 最终列表里的所有位置需满足：小于等于3位数，且小于等于255，且第一位不是0；或只有一位且为0
	if len(ip) <= 3 {
		num := convert(ip)
		ipTmp = append(ipTmp, num)
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

/*
class Solution(object):
    def restoreIpAddresses(self, s):
        """
        :type s: str
        :rtype: List[str]
        """
        self.res = []

        def backtrack(s, tmp):
            if len(s) == 0 and len(tmp) == 4:
                self.res.append('.'.join(tmp))
                return
            if len(tmp) < 4:
                for i in range(min(3, len(s))):
                    p, n = s[:i + 1], s[i + 1:]
                    if p and 0 <= int(p) <= 255 and str(int(p)) == p:
                        backtrack(n, tmp + [p])

        backtrack(s, [])
        return self.res
 */