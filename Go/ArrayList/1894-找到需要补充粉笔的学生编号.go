package ArrayList

// 直接模拟减会超时
func chalkReplacer(chalk []int, k int) int {
	i := 0
	for k != 0 {
		if i == len(chalk) {
			i = 0
		}
		if k < chalk[i] {
			break
		}
		k -= chalk[i]
		i++
	}
	return i
}

// 优化模拟减，先累加一轮，再取余
func chalkReplacer2(chalk []int, k int) int {
	total := 0
	for _, v := range chalk {
		total += v
	}
	k %= total
	for i, c := range chalk {
		if k < c {
			return i
		}
		k -= c
	}
	return 0
}