package ArrayList

import (
	"sort"
)

// 先排序，后判断区间重叠
func canAttendMeetings(intervals [][]int) bool {
	if len(intervals) == 0 || len(intervals) == 1 {
		return true
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i:=0; i<len(intervals)-1; i++ {
		if intervals[i][1] > intervals[i+1][0] {
			return false
		}
	}
	// 单独判断最后一个区间
	if intervals[len(intervals)-2][1] > intervals[len(intervals)-1][0] {
		return false
	}
	return true
}
