package ArrayList

import "fmt"

// 冒泡排序
func sortArrayBubble(nums []int)  {
	for i:=0; i<len(nums); i++ {
		for j:=0; j<len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	fmt.Println(nums)
}

// 插入排序
func sortArrayInsert(nums []int) []int {
	for i:=0; i<len(nums); i++ {
		min := nums[i]
		minIdx := i
		for j:=i+1; j<len(nums); j++ {
			if nums[j] < min {
				min = nums[j]
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
	fmt.Println(nums)
	return nums
}

// 快排
func sortArrayQuickSort(nums []int)  {
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func quickSort(arr []int, left, right int) {
	if left < right {
		index := partition(arr, left, right)
		quickSort(arr, left, index-1)
		quickSort(arr, index+1, right)
	}
}

func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot+1
	for i:=index+1; i<=right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	index--
	arr[index], arr[pivot] = arr[pivot], arr[index]
	return index
}