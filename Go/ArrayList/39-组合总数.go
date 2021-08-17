package ArrayList

var res39 [][]int
var path39 []int
var tmp39 []int
var t int
var sum int

// 回溯法搜索
func combinationSum(candidates []int, target int) [][]int {
	res39 = [][]int{}
	path39 = []int{}
	t = target
	backTracking39(candidates, 0)
	return res39
}

func backTracking39(nums []int, startIndex int)  {
	sum = sumPath()
	if sum == t {
		tmp39 = make([]int, len(path39))
		copy(tmp39, path39)
		res39 = append(res39, tmp39)
		return
	}
	for i:=startIndex; i<len(nums); i++ {
		path39 = append(path39, nums[i])
		// 如果加完直接超过target了，就不用再递归了，节省时间
		if sumPath() > t {
			path39 = path39[:len(path39)-1]
			continue
		}
		backTracking39(nums, i)
		path39 = path39[:len(path39)-1]
	}
}

func sumPath() (val int) {
	for i:=0; i<len(path39); i++ {
		val += path39[i]
	}
	return
}