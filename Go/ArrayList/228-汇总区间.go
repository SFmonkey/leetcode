package ArrayList

import "fmt"

// 双指针遍历
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	res := []string{}
	left, right := 0, 0
	for right < len(nums)-1 {
		if nums[right+1] != nums[right]+1 {
			if right == left {
				res = append(res, fmt.Sprintf("%d", nums[left]))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", nums[left], nums[right]))
			}
			left = right+1
		}
		right++
	}
	// 最后1或多个节点单独处理，因为前面len(nums)-1遍历不到
	if left == right {
		res = append(res, fmt.Sprintf("%d", nums[left]))
	}
	if right > left {
		res = append(res, fmt.Sprintf("%d->%d", nums[left], nums[right]))
	}
	return res
}
