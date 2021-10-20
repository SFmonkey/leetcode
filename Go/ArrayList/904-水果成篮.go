package ArrayList

// 滑动窗口+set
func totalFruit(fruits []int) int {
	basket := make(map[int]int)
	left, right, maxLen, length := 0, 0, 0, 0
	for right < len(fruits) {
		// 该类型已经在篮子中或已使用的篮子小于2，遍历下一个
		if _, ok := basket[fruits[right]]; ok {
			right++
			length++
			continue
		}
		if len(basket) < 2 {
			basket[fruits[right]] = right
			right++
			length++
		} else {
			// 不在篮子中且篮子已满
			// 清空篮子，向左遍历至到第一个重复的，并加入篮子
			basket = make(map[int]int)
			for i:=right-1; i>=0; i-- {
				if fruits[i] != fruits[right-1] {
					left = i+1
					break
				}
			}
			basket[fruits[left]] = left
			// 记录旧前长度
			if length > maxLen {
				maxLen = length
			}
			basket[fruits[right]] = right
			right++
			// 更新长度
			length = right - left
		}
	}
	if length > maxLen {
		maxLen = length
	}
	return maxLen
}