package ArrayList

import "sort"

// 回溯法+剪枝
var res90 [][]int
var path90 []int
var tmp90 []int

func subsetsWithDup(nums []int) [][]int {
	res90 = [][]int{}
	path90 = []int{}
	sort.Ints(nums)
	backTracking90(nums, 0)
	return res90
}

func backTracking90(nums []int, startIndex int)  {
	tmp90 = make([]int, len(path90))
	copy(tmp90, path90)
	res90 = append(res90, tmp90)
	for i:=startIndex; i<len(nums); i++ {
		// 剪枝/去重
		// 第一个元素不用判断，而且第一个就跳过的话，就不会递归了，startIndex对应的是每个子分支中最左边的一支
		// 相邻且相等的元素不处理，因为前一步search过了
		if i != startIndex && nums[i] == nums[i-1] {
			continue
		}
		path90 = append(path90, nums[i])
		backTracking90(nums, i+1)
		path90 = path90[:len(path90)-1]
	}
}