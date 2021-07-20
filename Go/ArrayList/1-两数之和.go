package ArrayList

func twoSum1(nums []int, target int) []int {
	ans := make([]int, 2, 2)
	for idx, _ := range nums {
		for jdx:=idx+1;jdx<len(nums);jdx++{
			if nums[idx]+nums[jdx] == target {
				ans[0] = idx
				ans[1] = jdx
				return ans
			}
		}
	}
	return ans
}

func twoSum2(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))
	ans := make([]int, 2, 2)
	for idx, n := range nums {
		if val, ok := numsMap[target-n]; ok {
			ans[0] = idx
			ans[1] = val
			return ans
		}
		numsMap[n] = idx
	}
	return ans
}