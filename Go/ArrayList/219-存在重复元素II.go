package ArrayList

// map滑动窗口
func containsNearbyDuplicate(nums []int, k int) bool {
	window := make(map[int]bool, len(nums))
	for i:=0; i<len(nums); i++ {
		_, ok := window[nums[i]]
		// 在window中直接返回
		if ok {
			return true
		}
		window[nums[i]] = true
		// 如果window长度大于k，说明当前窗口没有重复元素
		// 且窗口中i-k元素没用了，因为找不到它的在以k为大小的窗口内的重复元素，移除
		if len(window) > k {
			delete(window, nums[i-k])
		}
	}
	return false
}
