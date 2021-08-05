package ArrayList

func searchInsert(nums []int, target int) int {
	if target <= nums[0] {
		return 0
	}
	for i:=0; i<len(nums); i++ {
		if nums[i] == target {
			return i
		}
		if i+1 <= len(nums)-1 {
			if target > nums[i] && target < nums[i+1] {
				return i+1
			}
		}
	}
	return len(nums)
}