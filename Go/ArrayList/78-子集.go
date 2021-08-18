package ArrayList

import (
	"sort"
)
// 回溯法，i+1且全部搜索
var res78 [][]int
var path78 []int
var tmp78 []int

func subsets(nums []int) [][]int {
	res78 = [][]int{}
	path78 = []int{}
	sort.Ints(nums)
	backTrackingS(nums, 0)
	return res78
}

func backTrackingS(nums []int, startIndex int)  {
	print(len(path78), len(nums))
	if len(path78) <= len(nums) {
		tmp78 = make([]int, len(path78))
		copy(tmp78, path78)
		res78 = append(res78, tmp78)
	}
	for i:=startIndex; i<len(nums); i++ {
		path78 = append(path78, nums[i])
		backTrackingS(nums, i+1)
		path78 = path78[:len(path78)-1]
	}
}
