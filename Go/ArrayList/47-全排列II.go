package ArrayList

import "sort"

var res47 [][]int
var path47 []int
var tmp47 []int

func permuteUnique(nums []int) [][]int {
	res47 = [][]int{}
	path47 = []int{}
	sort.Ints(nums)
	backTrackingPU(nums, 0)
	return res47
}

func backTrackingPU(nums []int, startIndex int)  {
	if len(path47) == len(nums) {
		tmp47 = make([]int, len(path47))
		copy(tmp47, path47)
		res47 = append(res47, tmp47)
		return
	}
	for i:=startIndex; i<len(nums); i++ {
		if path47[len(path47)-1] == nums[i] {
			continue
		}
		path47 = append(path47, nums[i])
		backTrackingPU(nums, i)
		path47 = path47[:len(path47)-1]
	}
}
