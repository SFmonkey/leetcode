package DP

// dp[i] => 整数i拆分后的最大乘积
// j从1遍历到i-1，计算dp[i-i]*j，即拆分后半部分，再计算(i-j)*j => 因为题目是拆分成至少两个部分，每次取最大，即 max(dp[i],max(dp[i-j]*j,(i-j)*j))
// dp[1]=1, dp[2]=1
// 从头到尾顺序遍历
func integerBreak(n int) int {
	// 因为下标从1开始，0没有用，后面多申请一个位置，相当于整体后移一位
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i:=3; i<=n; i++ {
		for j:=1; j<=i-1; j++ {
			dp[i] = getMax(dp[i], getMax(dp[i-j]*j, (i-j)*j))
		}
	}
	return dp[n]
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
