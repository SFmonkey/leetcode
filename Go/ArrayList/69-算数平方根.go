package ArrayList

// 二分查找
func mySqrt(x int) int {
	left, right, mid := 0, x, x
	for left <= right {
		mid = (left+right)/2
		if x/mid > mid {
			left = mid + 1
		} else if x/mid < mid {
			right = mid - 1
		} else {
			return mid
		}
	}
	return right
}
