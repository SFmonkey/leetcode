package ArrayList

// 向右：(top, left) => (top, right)
// 向下：(top+1, right) => (bottom, right)
// 向左：(bottom, right-1) => (bottom, left+1)
// 向上：(bottom, left) => (top+1, left)
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	res := make([]int, len(matrix)*len(matrix[0]))
	top, left, bottom, right := 0, 0, len(matrix)-1, len(matrix[0])-1
	idx := 0
	for left <= right && top <= bottom {
		// 向右
		for col:=left; col<=right; col++ {
			res[idx] = matrix[top][col]
			idx++
		}
		// 向下
		for row:=top+1; row<=bottom; row++ {
			res[idx] = matrix[row][right]
			idx++
		}
		// 这个判断是避免重复，有些圈层不需要向左和向上遍历，比如单行
		if left < right && top < bottom {
			// 向左
			for col:=right-1; col>=left+1; col-- {
				res[idx] = matrix[bottom][col]
				idx++
			}
			// 向上
			for row:=bottom; row>=top+1; row-- {
				res[idx] = matrix[row][left]
				idx++
			}
		}
		// 下一圈
		top++
		left++
		bottom--
		right--
	}
	return res
}
