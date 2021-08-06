package ArrayList

/*
	从右向左search，找到第一个破坏升序的元素
	与右起第一个比它大的元素交换
	将其右边的数字按升序排序(因为之前确定了右向左升序，所以直接反转就行)
 */
func nextPermutation(nums []int)  {
	i, j := len(nums)-1, len(nums)-1
	for i > 0 {
		// 判断升序(右向左)
		if nums[i-1] >= nums[i] {
			i--
			continue
		}
		// 找到右起第一个比nums[i-1]大的元素，交换
		for {
			if nums[j] > nums[i-1] {
				nums[i-1], nums[j] = nums[j], nums[i-1]
				break
			}
			j--
		}
		break
	}
	// i-1之后的元素升序排列，反转；本身就是最大的话，i就是0
	j = len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}