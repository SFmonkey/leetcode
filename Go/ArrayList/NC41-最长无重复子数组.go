package ArrayList

// 同3-无重复最长子串
// 队列
func maxLength( arr []int ) int {
	maxLen := 0
	queue := make([]int, 0, len(arr))
	for i:=0; i<len(arr); i++ {
		// 有重复，一直出队列直到队列中没有重复，相当于左指针移动
		for isInSlice(queue, arr[i]) {
			queue = queue[1:]
		}
		queue = append(queue, arr[i])
		if maxLen < len(queue) {
			maxLen = len(queue)
		}
	}
	return maxLen
}

func isInSlice(nums []int, target int) bool {
	for i:=0; i<len(nums); i++ {
		if target == nums[i] {
			return true
		}
	}
	return false
}

func maxLength2(arr []int) int {
	maxLen := 0
	left, right := 0, 0
	set := make(map[int]bool)
	for right < len(arr) {
		if _, ok := set[arr[right]]; ok {
			delete(set, arr[left])
			left++
			continue
		}
		set[arr[right]] = true
		right++
		if maxLen < len(set) {
			maxLen = len(set)
		}
	}
	return maxLen
}