package ArrayList

// 双指针真头尾缩紧，每次移动高小的指针
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	var area, maxA = 0, 0
	for {
		if i >= j {
			break
		}
		if height[i] < height[j] {
			area = height[i] * (j-i)
			i++
		} else {
			area = height[j] * (j-i)
			j--
		}
		if area > maxA {
			maxA = area
		}
	}
	return maxA
}
