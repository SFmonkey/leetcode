package ArrayList

import (
	"sort"
)

// 排序+区间重叠，双指针
func mergeArea(intervals [][]int) [][]int {
	res := [][]int{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	start, end := intervals[0][0], intervals[0][1]
	for i:=1; i<len(intervals); i++ {
		if end < intervals[i][0] {
			// 入库，更新start，end
			res = append(res, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][1]
		}
		// 更新end
		if end < intervals[i][1] {
			end = intervals[i][1]
		}
	}
	// 最后一次单独入库
	res = append(res, []int{start, end})
	return res
}