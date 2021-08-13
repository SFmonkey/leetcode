package ArrayList

// 摩尔投票法,先假设第一个数过半数并设cnt=1；遍历后面的数如果相同则cnt+1，不同则减一，当cnt为0时则更换新的数字为候选数
func majorityElement(nums []int) int {
	res, cnt := 0, 0
	for i:=0; i<len(nums); i++ {
		if cnt <= 0 {
			res = nums[i]
			cnt++
			continue
		}
		if res == nums[i] {
			cnt++
		} else {
			cnt--
		}
	}
	return res
}
