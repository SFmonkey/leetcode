package DP

// dp[i][j]表示从起点(0,0)到(i,j)的路径数
// 到达(i,j)只能从两个方向：(i-1,j)和(i,j-1)，即dp[i][j]=dp[i-1][j]+dp[i][j-1]
// dp[i][0]和dp[0][j]这两条边的到达路径数都是1
// 逐层遍历
func uniquePaths(m int, n int) int {
	dp := [][]int{}
	for i:=0; i<m; i++ {
		dp[i][0] = 1
	}
	for j:=0; j<n; j++ {
		dp[0][j] = 1
	}
	for i:=1; i<m; i++ {
		for j:=1; j<n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
