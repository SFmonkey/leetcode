package ArrayList

import (
	"fmt"
	"sort"
)

var res47 [][]int
var path47 []int
var tmp47 []int
var repeat map[string]bool
var str string

func permuteUnique(nums []int) [][]int {
	res47 = [][]int{}
	path47 = []int{}
	repeat = make(map[string]bool)
	sort.Ints(nums)
	backTrackingPU(nums, 0)
	return res47
}

func backTrackingPU(nums []int, startIndex int)  {
	if len(path47) == len(nums) {
		tmp47 = make([]int, len(path47))
		copy(tmp47, path47)
		res47 = append(res47, tmp47)
		str = ""
		for j:=0; j<len(tmp47); j++ {
			str += fmt.Sprintf("%d", tmp47[j])
		}
		repeat[str] = true
		return
	}
	for i:=startIndex; i<len(nums); i++ {
		path47 = append(path47, nums[i])
		str = ""
		for j:=0; j<len(path47); j++ {
			str += fmt.Sprintf("%d", path47[j])
		}
		fmt.Println(str)
		if _, ok := repeat[str]; ok {
			path47 = path47[:len(path47)-1]
			continue
		}
		backTrackingPU(nums, 0)
		path47 = path47[:len(path47)-1]
	}
}


