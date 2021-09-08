package ArrayList

// 一次遍历，先根据newInterval[0]确定合并区间(左边界)，再根据newInterval[1]确定新区间的右边界
func insert(intervals [][]int, newInterval []int) [][]int {
	ans := [][]int{}
	i := 0
	// 找到需要合并的起始区间
	for i<len(intervals) && newInterval[0] > intervals[i][1] {
		ans = append(ans, intervals[i])
		i++
	}
	for ; i<len(intervals); {
		// 如果插入区间的右边界大于原区间的左边界，尝试更新左右边界
		if intervals[i][0] <= newInterval[1] {
			// 更新左边界
			if intervals[i][0] < newInterval[0] {
				newInterval[0] = intervals[i][0]
			}
			// 更新右边界
			if intervals[i][1] > newInterval[1] {
				newInterval[1] = intervals[i][1]
			}
			i++
		} else {
			// 超出范围，直接退出
			break
		}
	}
	ans = append(ans, newInterval)
	// 剩余部分直接放入结果数组
	for ; i<len(intervals); i++ {
		ans = append(ans, intervals[i])
	}
	return ans
}