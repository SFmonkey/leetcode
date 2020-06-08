/**
 * @param {number[]} cost
 * @return {number}
 */
/**
 * 动态规划
 * 到达本次所需要的花费应该是前两次中最小的那次
 */
var minCostClimbingStairs = function(cost) {
    let res = 0
    let dp = []
    dp[0] = cost[0]
    dp[1] = cost[1]
    let length = cost.length
    for(let i=2;i<length;i++) {
        dp[i] = Math.min(dp[i-2], dp[i-1]) + cost[i]
    }
    return Math.min(dp[length-2], dp[length-1])
};