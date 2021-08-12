package ArrayList

// 位运算，每个元素依次异或
// i ^ 0 = i
// i ^ i = 0
func singleNumber(nums []int) int {
	res := 0
	for i:=0; i<len(nums); i++ {
		res = res ^ nums[i]
	}
	return res
}
