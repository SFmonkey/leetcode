package ArrayList

import (
	"sort"
)

// 双指针
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 || (len(nums)==1 && nums[0] != target){
		return []int{-1,-1}
	}
	res := []int{}
	i, j := 0, len(nums)-1
	for i <= j {
		// 一边找到后，就不再继续寻找
		if i != -1 {
			if nums[i] == target {
				res = append(res, i)
				i = -1
			} else {
				i++
			}
		}
		if j != len(nums) {
			if nums[j] == target {
				res = append(res, j)
				// 保证j比i大
				j = len(nums)
			} else {
				j--
			}
		}
		if len(res) == 2 {
			break
		}
	}
	if i >= j {
		return []int{-1,-1}
	}
	sort.Ints(res)
	return res
}