package ArrayList

import "fmt"

// 双指针+hash表
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//numMap := make(map[int]int, len(nums))
	numMap := map[int]int{}
	i, j := 0, 1
	for {
		fmt.Println(i, j, nums, numMap)
		if j == len(nums) {
			break
		}
		if _, ok := numMap[nums[i]]; !ok {
			numMap[nums[i]] = i
			i++
			j++
		} else {
			if _, ok := numMap[nums[j]]; !ok {
				numMap[nums[j]] = j
				// 如果j不在map中，替换
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j++
			} else {
				// 如果j在map中，i保持原位，j继续向后搜索
				j++
			}
		}
	}
	if _, ok := numMap[nums[i]]; !ok {
		numMap[nums[i]] = i
	}
	return len(numMap)
}
