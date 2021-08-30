package Hash

// map存储，比较当前与后一位，后一位大做加法，小减法
func romanToInt(s string) int {
	nums := make(map[string]int, 7)
	nums["I"] = 1
	nums["V"] = 5
	nums["X"] = 10
	nums["L"] = 50
	nums["C"] = 100
	nums["D"] = 500
	nums["M"] = 1000
	val := 0
	for i:=0; i<len(s)-1; i++ {
		key := string(s[i])
		if nums[key] >= nums[string(s[i+1])] {
			val += nums[key]
		} else {
			val -= nums[key]
		}
	}
	// 最后一位直接加上去
	val += nums[string(s[len(s)-1])]
	return val
}