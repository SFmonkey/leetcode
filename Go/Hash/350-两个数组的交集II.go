package Hash

func intersect(nums1 []int, nums2 []int) []int {
	res := []int{}
	record := make(map[int]int)
	for i:=0; i<len(nums1); i++ {
		record[nums1[i]]++
	}
	for i:=0; i<len(nums2); i++ {
		if cnt, ok := record[nums2[i]]; ok {
			if cnt > 0 {
				res = append(res, nums2[i])
				record[nums2[i]]--
			}
		}
	}
	return res
}
