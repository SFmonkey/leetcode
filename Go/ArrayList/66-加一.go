package ArrayList

func plusOne(digits []int) []int {
	f := 1
	for i := len(digits)-1; i >= 0; {
		if digits[i]+f == 10 {
			digits[i] = 0
			i--
		} else {
			// 不需要进位的，加法结束，直接返回
			digits[i] += 1
			return digits
		}
	}
	// 在头部插入1
	digits = append([]int{1}, digits...)
	return digits
}
