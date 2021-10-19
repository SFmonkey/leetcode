package ArrayList

import "fmt"

// 平方+快排
func sortedSquares(nums []int) []int {
	// 平方
	for i:=0; i<len(nums); i++ {
		nums[i] = nums[i]*nums[i]
	}
	// 快排
	qs(nums, 0, len(nums)-1)
	fmt.Println(nums)
	return nums
}

func qs(arr []int, left, right int)  {
	if left < right {
		idx := partition_(arr, left, right)
		qs(arr, left, idx-1)
		qs(arr, idx+1, right)
	}
}

func partition_(arr []int, left, right int) int {
	pivot := left
	index := pivot+1
	for i:=index; i<=right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	index--
	arr[index], arr[pivot] = arr[pivot], arr[index]
	return index
}
// 双指针头尾判断
func sortedSquares2(nums []int) []int {
	res := make([]int, len(nums))
	i, j, k := 0, len(nums)-1, len(res)-1
	for ; k>=0; k-- {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			res[k] = nums[i]*nums[i]
			i++
		} else {
			res[k] = nums[j]*nums[j]
			j--
		}
	}
	return res
}

