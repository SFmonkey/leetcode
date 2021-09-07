package ArrayList

// 每个元素只遍历一次，使用三个map分别记录行、列、子数独的数字
func isValidSudoku(board [][]byte) bool {
	rows := make(map[int]map[byte]bool, 9)
	cols := make(map[int]map[byte]bool, 9)
	boxs := make(map[int]map[byte]bool, 9)
	boxIdx := 0
	for i:=0; i<9; i++ {
		rows[i] = make(map[byte]bool, 9)
		for j:=0; j<9; j++ {
			if cols[j] == nil {
				cols[j] = make(map[byte]bool, 9)
			}
			num := board[i][j]
			if num == '.' {
				continue
			}
			boxIdx = (i/3)*3+j/3
			if boxs[boxIdx] == nil {
				boxs[boxIdx] = make(map[byte]bool, 9)
			}
			// 如果数字已经出现过，直接返回false
			if _, ok := rows[i][num]; !ok {
				rows[i][num] = true
			} else {
				return false
			}
			if _, ok := cols[j][num]; !ok {
				cols[j][num] = true
			} else {
				return false
			}
			if _, ok := boxs[boxIdx][num]; !ok {
				boxs[boxIdx][num] = true
			} else {
				return false
			}
		}
	}
	return true
}