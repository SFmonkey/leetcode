package ArrayList

// 三次反转，整体反转->反转前k个元素->反转剩余元素
func rotateArr(nums []int, k int)  {
	if k >= len(nums) {
		k = k%len(nums)
	}
	// 反转全部
	reverse(nums, 0, len(nums)-1)
	// 反转前k个
	reverse(nums, 0, k-1)
	// 反转剩余
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, i, j int)  {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}