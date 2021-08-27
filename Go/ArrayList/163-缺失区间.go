package ArrayList

import "fmt"

// 双指针寻找
func findMissingRanges(nums []int, lower int, upper int) []string {
	missRange := []string{}
	nums = append([]int{lower-1}, nums...)
	nums = append(nums, upper+1)
	i, j := 0, 1
	for j < len(nums) {
		if nums[j] != nums[i]+1 {
			r := ""
			if nums[j]-nums[i] == 2 {
				r = fmt.Sprintf("%d", nums[i]+1)
			} else {
				r = fmt.Sprintf("%d->%d", nums[i]+1, nums[j]-1)
			}
			missRange = append(missRange, r)
		}
		i++
		j++
	}
	fmt.Println(missRange)
	return missRange
}
