package ArrayList

// 双指针头尾缩进交换元素
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] == val && nums[right] != val {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		} else if nums[left] == val && nums[right] == val {
			right--
		} else if nums[left] != val {
			left++
		}
	}
	if nums[left] == val {
		return left
	} else {
		return left+1
	}
}

// 其实不用双指针，快慢指针->慢指针覆盖原数组(会改变元素顺序)
func removeElement2(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for i:=0; i<len(nums); i++ {
		if nums[i] != val {
			nums[slow] = nums[i]
			slow++
		}
	}
	return slow
}
