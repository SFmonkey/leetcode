package ArrayList

var res46 [][]int
var path46 []int
var tmp46 []int

// 求排列，回溯法
func permute(nums []int) [][]int {
	backTracking46(nums, 0)
	return res46
}

func backTracking46(nums []int, startIndex int)  {
	if len(path46) == len(nums) {
		tmp46 = make([]int, len(nums))
		copy(tmp46, path46)
		res46 = append(res46, tmp46)
		return
	}
	for i:=startIndex; i<len(nums); i++ {
		// 需要剪枝，去除掉重复的元素
		if inSlice(nums[i]) {
			continue
		}
		path46 = append(path46, nums[i])
		backTracking46(nums, 0)
		path46 = path46[:len(path46)-1]
	}
}

func inSlice(val int) bool {
	for j:=0; j<len(path46); j++ {
		if val == path46[j] {
			return true
		}
	}
	return false
}