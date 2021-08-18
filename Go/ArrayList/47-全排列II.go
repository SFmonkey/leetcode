package ArrayList

import (
	"sort"
)

var res47 [][]int
var path47 []int
var tmp47 []int
var visit []bool

func permuteUnique(nums []int) [][]int {
	res47 = [][]int{}
	path47 = []int{}
	visit = make([]bool, len(nums))
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
		if visit[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !visit[i-1] {
			continue
		}
		path47 = append(path47, nums[i])
		visit[i] = true
		backTrackingPU(nums, 0)
		path47 = path47[:len(path47)-1]
		visit[i] = false
	}
}


