package ArrayList

import (
	"sort"
)

// 排序+双指针
func threeSum(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 || len(nums) < 3{
		return res
	}
	// 排序后双指针头尾缩进搜索
	sort.Ints(nums)
	var left, right int
	tmp := 0
	for i:=0; i<len(nums); i++ {
		if nums[i] > 0 {
			return res
		}
		// 避免重复答案
		if i>0 && nums[i] == nums[i-1] {
			continue
		}
		left = i+1
		right = len(nums)-1
		for left < right {
			tmp = nums[i]+nums[left]+nums[right]
			if tmp == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 避免重复答案
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if tmp > 0 {
				// 大了，右指针左移
				right--
			} else {
				// 小了，左指针右移
				left++
			}
		}
	}
	return res
}
