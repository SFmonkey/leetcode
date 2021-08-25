package ArrayList

// 快慢指针，后向前覆盖，不需要替换
func removeDuplicates2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	// j指向需要被替换的元素，i去寻找替换j的元素
	j := 2
	for i:=2; i<len(nums); i++ {
		if nums[i] != nums[j-2] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}