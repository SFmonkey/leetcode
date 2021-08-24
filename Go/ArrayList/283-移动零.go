package ArrayList

// 双指针，0元素与左起第一个非0元素交换(保证非0元素相对顺序)
func moveZeroes(nums []int)  {
	i, j := 0, 1
	for i < len(nums) {
		if nums[i] == 0 {
			j = i+1
			for j < len(nums) {
				if nums[j] != 0 {
					break
				}
				j++
			}
			if j < len(nums) {
				nums[i], nums[j] = nums[j], nums[i]
			} else {
				return
			}
		}
		i++
	}
}
