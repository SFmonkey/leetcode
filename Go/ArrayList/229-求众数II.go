package ArrayList

// 摩尔投票法
func majorityElement2(nums []int) []int {
	tmp1, cnt1, tmp2, cnt2 := nums[0], 0, nums[0], 0
	for i:=0; i<len(nums); i++ {
		// 投票，碰到相同的+1
		if tmp1 == nums[i] {
			cnt1++
			continue
		}
		if tmp2 == nums[i] {
			cnt2++
			continue
		}
		// 有一个计数为0了，更新候选
		if cnt1 == 0 {
			tmp1 = nums[i]
			cnt1++
			continue
		}
		if cnt2 == 0 {
			tmp2 = nums[i]
			cnt2++
			continue
		}
		// 没有匹配到相同的，也没有消耗为0的，两个计数都减一
		cnt1--
		cnt2--
	}
	// 投票完得到的两个数，未必是符合条件的(超过1/3)，所以需要重新遍历一遍计数
	cnt1 = 0
	cnt2 = 0
	for i:=0; i<len(nums); i++ {
		if tmp1 == nums[i] {
			cnt1++
		}else if tmp2 == nums[i] {
			cnt2++
		}
	}
	res := []int{}
	if cnt1 > len(nums)/3 {
		res = append(res, tmp1)
	}
	if cnt2 > len(nums)/3 {
		res = append(res, tmp2)
	}
	return res
}