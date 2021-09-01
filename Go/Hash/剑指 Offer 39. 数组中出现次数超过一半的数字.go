package Hash

// hashmap
func majorityElement(nums []int) int {
	count := make(map[int]int)
	maxNum, maxCnt := 0, 0
	for i:=0; i<len(nums); i++ {
		count[nums[i]] += 1
		if count[nums[i]] >= len(nums)/2 {
			if count[nums[i]] > maxCnt {
				maxNum = nums[i]
				maxCnt = count[nums[i]]
			}
		}
	}
	return maxNum
}

// 摩尔投票法
// 先假设第一个数过半数并设cnt=1；遍历后面的数如果相同则cnt+1，不同则减一，当cnt为0时则更换新的数字为候选数
func majorityElement2(nums []int) int {
	cnt, max := 0, 0
	for i:=0; i<len(nums); i++ {
		if cnt <= 0 {
			max = nums[i]
			cnt++
			continue
		}
		if max == nums[i] {
			cnt++
		} else {
			cnt--
		}
	}
	return max
}