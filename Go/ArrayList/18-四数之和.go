package ArrayList

import "sort"

// 排序+双指针；类似15-三数之和，多固定一个值
func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	ret := [][]int{}
	var fixIdx, left, right int
	sort.Ints(nums)
	for i:=0; i<len(nums)-3; i++ {
		// 避免重复
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for fixIdx=i+1; fixIdx<len(nums); fixIdx++ {
			// 避免重复
			if fixIdx > i+1 && nums[fixIdx] == nums[fixIdx-1] {
				continue
			}
			left = fixIdx+1
			right = len(nums)-1
			for left < right {
				tmp := nums[i]+nums[fixIdx]+nums[left]+nums[right]
				if tmp < target {
					// 小了，左指针向右
					left++
				} else if tmp > target {
					// 大了，右指针向左
					right--
				} else {
					ret = append(ret, []int{nums[i],nums[fixIdx],nums[left],nums[right]})
					// 避免重复
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				}
			}
		}
	}
	return ret
}

