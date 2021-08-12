package ArrayList

import "fmt"

// 三指针，左右两个指针分别记录最后一个0和倒数相连2的位置，中间指针遍历，0往左边放，2往右边放
func sortColors(nums []int)  {
	if len(nums) <= 1 {
		return
	}
	left, right, mid := 0, len(nums)-1, 0
	// 左右指针分别移动到最后一个0和倒数相连的2
	for left<=len(nums)-1 && 0 == nums[left] {
		left++
	}
	for right>=0 && 2 == nums[right] {
		right--
	}
	fmt.Println(left,mid, right)
	// 从left后面的第二个元素开始遍历，因为left后面的第一个元素肯定不是0
	if len(nums) <= 3 {
		mid = left
	} else {
		mid = left+1
	}
	for mid<=right {
		if 0 == nums[mid] {
			fmt.Println("switch 0")
			nums[left], nums[mid] = nums[mid], nums[left]
			left++
		} else if 2 == nums[mid] {
			fmt.Println("switch 2")
			nums[right], nums[mid] = nums[mid], nums[right]
			right--
			mid++
		} else {
			mid++
		}
		fmt.Println(nums, left,mid,right)
		fmt.Println("-------------------------")
	}
}
