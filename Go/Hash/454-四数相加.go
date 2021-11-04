package Hash

// 先遍历计算a+b的值，记录两个值相加出现的次数
// 在和c+d比较，相加为0了就是找到了
// 其实就是把a+b和c+d分别看作一个整体，相当于就是两数之和了
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	recordAdd := make(map[int]int)
	count := 0
	for _, a := range nums1 {
		for _, b := range nums2 {
			recordAdd[a+b]++
		}
	}
	for _, c := range nums3 {
		for _, d := range nums4 {
			// a+b+c+d=0 => a+b=0-(c+d)
			// go初始化map的值都是0，不存在的也是0
			count += recordAdd[0-(c+d)]
		}
	}
	return count
}
