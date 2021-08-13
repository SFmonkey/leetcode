package ArrayList

// 三指针，中间指针遍历，0往左边放，2往右边放
func sortColors(nums []int)  {
	if len(nums) <= 1 {
		return
	}
	left, right, cur := 0, len(nums)-1, 0
	for cur <= right {
		if 0 == nums[cur] {
			nums[cur], nums[left] = nums[left], nums[cur]
			left++
			// 因为left向right遍历，所以左边发生交换后cur需要继续向后移动
			cur++
		} else if 2 == nums[cur] {
			nums[cur], nums[right] = nums[right], nums[cur]
			// 右边交换，交换到cur的新元素有可能是0，有可能需要继续往左边交换，所以cur不动
			right--
		} else {
			// 不需要交换，cur继续向后移动
			cur++
		}
	}
}
