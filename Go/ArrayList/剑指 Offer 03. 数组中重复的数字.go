package ArrayList

import "sort"

// 排序+比对相邻
func findRepeatNumber1(nums []int) int {
	sort.Ints(nums)
	for i:=0; i<len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

// map记录是否出现过
func findRepeatNumber(nums []int) int {
	record := make(map[int]bool, len(nums))
	for i:=0; i<len(nums); i++ {
		if _, ok := record[nums[i]]; !ok {
			record[nums[i]] = true
		} else {
			return nums[i]
		}
	}
	return -1
}
