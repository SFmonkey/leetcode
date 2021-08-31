package Hash

import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	intersect := []int{}
	record := make(map[int]bool, len(nums1))
	for i:=0; i<len(nums1); i++ {
		record[nums1[i]] = true
	}
	// nums2排序，相邻相同的元素跳过
	sort.Ints(nums2)
	// 尾部添加一个不同值，避免[x,x,x,x,x]case
	nums2 = append(nums2, nums2[len(nums2)-1]+1)
	for i:=0; i<len(nums2)-1; i++ {
		if nums2[i] == nums2[i+1] {
			continue
		}
		if _, ok := record[nums2[i]]; ok {
			intersect = append(intersect, nums2[i])
		}
	}
	return intersect
}