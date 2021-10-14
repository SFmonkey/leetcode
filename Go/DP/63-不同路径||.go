package DP

// dp[i][j] => 从(0,0)到(i,j)的不同路径数
// 到达(i,j)只能从两个方向：(i-1,j)和(i,j-1)，即dp[i][j]=dp[i-1][j]+dp[i][j-1]
// dp[i][0]和dp[0][j]这两条边的到达路径数都是1,有障碍的点路径为0,两条边上有障碍之后的点也是0
// 逐层遍历
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// 遇到obstacleGrid为1的，后面就不再遍历了，默认是0
	for i:=0; i<m && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}
	for j:=0; j<n && obstacleGrid[0][j] == 0; j++ {
		dp[0][j] = 1
	}
	for i:=1; i<m; i++ {
		for j:=1; j<n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
