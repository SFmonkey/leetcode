package ArrayList

import "fmt"

// 先找到第一个奇数后面的第一个偶数，每次向后找到奇数和偶数的起始下标，依次互换
func sortArrayByParity(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	// 找到第一个奇数
	i, j := 0, 0
	j = find(nums, i, 1)
	// 没有奇数，直接返回
	if j >= len(nums) {
		return nums
	}
	// 找到第一个奇数后面的第一个偶数
	i = j+1
	i = find(nums, i, 0)
	// j后面没有偶数，说明已经是符合要求的排序，直接返回
	if i >= len(nums) {
		return nums
	}
	// 依次奇偶互换
	for {
		nums[i], nums[j] = nums[j], nums[i]
		// 寻找下一个偶数
		i = find(nums, i, 0)
		// 寻找下个奇数
		j = find(nums, j, 1)
		fmt.Println(i, j)
		if i >= len(nums) || j >= len(nums) {
			break
		}
	}
	return nums
}

func find(nums []int, startIndex, tag int) int {
	for startIndex < len(nums) {
		if nums[startIndex]%2 == tag {
			break
		}
		startIndex++
	}
	return startIndex
}
