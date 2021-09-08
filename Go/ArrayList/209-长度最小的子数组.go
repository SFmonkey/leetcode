package ArrayList

// 滑动窗口
func minSubArrayLen(target int, nums []int) int {
	i, j, total, minLen := 0, 0, nums[0], len(nums)
	for i < len(nums) {
		// 大于target，找到了符合条件的子数组，更新最小长度，窗口左侧+1
		if total >= target {
			length := j-i+1
			if length < minLen {
				minLen = length
			}
			// 减掉左窗口一个位置，进入下一轮循环，继续判断大小
			total -= nums[i]
			i++
		} else {
			// 小于target，子数组和不够，窗口右侧+1，继续求和
			j++
			if j >= len(nums) {
				break
			}
			total += nums[j]
		}
	}
	// 如果遍历完，左窗口还是初始位置，说明累加和小于target，返回0
	if j >= len(nums) && i == 0 {
		return 0
	}
	return minLen
}