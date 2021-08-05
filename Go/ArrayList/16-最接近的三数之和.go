package ArrayList

import (
	"fmt"
	"math"
	"sort"
)

// 排序+双指针头尾缩进搜索
func threeSumClosest(nums []int, target int) int {
	if 3 == len(nums) {
		return nums[0]+nums[1]+nums[2]
	}
	sort.Ints(nums)
	left, right, tmp, diff := 1, len(nums)-1, 0, 0
	var closest, minDiff int
	minDiff = 1000000
	for i:=0; i<len(nums); i++ {
		for left < right {
			tmp = nums[i]+nums[left]+nums[right]
			diff = int(math.Abs(float64(tmp-target)))
			// 如果正好等于target，直接返回
			if 0 == diff {
				return target
			}
			fmt.Println(nums[i], nums[left], nums[right], tmp, diff, minDiff)
			// 小于最小差值，更新返回值
			if diff < minDiff {
				closest = tmp
				minDiff = diff
			}
			// 移动指针，大于target，右指针左移; 小于target，左指针右移
			if tmp > target {
				right--
			} else {
				left++
			}
		}
	}
	return closest
}