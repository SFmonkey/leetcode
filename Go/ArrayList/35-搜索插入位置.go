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

// 二分法, 二分前提：有序&&无重复
func searchInsert2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left+right)/2
		if target < nums[mid] {
			right = mid-1
		} else if target > nums[mid] {
			left = mid+1
		} else {
			return mid
		}
	}
	return right+1
}