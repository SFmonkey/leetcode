package DP

// dp[i][j] => 从(0,0)到(i,j)的最小路径和
// dp[i][j] = min(dp[i-1][j], dp[i][j-1])+grid[i][j]
// 两条边只有一个方向能走，所以当i=0或j=0时，不用计算min，直接是dp[i][j-1]或dp[i-1][j]
func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}
	for i:=0; i<len(dp); i++ {
		for j:=0; j<len(dp[0]); j++ {
			if i == j && i == 0 {
				dp[i][j] = grid[i][j]
				continue
			}
			if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = getMin(dp[i][j-1], dp[i-1][j]) + grid[i][j]
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func getMin(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
