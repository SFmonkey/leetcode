package ArrayList

import "fmt"

// 第i个数组的第j个元素 -> 第j个数组的倒数第i个位置
type Key [2]int
func rotate(matrix [][]int) {
	record := make(map[Key]int, len(matrix)*len(matrix[0]))
	for i:=0; i<len(matrix); i++ {
		for j:=0; j<len(matrix[0]); j++ {
			record[Key{j, len(matrix)-1-i}] = matrix[i][j]
		}
	}
	for k, v := range(record) {
		matrix[k[0]][k[1]] = v
	}
	fmt.Println(matrix)
}