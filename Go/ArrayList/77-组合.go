package ArrayList

// 回溯法
var res = [][]int{}
var path = []int{}
var tmp []int

func combine(n int, k int) [][]int {
	backTracking(n, k, 1)
	return res
}

func backTracking(n, k, startIndex int)  {
	if len(path) == k {
		tmp = make([]int, k)
		// golang这里必须copy一下，slice append的坑详见https://juejin.cn/post/6844904001289322504
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := startIndex; i <= n; i++ {
		// 走到下一个状态
		path = append(path, i)
		backTracking(n, k, i+1)
		// 回到上一个状态
		path = path[:len(path)-1]
	}
}
